package models

import "time"

const (
	ElectionStatusPublished = "published"
	ElectionStatusDraft     = "draft"
	ElectionStatusOngoing   = "ongoing"
	ElectionStatusArchived  = "archived"
	ElectionStatusCompleted = "completed"
)

type Election struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:255;not null" validate:"required,max=255"`
	Status    string    `gorm:"size:50;not null" validate:"required,oneof=published draft ongoing archived completed"`
	StartDate time.Time `gorm:"not null" validate:"required"`
	EndDate   time.Time `gorm:"not null" validate:"required,gtfield=StartDate"`
	Secret    string    `gorm:"size:32;not null" validate:"required,len=32"`

	Candidates   []Candidate         `gorm:"constraint:OnDelete:CASCADE" validate:"-"`
	Voters       []Voter             `gorm:"constraint:OnDelete:CASCADE" validate:"-"`
	VotingPlaces []VotingPlace       `gorm:"constraint:OnDelete:CASCADE" validate:"-"`
	Deployment   *ElectionDeployment `gorm:"constraint:OnDelete:CASCADE" validate:"-"`
}

func (e *Election) Validate() error {
	return validate.Struct(e)
}
