package user

import (
	"context"

	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/domain/entity"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/domain/enum"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/interface/http/dto"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/errorpkg"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/log"
)

func (s *userService) Register(ctx context.Context, req *dto.RegisterRequest) (string, *dto.UserResponse, error) {
	passwordHash, err := s.bcrypt.Hash(req.Password)
	if err != nil {
		log.Error(ctx).Err(err).Msg("failed to hash password")
		return "", nil, errorpkg.ErrInternalServer()
	}

	newUser := &entity.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: passwordHash,
		Role:         enum.UserRoleStudent,
	}

	if err := s.repo.CreateUser(newUser); err != nil {
		if err.Error() == "user already exists" {
			log.Warn(ctx).Str("user.email", req.Email).Msg("user already exists")
			return "", nil, errorpkg.ErrEmailAlreadyRegistered()
		}

		log.Error(ctx).Err(err).Msg("failed to create user")
		return "", nil, errorpkg.ErrInternalServer()
	}

	log.Info(ctx).Str("user.email", newUser.Email).Msg("user registered successfully")

	loginReq := &dto.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	return s.Login(ctx, loginReq)
}
