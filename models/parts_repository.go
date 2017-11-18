package models

var PartsRepository = partsRepository{}

type partsRepository struct {
}

func (p *partsRepository) RegisterPart(fileName string, partsTypeID int) (*OriginalParts) {
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

func (p *partsRepository) FindAll() ([]OriginalParts) {
	originalParts := []OriginalParts{}
	err := masterDB.Find(&originalParts).Error()

	if err != nil {
		logger.Errorln(err)
	}
	return originalParts
}
