package models

type VotingPlace struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"size:255;not null" validate:"required,max=255"`
	Address string `gorm:"size:255;not null" validate:"required"`
	Secret  string `gorm:"size:32;not null" validate:"required,len=32"`

	ElectionID uint     `gorm:"not null;index" validate:"required"`
	Election   Election `gorm:"foreignKey:ElectionID;constraint:OnDelete:CASCADE" validate:"-"`

	VotingBooths []VotingBooth `gorm:"constraint:OnDelete:CASCADE" validate:"-"`
	Voters       []Voter       `gorm:"constraint:OnDelete:CASCADE" validate:"-"`
}

func (vp *VotingPlace) Validate() error {
	return validate.Struct(vp)
}
