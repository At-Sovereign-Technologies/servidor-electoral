package models

type VotingBooth struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"size:255;not null" validate:"required,max=255"`
	Secret string `gorm:"size:32;not null" validate:"required,len=32"`

	VotingPlaceID uint        `gorm:"not null;index" validate:"required"`
	VotingPlace   VotingPlace `gorm:"foreignKey:VotingPlaceID;constraint:OnDelete:CASCADE" validate:"-"`
}

func (vb *VotingBooth) Validate() error {
	return validate.Struct(vb)
}
