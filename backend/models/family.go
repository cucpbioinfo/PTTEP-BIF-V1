package models

import (
	"github.com/google/uuid"
	"time"
)

type Family struct {
	tableName  struct{}  `pg:"family"`
	FamilyID   uuid.UUID `pg:"family_id,pk"`
	FamilyName string    `pg:"family_name"`

	SuperFamilyID uuid.UUID    `pg:"super_family_id"`
	SuperFamily   *SuperFamily `pg:"rel:has-one"`

	OrderID uuid.UUID `pg:"order_id"`
	Order   *Order    `pg:"rel:has-one"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
