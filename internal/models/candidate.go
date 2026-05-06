package models

const (
	CandidateStatusDraft     = "draft"
	CandidateStatusSubmitted = "submitted"
	CandidateStatusInReview  = "in_review"
	CandidateStatusApproved  = "approved"
	CandidateStatusPublished = "published"
	CandidateStatusRejected  = "rejected"
	CandidateStatusBlocked   = "blocked"
	CandidateStatusReplaced  = "replaced"
	CandidateStatusRevoked   = "revoked"
)

type Candidate struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"size:255;not null" validate:"required,max=255"`
	Document   string `gorm:"size:100;not null;index:,unique,composite:idx_candidate_doc" validate:"required"`
	Party      string `gorm:"size:255;not null" validate:"required"`
	Location   string `gorm:"size:255;not null" validate:"required"`
	PictureURL string `gorm:"size:500;not null" validate:"required,url"`

	Status string `gorm:"size:50;not null" validate:"required,oneof=draft submitted in_review approved published rejected blocked replaced revoked"`

	ElectionID uint     `gorm:"not null;index:,unique,composite:idx_candidate_doc" validate:"required"`
	Election   Election `gorm:"foreignKey:ElectionID;constraint:OnDelete:CASCADE" validate:"-"`
}

func (c *Candidate) Validate() error {
	return validate.Struct(c)
}
