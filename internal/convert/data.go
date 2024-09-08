package convert

type FileTypes struct {
	Document string
	Audio    string
	Photo    string
}

func GetTypes() FileTypes {
	return FileTypes{
		Document: "documents",
		Audio:    "audio",
		Photo:    "photo",
	}
}

const RAND_NAME_LEN = 16

type SavedFile struct {
	Name   string
	Output string
	Type   string
	Target string
}

func DocExtensions() []string {
	return []string{"pdf", "doc", "epub", "docx", "rtf", "tiff"}
}

func AudioExtensions() []string {
	return []string{"wav", "mp3", "ogg", "flac"}
}
