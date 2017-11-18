package models

var PartsRepository = partsRepository{}

type partsRepository struct {
}

func (d *partsRepository) RegisterPart(fileName string, partsTypeID int) (*OriginalParts) {
	part := OriginalParts{
		Filename: fileName,
		PartsTypeID: partsTypeID,
	}
	err := masterDB.Create(&part).Error()
	if err != nil {
		logger.Errorln(err)
	}
	return &part
}
