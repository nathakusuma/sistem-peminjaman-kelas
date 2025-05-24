package service

import (
	"context"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/interface/http/dto"
)

type IUserService interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (string, *dto.UserResponse, error)
	Login(ctx context.Context, req *dto.LoginRequest) (string, *dto.UserResponse, error)
}
