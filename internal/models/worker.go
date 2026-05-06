package models

type ElectionWorker struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:255;not null" validate:"required,max=255"`
	Revoked   bool   `gorm:"not null" validate:"-"`
	CreatedAt int64  `gorm:"autoCreateTime" validate:"-"`

	ElectionID uint     `gorm:"not null;index" validate:"required"`
	Election   Election `gorm:"foreignKey:ElectionID;constraint:OnDelete:CASCADE" validate:"-"`
}

func (w *ElectionWorker) Validate() error {
	return validate.Struct(w)
}
