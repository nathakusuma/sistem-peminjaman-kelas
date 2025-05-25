package entity

import (
	"time"

	"github.com/google/uuid"
)

type Reply struct {
	ID         uuid.UUID `db:"id"`
	AdminEmail string    `db:"admin_email"`
	RoomID     string    `db:"room_id"`
	IsApproved bool      `db:"is_approved"`
	Note       *string   `db:"note"`
	CreatedAt  time.Time `db:"created_at"`

	Admin *User `db:"admin"`
}
