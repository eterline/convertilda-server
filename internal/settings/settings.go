package settings

import (
	"net"
)

// TODO: Сделать парсинг файла конфига. Возможно задать ему несколько мест расположения.

type Config struct {
	Adress
	Logging
	DbName string
}

type Adress struct {
	IP   net.IP
	Port int
}

type Logging struct {
	LogPath  string
	LogLevel int
}
