package models

type OriginalParts struct {
	Model
	PartsTypeID int
	Filename    string
}

const (
	Eyebrow = 1
	Eye     = 2
	Nose    = 3
	Mouth   = 4
)
