package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/domain/entity"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/interface/http/dto"
)

type IProposalRepository interface {
	CreateProposal(ctx context.Context, proposal *entity.Proposal) error
	GetProposals(ctx context.Context,
		pageReq *dto.PaginationRequest) ([]*entity.Proposal, int64, error)
	GetProposalsByUser(ctx context.Context, userEmail string,
		pageReq *dto.PaginationRequest) ([]*entity.Proposal, int64, error)
	GetProposalDetail(ctx context.Context, id uuid.UUID) (*entity.Proposal, error)
	CreateReply(ctx context.Context, reply *entity.Reply) error
}
