package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/interface/http/dto"
)

type IProposalService interface {
	CreateProposal(ctx context.Context, req *dto.CreateProposalRequest) error
	GetProposals(ctx context.Context,
		pageReq *dto.PaginationRequest) ([]*dto.ProposalResponse, *dto.PaginationResponse, error)
	GetProposalDetail(ctx context.Context, id uuid.UUID) (*dto.ProposalResponse, error)
	CreateReply(ctx context.Context, req *dto.CreateReplyRequest) error
}
