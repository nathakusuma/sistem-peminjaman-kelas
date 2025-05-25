package dto

import (
	"time"

	"github.com/google/uuid"

	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/domain/entity"
)

type ProposalResponse struct {
	ID           uuid.UUID `json:"id"`
	Purpose      string    `json:"purpose"`
	ProposerName string    `json:"proposer_name"`
	Status       string    `json:"status"`
	Course       string    `json:"course,omitempty"`
	ClassID      string    `json:"class_id,omitempty"`
	Lecturer     string    `json:"lecturer,omitempty"`
	StartsAt     time.Time `json:"starts_at,omitempty"`
	EndsAt       time.Time `json:"ends_at,omitempty"`
	Occupancy    int       `json:"occupancy,omitempty"`
	Note         *string   `json:"note,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`

	Reply *ReplyResponse `json:"reply,omitempty"`
}

type ReplyResponse struct {
	AdminName  string    `json:"admin_name"`
	RoomID     string    `json:"room_id"`
	IsApproved bool      `json:"is_approved"`
	Note       *string   `json:"note,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}

func (r *ProposalResponse) FromEntityMinimal(e *entity.Proposal) {
	r.ID = e.ID
	r.Purpose = e.Purpose
	r.ProposerName = e.Proposer.Name
	r.Status = e.GetStatus()
}

func (r *ProposalResponse) FromEntityDetail(e *entity.Proposal) {
	r.ID = e.ID
	r.Purpose = e.Purpose
	r.ProposerName = e.Proposer.Name
	r.Status = e.GetStatus()
	r.Course = e.Course
	r.ClassID = e.ClassID
	r.Lecturer = e.Lecturer
	r.StartsAt = e.StartsAt
	r.EndsAt = e.EndsAt
	r.Occupancy = e.Occupancy
	r.Note = e.Note
	r.CreatedAt = e.CreatedAt

	if e.Reply != nil {
		r.Reply = &ReplyResponse{
			AdminName:  e.Reply.Admin.Name,
			RoomID:     e.Reply.RoomID,
			IsApproved: e.Reply.IsApproved,
			Note:       e.Reply.Note,
			CreatedAt:  e.Reply.CreatedAt,
		}
	}
}

type CreateProposalRequest struct {
	Purpose   string    `json:"purpose" validate:"required,max=50"`
	Course    string    `json:"course" validate:"required,max=50"`
	ClassID   string    `json:"class_id" validate:"required,max=3"`
	Lecturer  string    `json:"lecturer" validate:"required,max=255"`
	StartsAt  time.Time `json:"starts_at" validate:"required,gt=now"`
	EndsAt    time.Time `json:"ends_at" validate:"required,gtfield=StartsAt"`
	Occupancy int       `json:"occupancy" validate:"required,max=32767"`
	Note      *string   `json:"note" validate:"omitempty,max=1000"`
}

func (req *CreateProposalRequest) ToEntity(id uuid.UUID, userEmail string) *entity.Proposal {
	return &entity.Proposal{
		ID:            id,
		ProposerEmail: userEmail,
		Purpose:       req.Purpose,
		Course:        req.Course,
		ClassID:       req.ClassID,
		Lecturer:      req.Lecturer,
		StartsAt:      req.StartsAt,
		EndsAt:        req.EndsAt,
		Occupancy:     req.Occupancy,
		Note:          req.Note,
		CreatedAt:     time.Now(),
	}
}

type CreateReplyRequest struct {
	ProposalID uuid.UUID `param:"proposal_id" validate:"required"`
	RoomID     string    `json:"room_id" validate:"required,max=20"`
	IsApproved bool      `json:"is_approved" validate:"required"`
	Note       *string   `json:"note" validate:"omitempty,max=1000"`
}

func (req *CreateReplyRequest) ToEntity(adminEmail string) *entity.Reply {
	return &entity.Reply{
		ID:         req.ProposalID,
		AdminEmail: adminEmail,
		RoomID:     req.RoomID,
		IsApproved: req.IsApproved,
		Note:       req.Note,
		CreatedAt:  time.Now(),
	}
}
