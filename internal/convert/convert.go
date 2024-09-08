package convert

import (
	"os/exec"
	"strings"

	"golang.org/x/exp/rand"
)

func New(src string, out string, typ string, targ string) *SavedFile {
	return &SavedFile{
		Name:   src,
		Output: out,
		Type:   typ,
		Target: targ,
	}
}

func (f *SavedFile) Convert(down, save string) error {
	types := GetTypes()
	switch f.Type {
	case types.Document:
		err := convertOffice(down, f.Output, f.Target, save)
		if err != nil {
			return err
		}
	case types.Audio:
		err := convertAudio(down, f.Output, f.Target, save)
		if err != nil {
			return err
		}
	}
	return nil
}

func RandStringBytesRmndr(n int) string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func convertOffice(store string, input string, target string, output string) error {
	cmd := exec.Command("soffice", "--headless", "--convert-to", target, store+input, "--outdir", output)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func convertAudio(store string, input string, target string, output string) error {
	out := strings.Split(input, ".")[0] + "." + target
	cmd := exec.Command("ffmpeg", "-i", store+input, output+out)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
