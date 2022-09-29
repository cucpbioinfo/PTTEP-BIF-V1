package models

import (
	"time"

	"github.com/google/uuid"
)

type InfraOrder struct {
	tableName      struct{}  `pg:"infra_order"`
	InfraOrderID   uuid.UUID `pg:"infra_order_id"`
	InfraOrderName string    `pg:"infra_order_name"`

	OrderID uuid.UUID `pg:"order_id"`
	Order   *Order    `pg:"rel:has-one"`

	SubOrderID uuid.UUID `pg:"sub_order_id"`
	SubOrder   *SubOrder `pg:"rel:has-one"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
