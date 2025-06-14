package proposal

import (
	"context"
	"strings"

	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/domain/ctxkey"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/domain/enum"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/infrastructure/config"
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

	proposal, err := s.repo.GetProposalDetail(ctx, req.ProposalID)
	if err != nil {
		log.Error(ctx).Err(err).Msg("failed to get proposal by ID")
		return errorpkg.ErrInternalServer()
	}

	if config.GetEnv().Env == enum.EnvProduction {
		// Send email notification to the user
		go func() {
			err = s.mailer.Send(
				proposal.ProposerEmail,
				"Konfirmasi Pengajuan Kelas",
				"reply_student_notification.html",
				map[string]any{
					"name":        proposal.Proposer.Name,
					"is_approved": reply.IsApproved,
				},
			)
			if err != nil {
				log.Error(ctx).Err(err).Msg("failed to send email notification")
			}
		}()
	}

	log.Info(ctx).Any("reply", reply).Msg("reply created successfully")
	return nil
}
