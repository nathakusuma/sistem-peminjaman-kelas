package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/domain/entity"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/domain/repository"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/interface/http/dto"
)

type proposalRepository struct {
	db *pgx.Conn
}

func NewProposalRepository(db *pgx.Conn) repository.IProposalRepository {
	return &proposalRepository{
		db: db,
	}
}

// Helper structs for scanning database results
type proposalScanMinimal struct {
	ID            uuid.UUID `db:"id"`
	Purpose       string    `db:"purpose"`
	ProposerName  string    `db:"proposer_name"`
	ReplyApproved *bool     `db:"reply_is_approved"`
}

type proposalScanDetail struct {
	// Proposal fields
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

	// Proposer fields
	ProposerName string `db:"proposer_name"`

	// Reply fields (nullable)
	ReplyID         *uuid.UUID `db:"reply_id"`
	ReplyAdminEmail *string    `db:"reply_admin_email"`
	ReplyRoomID     *string    `db:"reply_room_id"`
	ReplyApproved   *bool      `db:"reply_is_approved"`
	ReplyNote       *string    `db:"reply_note"`
	ReplyCreatedAt  *time.Time `db:"reply_created_at"`

	// Admin fields (nullable)
	ReplyAdminName *string `db:"reply_admin_name"`
}

func (r *proposalRepository) CreateProposal(ctx context.Context, proposal *entity.Proposal) error {
	query := `
		INSERT INTO proposals (id, proposer_email, purpose, course, class_id, lecturer, starts_at, ends_at, occupancy, note, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`

	_, err := r.db.Exec(ctx, query,
		proposal.ID,
		proposal.ProposerEmail,
		proposal.Purpose,
		proposal.Course,
		proposal.ClassID,
		proposal.Lecturer,
		proposal.StartsAt,
		proposal.EndsAt,
		proposal.Occupancy,
		proposal.Note,
		proposal.CreatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create proposal: %w", err)
	}

	return nil
}

// DRY helper method for both GetProposals and GetProposalsByUser
func (r *proposalRepository) getProposalsBase(ctx context.Context, whereClause string, args []interface{}, pageReq *dto.PaginationRequest) ([]*entity.Proposal, int64, error) {
	// Build count query
	countQuery := `
		SELECT COUNT(*)
		FROM proposals p
		JOIN users u ON p.proposer_email = u.email
		LEFT JOIN replies r ON p.id = r.id
	`
	if whereClause != "" {
		countQuery += " WHERE " + whereClause
	}

	// Get total count
	var totalCount int64
	err := r.db.QueryRow(ctx, countQuery, args...).Scan(&totalCount)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get total count: %w", err)
	}

	// Build main query for minimal data (used in list views)
	query := `
		SELECT 
			p.id,
			p.purpose,
			u.name as proposer_name,
			r.is_approved as reply_is_approved
		FROM proposals p
		JOIN users u ON p.proposer_email = u.email
		LEFT JOIN replies r ON p.id = r.id
	`
	if whereClause != "" {
		query += " WHERE " + whereClause
	}

	// Order by: proposals with no replies first (priority), then by id desc (latest first for UUIDv7)
	query += `
		ORDER BY 
			CASE WHEN r.id IS NULL THEN 0 ELSE 1 END,
			p.id DESC
		LIMIT $` + fmt.Sprintf("%d", len(args)+1) + ` OFFSET $` + fmt.Sprintf("%d", len(args)+2)

	// Add pagination parameters
	offset := (pageReq.Page - 1) * pageReq.Size
	paginationArgs := append(args, pageReq.Size, offset)

	// Execute query
	rows, err := r.db.Query(ctx, query, paginationArgs...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query proposals: %w", err)
	}
	defer rows.Close()

	// Scan results
	scannedProposals, err := pgx.CollectRows(rows, pgx.RowToStructByName[proposalScanMinimal])
	if err != nil {
		return nil, 0, fmt.Errorf("failed to scan proposals: %w", err)
	}

	// Convert to entity (minimal data for list views)
	proposals := make([]*entity.Proposal, len(scannedProposals))
	for i, sp := range scannedProposals {
		proposals[i] = &entity.Proposal{
			ID:      sp.ID,
			Purpose: sp.Purpose,
			Proposer: &entity.User{
				Name: sp.ProposerName,
			},
		}

		// Set reply if exists (needed for GetStatus())
		if sp.ReplyApproved != nil {
			proposals[i].Reply = &entity.Reply{
				IsApproved: *sp.ReplyApproved,
			}
		}
	}

	return proposals, totalCount, nil
}

func (r *proposalRepository) GetProposals(ctx context.Context, pageReq *dto.PaginationRequest) ([]*entity.Proposal, int64, error) {
	return r.getProposalsBase(ctx, "", []interface{}{}, pageReq)
}

func (r *proposalRepository) GetProposalsByUser(ctx context.Context, userEmail string, pageReq *dto.PaginationRequest) ([]*entity.Proposal, int64, error) {
	return r.getProposalsBase(ctx, "p.proposer_email = $1", []interface{}{userEmail}, pageReq)
}

func (r *proposalRepository) GetProposalDetail(ctx context.Context, id uuid.UUID) (*entity.Proposal, error) {
	query := `
		SELECT 
			p.id, p.proposer_email, p.purpose, p.course, p.class_id, p.lecturer,
			p.starts_at, p.ends_at, p.occupancy, p.note, p.created_at, 
			u.name as proposer_name,
			r.id as reply_id, r.admin_email as reply_admin_email, r.room_id as reply_room_id,
			r.is_approved as reply_is_approved, r.note as reply_note, r.created_at as reply_created_at,
			ua.name as reply_admin_name
		FROM proposals p
		JOIN users u ON p.proposer_email = u.email
		LEFT JOIN replies r ON p.id = r.id
		LEFT JOIN users ua ON r.admin_email = ua.email
		WHERE p.id = $1
	`

	rows, err := r.db.Query(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to query proposal detail: %w", err)
	}
	defer rows.Close()

	scannedProposals, err := pgx.CollectRows(rows, pgx.RowToStructByName[proposalScanDetail])
	if err != nil {
		return nil, fmt.Errorf("failed to scan proposal detail: %w", err)
	}

	if len(scannedProposals) == 0 {
		return nil, fmt.Errorf("proposal not found")
	}

	sp := scannedProposals[0]

	// Build complete proposal entity with all details
	proposal := &entity.Proposal{
		ID:            sp.ID,
		ProposerEmail: sp.ProposerEmail,
		Purpose:       sp.Purpose,
		Course:        sp.Course,
		ClassID:       sp.ClassID,
		Lecturer:      sp.Lecturer,
		StartsAt:      sp.StartsAt,
		EndsAt:        sp.EndsAt,
		Occupancy:     sp.Occupancy,
		Note:          sp.Note,
		CreatedAt:     sp.CreatedAt,
		Proposer: &entity.User{
			Email: sp.ProposerEmail,
			Name:  sp.ProposerName,
		},
	}

	// Set reply if exists
	if sp.ReplyID != nil {
		proposal.Reply = &entity.Reply{
			ID:         *sp.ReplyID,
			AdminEmail: *sp.ReplyAdminEmail,
			RoomID:     *sp.ReplyRoomID,
			IsApproved: *sp.ReplyApproved,
			Note:       sp.ReplyNote,
			CreatedAt:  *sp.ReplyCreatedAt,
		}

		// Set admin information if available
		if sp.ReplyAdminName != nil {
			proposal.Reply.Admin = &entity.User{
				Email: *sp.ReplyAdminEmail,
				Name:  *sp.ReplyAdminName,
			}
		}
	}

	return proposal, nil
}

func (r *proposalRepository) CreateReply(ctx context.Context, reply *entity.Reply) error {
	query := `
		INSERT INTO replies (id, admin_email, room_id, is_approved, note, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.Exec(ctx, query,
		reply.ID,
		reply.AdminEmail,
		reply.RoomID,
		reply.IsApproved,
		reply.Note,
		reply.CreatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to add reply: %w", err)
	}

	return nil
}
