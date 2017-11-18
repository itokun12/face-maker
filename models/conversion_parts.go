package models

import "github.com/jinzhu/gorm"

type ConversionParts struct {
	gorm.Model
	OriginalPartsID int
	FileName        string
}
