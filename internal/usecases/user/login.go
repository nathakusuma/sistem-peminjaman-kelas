package user

import (
	"context"

	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/interface/http/dto"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/errorpkg"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/log"
)

func (s *userService) Login(ctx context.Context, req *dto.LoginRequest) (string, *dto.UserResponse, error) {
	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		if err.Error() == "user not found" {
			return "", nil, errorpkg.ErrCredentialsNotMatch()
		}

		log.Error(ctx).Err(err).Msg("failed to get user by id")
		return "", nil, errorpkg.ErrInternalServer()
	}

	if !s.bcrypt.Compare(req.Password, user.PasswordHash) {
		return "", nil, errorpkg.ErrCredentialsNotMatch()
	}

	token, err := s.jwt.Create(user.Email, user.Role)
	if err != nil {
		log.Error(ctx).Err(err).Msg("failed to create jwt token")
		return "", nil, errorpkg.ErrInternalServer()
	}

	userResp := &dto.UserResponse{
		Email: user.Email,
		Name:  user.Name,
		Role:  user.Role,
	}

	log.Info(ctx).Str("user.email", user.Email).Msg("user logged in successfully")
	return token, userResp, nil
}
