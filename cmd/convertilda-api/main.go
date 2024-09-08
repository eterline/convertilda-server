package main

import (
	"github.com/eterline/convertilda-api/internal/api"
	"github.com/eterline/convertilda-api/internal/database"
	"github.com/eterline/convertilda-api/internal/logging"
	"github.com/eterline/convertilda-api/internal/settings"
)

func main() {
	cfg := settings.MustArgs()

	logging.InitLogfile(cfg.Logging)
	db := database.ConnDB(cfg)
	srv := api.New(cfg, db)
	srv.Run()
}
