package proposal

import (
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/domain/repository"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/domain/service"
)

type proposalService struct {
	repo repository.IProposalRepository
}

func NewProposalService(repo repository.IProposalRepository) service.IProposalService {
	return &proposalService{
		repo: repo,
	}
}
