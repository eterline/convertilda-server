package api

type file struct {
	Target   string `reqHeader:"target"`
	Bitrate  int    `reqHeader:"bitrate"`
	Quallity string `reqHeader:"quallity"`
}
