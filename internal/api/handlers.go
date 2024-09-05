package api

import (
	"fmt"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
)

const STORAGE = "./store/download/"

func (s *Server) document(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		return err
	}
	files := form.File["documents"]
	if err := save(files, c); err != nil {
		c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		return err
	}
	// TODO: доделать запись в бд под орм и отправку в конвертацию.
	target := new(file)
	c.ReqHeaderParser(target)
	fmt.Println(target.Target == "")
	c.Status(fiber.StatusAccepted)
	return nil
}

// TODO: Для форнта реализовать проверку размера файла, апи пускай этим не занимается.
func save(f []*multipart.FileHeader, c *fiber.Ctx) error {
	if len(f) == 1 {
		file := f[0]
		if err := c.SaveFile(file, fmt.Sprint(STORAGE, file.Filename)); err != nil {
			c.Status(fiber.StatusInternalServerError)
			return err
		}
		c.Status(fiber.StatusAccepted)
		return nil
	}
	c.Status(fiber.StatusBadRequest)
	return &fiber.Error{
		Code:    fiber.StatusBadRequest,
		Message: "Too many files in request.",
	}
}
