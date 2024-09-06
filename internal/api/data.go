package api

import (
	"github.com/gofiber/fiber/v2"
)

type fileTarg struct {
	Target   string `reqHeader:"target"`
	Bitrate  int    `reqHeader:"bitrate"`
	Quallity string `reqHeader:"quallity"`
}

func getDocExtensions() []string {
	return []string{"pdf", "doc", "epub", "docx"}
}

func getAudioExtensions() []string {
	return []string{"wav", "flac", "mp3", "ogg"}
}

func TooManyFilesErr() error {
	return &fiber.Error{
		Code:    fiber.StatusBadRequest,
		Message: "Too many files in request. Must be one.",
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
