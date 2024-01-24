package models

import (
	"time"

)

type Todo struct {
	ID          int       `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	DueDate     time.Time `db:"due_date"`
	IsCompleted bool      `db:"is_completed"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
