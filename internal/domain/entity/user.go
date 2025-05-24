package entity

import "github.com/nathakusuma/sistem-peminjaman-kelas/internal/domain/enum"

type User struct {
	Email        string        `db:"email"`
	PasswordHash string        `db:"password_hash"`
	Name         string        `db:"name"`
	Role         enum.UserRole `db:"role"`
}
