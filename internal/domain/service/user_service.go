package service

import (
	"context"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/interface/http/dto"
)

type IUserService interface {
	Login(ctx context.Context, email, password string) (string, *dto.UserResponse, error)
}
