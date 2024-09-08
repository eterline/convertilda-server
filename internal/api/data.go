package api

import (
	"github.com/gofiber/fiber/v2"
)

const RAND_NAME_LEN = 16

type fileTarg struct {
	Target   string `reqHeader:"target"`
	Bitrate  int    `reqHeader:"bitrate"`
	Quallity string `reqHeader:"quallity"`
}

func TooManyFilesErr() error {
	return &fiber.Error{
		Code:    fiber.StatusBadRequest,
		Message: "Too many files in request. Must be one.",
	}
}

func FileConvertionErr() error {
	return &fiber.Error{
		Code:    fiber.StatusInternalServerError,
		Message: "File converting failed.",
	}
}

func UncorrectTargetErr() error {
	return &fiber.Error{
		Code:    fiber.StatusBadRequest,
		Message: "Uncorrect target for filetype.",
	}
}

func IncorrectExtensionErr() error {
	return &fiber.Error{
		Code:    fiber.StatusBadRequest,
		Message: "Incorrect file extension.",
	}
}

func IncorrectTypeErr() error {
	return &fiber.Error{
		Code:    fiber.StatusNotAcceptable,
		Message: "Incorrect file type.",
	}
}

type File struct {
	URL string
}
