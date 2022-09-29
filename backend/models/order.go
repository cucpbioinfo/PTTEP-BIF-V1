package models

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	tableName struct{}  `pg:"order"`
	OrderID   uuid.UUID `pg:"order_id,pk"`
	OrderName string    `pg:"order_name"`

	SuperOrderID uuid.UUID   `pg:"super_order_id"`
	SuperOrder   *SuperOrder `pg:"rel:has-one"`

	ClassID uuid.UUID `pg:"class_id"`
	Class   *Class    `pg:"rel:has-one"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
