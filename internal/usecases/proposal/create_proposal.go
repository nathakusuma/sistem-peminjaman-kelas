package proposal

import (
	"context"

	"github.com/google/uuid"

	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/domain/enum"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/interface/http/dto"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/errorpkg"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/log"
)

func (s *proposalService) CreateProposal(ctx context.Context, req *dto.CreateProposalRequest) error {
	userRole := ctx.Value("user.role").(enum.UserRole)
	if userRole != enum.UserRoleStudent {
		return errorpkg.ErrForbiddenRole()
	}

	id, err := uuid.NewV7()
	if err != nil {
		log.Error(ctx).Err(err).Msg("failed to generate UUID for proposal")
		return errorpkg.ErrInternalServer()
	}

	userEmail := ctx.Value("user.email").(string)

	proposal := req.ToEntity(id, userEmail)

	err = s.repo.CreateProposal(ctx, proposal)
	if err != nil {
		log.Error(ctx).Err(err).Msg("failed to create proposal")
		return errorpkg.ErrInternalServer()
	}

	log.Info(ctx).Any("proposal", proposal).Msg("proposal created successfully")
	return nil
}
