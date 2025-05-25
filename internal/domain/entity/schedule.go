package entity

import (
	"time"

	"github.com/google/uuid"
)

type Schedule struct {
	ID           uuid.UUID `db:"id"`
	Day          string    `db:"day"`
	StartTime    time.Time `db:"start_time"`
	FinishTime   time.Time `db:"finish_time"`
	RoomID       string    `db:"room_id"`
	Course       string    `db:"course"`
	ClassID      string    `db:"class_id"`
	IsLaboratory bool      `db:"is_laboratory"`
	Lecturer     string    `db:"lecturer"`
	Major        string    `db:"major"`
	StartDate    time.Time `db:"start_date"`
	FinishDate   time.Time `db:"finish_date"`
}
