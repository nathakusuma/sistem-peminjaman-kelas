package proposal

import (
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/domain/repository"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/domain/service"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/mail"
)

type proposalService struct {
	repo   repository.IProposalRepository
	mailer mail.IMailer
}

func NewProposalService(
	repo repository.IProposalRepository,
	mailer mail.IMailer,
) service.IProposalService {
	return &proposalService{
		repo:   repo,
		mailer: mailer,
	}
}
