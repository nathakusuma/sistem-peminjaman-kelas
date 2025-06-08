package proposal

import (
	"context"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/domain/ctxkey"
	"strings"

	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/interface/http/dto"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/errorpkg"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/log"
)

func (s *proposalService) CreateReply(ctx context.Context, req *dto.CreateReplyRequest) error {
	adminEmail := ctx.Value(ctxkey.UserEmail).(string)

	reply := req.ToEntity(adminEmail)

	err := s.repo.CreateReply(ctx, reply)
	if err != nil {
		if strings.HasPrefix(err.Error(), "proposal not found") {
			return errorpkg.ErrNotFound().WithDetail("Proposal not found for the given ID")
		}
		if strings.HasPrefix(err.Error(), "reply already exists") {
			return errorpkg.ErrReplyAlreadyExists()
		}
		log.Error(ctx).Err(err).Msg("failed to create reply")
		return errorpkg.ErrInternalServer()
	}

	log.Info(ctx).Any("reply", reply).Msg("reply created successfully")
	return nil
}
