package convert

import "golang.org/x/exp/rand"

const RAND_NAME_LEN = 16

type SavedFile struct {
	Name   string
	Output string
	Type   string
	Target string
}

func New(n string, typ string, targ string) *SavedFile {
	return &SavedFile{
		Name:   n,
		Output: randStringBytesRmndr(RAND_NAME_LEN),
		Type:   typ,
		Target: targ,
	}
}

func (f *SavedFile) Convert() (string, error) {
	switch f.Type {
	// TODO: soffice от пакета libreoffice подойдет для конвертации, надо сделать проверку на его наличие.
	case "documents":

	}
	return "", nil
}

func randStringBytesRmndr(n int) string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
