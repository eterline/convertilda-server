package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eterline/convertilda-api/internal/settings"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	app      *fiber.App
	settings settings.Config
}

func (s *Server) Run() {
	s.routeInit()
	addr := fmt.Sprintf(`%s:%v`, string(s.settings.IP), s.settings.Port)
	log.Printf("Serever has been started on: %s", addr)
	s.app.Listen(addr)
}

func New(cfg settings.Config) *Server {
	return &Server{
		app:      fiber.New(),
		settings: cfg,
	}
}

func (s *Server) routeInit() {

	// TODO: Вынести логгер в отдельное место и настроить его в запись файла.
	s.app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	s.app.Use(filesystem.New(filesystem.Config{
		Root:   http.Dir("./converted"),
		Browse: true,
	}))
	s.app.Post("/api/convert/document", s.document)
}
