package main

import (
	"github.com/eterline/convertilda-api/internal/api"
	"github.com/eterline/convertilda-api/internal/database"
	"github.com/eterline/convertilda-api/internal/logging"
	"github.com/eterline/convertilda-api/internal/settings"
)

func main() {
	addr := settings.Adress{
		IP:   []byte("0.0.0.0"),
		Port: 8080,
	}
	logs := settings.Logging{
		LogPath:  "./logs/",
		LogLevel: 1,
	}
	cfg := settings.Config{
		Adress:  addr,
		Logging: logs,
		DbName:  "test.db",
	}
	logging.InitLogfile(cfg.Logging)
	db := database.ConnDB(cfg)
	srv := api.New(cfg, db)
	srv.Run()
}
