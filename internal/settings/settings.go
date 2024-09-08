package settings

import (
	"flag"
	"net"
)

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

func MustArgs() Config {
	port := flag.Int("port", 8080, "App listening port.")
	ip := flag.String("ip", "0.0.0.0", "App listening ip.")
	LogPath := flag.String("log", "./logs/", "Path for log files.")
	LogLevel := flag.Int("level", 1, "Logging level.")
	DbName := flag.String("db", "app-api.db", "Set name for database.")

	flag.Parse()
	addr := Adress{
		IP:   []byte(*ip),
		Port: *port,
	}
	log := Logging{
		LogPath:  *LogPath,
		LogLevel: *LogLevel,
	}
	return Config{
		Adress:  addr,
		Logging: log,
		DbName:  *DbName,
	}
}
