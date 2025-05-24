package repository

import "github.com/nathakusuma/sistem-peminjaman-kelas/internal/domain/entity"

type IUserRepository interface {
	CreateUser(user *entity.User) error
	GetUserByEmail(email string) (*entity.User, error)
}
