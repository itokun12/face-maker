package models

import "github.com/jinzhu/gorm"

type OriginalParts struct {
	gorm.Model
	PartsTypeID int
	FilePath    string
}

const (
	Eyebrow = 1
	Eye     = 2
	Nose    = 3
	Mouth   = 4
)
