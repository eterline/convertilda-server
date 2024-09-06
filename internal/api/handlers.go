package api

import (
	"fmt"
	"mime/multipart"
	"path"
	"slices"
	"strings"

	"github.com/eterline/convertilda-api/internal/database"
	"github.com/gofiber/fiber/v2"
)

const STORAGE = "./store/download/"

func (s *Server) document(c *fiber.Ctx) error {
	fileType := "documents"
	target := new(fileTarg)
	c.ReqHeaderParser(target)
	fmt.Println(target.Target)
	if !slices.Contains(getDocExtensions(), target.Target) {
		return UncorrectTargetErr()
	}
	form, err := c.MultipartForm()
	if err != nil {
		c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		return err
	}

	srcName, err := saveFormFile(fileType, form, c, getDocExtensions())
	if err != nil {
		c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		return err
	}

	// TODO: изменить сдесь вход.
	data := &database.ConvertFiles{
		SourceFile: srcName,
		OutFile:    "s",
		Target:     target.Target,
		FileType:   fileType,
	}
	s.db.Create(data)

	fmt.Println(srcName)
	c.Status(fiber.StatusAccepted)
	return nil
}

// TODO: Для форнта реализовать проверку размера файла, апи пускай этим не занимается.
func saveFormFile(field string, form *multipart.Form, c *fiber.Ctx, exts []string) (string, error) {
	f := form.File[field]
	fmt.Println(len(f) == 1)
	if len(f) == 1 {
		file := f[0]
		filen := strings.ReplaceAll(file.Filename, " ", "")
		if slices.Contains(exts, path.Ext(filen)[1:]) {
			if err := c.SaveFile(file, fmt.Sprint(STORAGE, filen)); err != nil {
				c.Status(fiber.StatusInternalServerError)
				return "", err
			}
			c.Status(fiber.StatusAccepted)
			return filen, nil
		}
		return "", IncorrectExtensionErr()
	}
	return "", TooManyFilesErr()
}
