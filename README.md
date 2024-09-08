
# Convertilda

<hr>

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

API server for my project.
![Logo](/convertilda.png)

<hr>

## Requires

- FFmpeg
- LibreOffice
- go 1.23
CGO_ENABLED

```
sudo apt --no-install-recommends install ffmpeg libreoffice
```

## Install

To build project:
```bash
git clone 'https://github.com/eterline/convertilda-server.git'
cd ./convertilda-server
CGO_ENABLED=1 go build ./cmd/convertilda-api/...
```
Start:
```

./convertilda-api
```

## Settings

Help list:
```
# ./convertilda-api -h
```
```
Usage of ./convertilda-api:
  -db string
        Set name for database. (default "app-api.db")
  -ip string
        App listening ip. (default "0.0.0.0")
  -level int
        Logging level. (default 1)
  -log string
        Path for log files. (default "./logs/")
  -port int
        App listening port. (default 8080)
```

## API

In process...

<hr>

## License

[MIT](https://choosealicense.com/licenses/mit/)

