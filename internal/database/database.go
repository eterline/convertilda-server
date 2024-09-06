package database

import (
	"log"
	"os"

	"github.com/eterline/convertilda-api/internal/settings"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnDB(cfg settings.Config) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(cfg.DbName), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection to database failed...")
		os.Exit(2)
	}
	log.Println("Connection to database success.")
	db.AutoMigrate(&ConvertFiles{})
	return db
}
