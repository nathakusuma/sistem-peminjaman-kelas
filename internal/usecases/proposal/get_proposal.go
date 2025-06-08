package proposal

import (
	"context"
	"github.com/google/uuid"

	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/domain/entity"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/domain/enum"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/interface/http/dto"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/errorpkg"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/log"
)

func (s *proposalService) GetProposals(ctx context.Context, pageReq *dto.PaginationRequest) ([]*dto.ProposalResponse, *dto.PaginationResponse, error) {
	var proposals []*entity.Proposal
	var totalCount int64
	var err error

	userRole := ctx.Value("user.role").(enum.UserRole)
	if userRole == enum.UserRoleAdmin {
		proposals, totalCount, err = s.repo.GetProposals(ctx, pageReq)
	} else {
		userEmail := ctx.Value("user.email").(string)
		proposals, totalCount, err = s.repo.GetProposalsByUser(ctx, userEmail, pageReq)
	}

	if err != nil {
		log.Error(ctx).Err(err).Msg("failed to get proposals")
		return nil, nil, errorpkg.ErrInternalServer()
	}

	pageResp := &dto.PaginationResponse{
		CurrentPage: pageReq.Page,
		Size:        pageReq.Size,
		TotalCount:  totalCount,
		TotalPages:  (int)(totalCount / int64(pageReq.Size)),
		HasNext:     totalCount > int64(pageReq.Page)*int64(pageReq.Size),
		HasPrev:     pageReq.Page > 1,
	}

	proposalResp := make([]*dto.ProposalResponse, len(proposals))
	for i, proposal := range proposals {
		proposalResp[i] = new(dto.ProposalResponse)
		proposalResp[i].FromEntityMinimal(proposal)
	}

	return proposalResp, pageResp, nil
}

func (s *proposalService) GetProposalDetail(ctx context.Context, id uuid.UUID) (*dto.ProposalResponse, error) {
	proposal, err := s.repo.GetProposalDetail(ctx, id)
	if err != nil {
		log.Error(ctx).Err(err).Msg("failed to get proposal detail")
		return nil, errorpkg.ErrInternalServer()
	}

	userRole := ctx.Value("user.role").(enum.UserRole)
	userEmail := ctx.Value("user.email").(string)
	if userRole != enum.UserRoleAdmin && proposal.ProposerEmail != userEmail {
		return nil, errorpkg.ErrForbiddenUser()
	}

	proposalResp := &dto.ProposalResponse{}
	proposalResp.FromEntityDetail(proposal)

	return proposalResp, nil
}
