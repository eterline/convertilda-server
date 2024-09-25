package api

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"path"
	"slices"
	"strings"

	"github.com/eterline/convertilda-api/internal/convert"
	"github.com/eterline/convertilda-api/internal/database"
	"github.com/gofiber/fiber/v2"
)

const STORAGE = "./store/download/"
const OUTPUT = "./store/converted/"

func (s *Server) process(c *fiber.Ctx) error {
	file := convert.GetTypes()
	var fileType string
	var extList []string
	switch c.Params("filetype") {
	case file.Audio:
		fileType = file.Audio
		extList = convert.AudioExtensions()
	case file.Document:
		fileType = file.Document
		extList = convert.DocExtensions()
	case file.Photo:
		fileType = file.Photo
	default:
		fileType = "none"
	}
	target := new(fileTarg)
	c.ReqHeaderParser(target)
	if !slices.Contains(extList, target.Target) {
		return UncorrectTargetErr()
	}
	form, err := c.MultipartForm()
	if err != nil {
		c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		return err
	}
	srcName, outputName, err := saveFormFile("file", form, c, extList)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		return err
	}

	converTo := convert.New(srcName, outputName, fileType, target.Target)
	err = converTo.Convert(STORAGE, OUTPUT)
	if err != nil {
		return FileConvertionErr()
	}

	res := File{
		URL: c.BaseURL() + "/converted/" + strings.Split(converTo.Output, ".")[0] + "." + converTo.Target,
	}
	r, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	data := &database.ConvertFiles{
		SourceFile: converTo.Name,
		OutFile:    converTo.Output,
		Target:     converTo.Target,
		FileType:   converTo.Type,
		URL:        res.URL,
	}
	s.db.Create(data)
	c.Status(fiber.StatusAccepted).SendString(string(r))
	return nil
}

func (s *Server) sendFile(c *fiber.Ctx) error {
	filename := c.Params("filename")
	return c.Status(fiber.StatusAccepted).SendFile(OUTPUT + filename)
}

func saveFormFile(field string, form *multipart.Form, c *fiber.Ctx, exts []string) (string, string, error) {
	f := form.File[field]
	if len(f) == 1 {
		file := f[0]
		extension := path.Ext(file.Filename)
		if slices.Contains(exts, extension[1:]) {
			saved := convert.RandStringBytesRmndr(RAND_NAME_LEN) + extension
			if err := c.SaveFile(file, fmt.Sprint(STORAGE, saved)); err != nil {
				c.Status(fiber.StatusInternalServerError)
				return "", "", err
			}
			c.Status(fiber.StatusAccepted)
			return file.Filename, saved, nil
		}
		return "", "", IncorrectExtensionErr()
	}
	return "", "", TooManyFilesErr()
}
