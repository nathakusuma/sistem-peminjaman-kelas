package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/domain/entity"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/domain/repository"
)

type userRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) repository.IUserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(user *entity.User) error {
	query := `
        INSERT INTO users (email, password_hash, name, role) 
        VALUES ($1, $2, $3, $4)
    `

	_, err := r.db.Exec(
		context.Background(),
		query,
		user.Email,
		user.PasswordHash,
		user.Name,
		user.Role,
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok && pgErr.Code == "23505" {
			return errors.New("user already exists")
		}

		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (r *userRepository) GetUserByEmail(id string) (*entity.User, error) {
	query := `
        SELECT email, password_hash, name, role 
        FROM users 
        WHERE email = $1
    `

	rows, err := r.db.Query(context.Background(), query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to query user: %w", err)
	}
	defer rows.Close()

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[entity.User])
	if err != nil {
		return nil, fmt.Errorf("failed to scan user: %w", err)
	}

	if len(users) == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return &users[0], nil
}
