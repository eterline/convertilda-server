package database

import "time"

type ConvertFiles struct {
	ID         uint `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	CreatedAt  time.Time
	SourceFile string `json:"source_file"`
	OutFile    string `json:"outpit_file"`
	Target     string `json:"target_ext"`
	FileType   string `json:"type"`
}
