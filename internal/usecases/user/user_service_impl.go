package user

import (
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/domain/repository"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/domain/service"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/bcrypt"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/jwt"
)

type userService struct {
	repo   repository.IUserRepository
	bcrypt bcrypt.IBcrypt
	jwt    jwt.IJwt
}

func NewUserService(
	userRepo repository.IUserRepository,
	bcrypt bcrypt.IBcrypt,
	jwt jwt.IJwt,
) service.IUserService {
	return &userService{
		repo:   userRepo,
		bcrypt: bcrypt,
		jwt:    jwt,
	}
}
