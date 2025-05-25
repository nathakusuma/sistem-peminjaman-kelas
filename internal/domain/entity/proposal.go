package entity

import (
	"time"

	"github.com/google/uuid"
)

type Proposal struct {
	ID            uuid.UUID `db:"id"`
	ProposerEmail string    `db:"proposer_email"`
	Purpose       string    `db:"purpose"`
	Course        string    `db:"course"`
	ClassID       string    `db:"class_id"`
	Lecturer      string    `db:"lecturer"`
	StartsAt      time.Time `db:"starts_at"`
	EndsAt        time.Time `db:"ends_at"`
	Occupancy     int       `db:"occupancy"`
	Note          *string   `db:"note"`
	CreatedAt     time.Time `db:"created_at"`

	Proposer *User  `db:"proposer"`
	Reply    *Reply `db:"reply"`
}

func (p *Proposal) GetStatus() string {
	if p.Reply == nil {
		return "pending"
	}
	if p.Reply.IsApproved {
		return "approved"
	}
	return "rejected"
}
