package models

import (
	"github.com/google/uuid"
	"time"
)

type SuperOrder struct {
	tableName      struct{}  `pg:"super_order"`
	SuperOrderID   uuid.UUID `pg:"super_order_id,pk"`
	SuperOrderName string    `pg:"super_order_name"`

	SubClassID uuid.UUID `pg:"sub_class_id"`
	SubClass   *SubClass `pg:"rel:has-one"`

	ClassID uuid.UUID `pg:"class_id"`
	Class   *Class    `pg:"rel:has-one"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
