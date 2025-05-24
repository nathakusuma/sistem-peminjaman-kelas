package dto

import "github.com/nathakusuma/sistem-peminjaman-kelas/internal/domain/enum"

type UserResponse struct {
	Email string        `json:"email"`
	Name  string        `json:"name"`
	Role  enum.UserRole `json:"role"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email,max=320"`
	Password string `json:"password" validate:"required,min=8,max=72"`
}

type LoginResponse struct {
	Token string        `json:"token"`
	User  *UserResponse `json:"user"`
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email,max=320"`
	Password string `json:"password" validate:"required,min=8,max=72"`
	Name     string `json:"name" validate:"required,max=255"`
}
