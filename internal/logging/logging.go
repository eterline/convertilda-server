package logging

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/eterline/convertilda-api/internal/settings"
)

func InitLogfile(cfg settings.Logging) {
	logfile := cfg.LogPath + logname()
	file, _ := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	w := io.MultiWriter(os.Stdout, file)
	log.SetOutput(w)
}

func logname() string {
	return fmt.Sprintf(
		"api_%v_%v_%v-%v_%v_%v.log",
		time.Now().Year(), time.Now().Month(), time.Now().Day(),
		time.Now().Hour(), time.Now().Minute(), time.Now().Second(),
	)
}
