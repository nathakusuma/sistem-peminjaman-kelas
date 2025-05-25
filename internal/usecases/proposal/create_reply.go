package proposal

import (
	"context"

	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/domain/enum"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/interface/http/dto"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/errorpkg"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/log"
)

func (s *proposalService) CreateReply(ctx context.Context, req *dto.CreateReplyRequest) error {
	userRole := ctx.Value("user.role").(enum.UserRole)
	if userRole != enum.UserRoleAdmin {
		return errorpkg.ErrForbiddenRole()
	}

	userEmail := ctx.Value("user.email").(string)

	reply := req.ToEntity(userEmail)

	err := s.repo.CreateReply(ctx, reply)
	if err != nil {
		log.Error(ctx).Err(err).Msg("failed to create reply")
		return errorpkg.ErrInternalServer()
	}

	log.Info(ctx).Any("reply", reply).Msg("reply created successfully")
	return nil
}
