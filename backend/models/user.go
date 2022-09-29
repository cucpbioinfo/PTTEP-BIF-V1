package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	tableName struct{}  `pg:"user"`
	UserID    uuid.UUID `pg:"user_id,pk"`
	Email     string    `pg:"email"`
	Password  string    `pg:"password"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
