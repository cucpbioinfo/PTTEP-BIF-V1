package models

import (
	"github.com/google/uuid"
	"time"
)

type SubOrder struct {
	tableName    struct{}  `pg:"sub_order"`
	SubOrderID   uuid.UUID `pg:"sub_order_id,pk"`
	SubOrderName string    `pg:"sub_order_name"`

	OrderID uuid.UUID `pg:"order_id"`
	Order   *Order    `pg:"rel:has-one"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
