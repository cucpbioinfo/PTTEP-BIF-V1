package models

import (
	"time"

	"github.com/google/uuid"
)

type Method struct {
	tableName  struct{}  `pg:"method"`
	MethodID   uuid.UUID `pg:"method_id,pk"`
	MethodName string    `pg:"method_name"`

	IdentificationID uuid.UUID       `pg:"identification_id"`
	Identification   *Identification `pg:"rel:has-one"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
