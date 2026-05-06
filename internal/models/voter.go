package models

type Voter struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:255;not null" validate:"required,max=255"`
	Document string `gorm:"size:100;not null;index:,unique,composite:idx_voter_doc" validate:"required"`

	ElectionID uint     `gorm:"not null;index:,unique,composite:idx_voter_doc" validate:"required"`
	Election   Election `gorm:"foreignKey:ElectionID;constraint:OnDelete:CASCADE" validate:"-"`

	VotingPlaceID uint        `gorm:"not null;index" validate:"required"`
	VotingPlace   VotingPlace `gorm:"foreignKey:VotingPlaceID;constraint:OnDelete:CASCADE" validate:"-"`
}

func (v *Voter) Validate() error {
	return validate.Struct(v)
}
