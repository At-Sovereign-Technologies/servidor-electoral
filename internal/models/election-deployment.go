package models

type ElectionDeployment struct {
	ID uint `gorm:"primaryKey"`

	ElectionID uint     `gorm:"not null;uniqueIndex"`
	Election   Election `gorm:"foreignKey:ElectionID;constraint:OnDelete:CASCADE" validate:"-"`

	QueryURL    string `gorm:"size:500;not null" validate:"required,url"`
	DatabaseURI string `gorm:"size:500;not null" validate:"required,url"`
	QueueURI    string `gorm:"size:500;not null" validate:"required,url"`
}

func (ed *ElectionDeployment) Validate() error {
	return validate.Struct(ed)
}
