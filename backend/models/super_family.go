package models

import (
	"github.com/google/uuid"
	"time"
)

type SuperFamily struct {
	tableName       struct{}  `pg:"super_family"`
	SuperFamilyID   uuid.UUID `pg:"super_family_id,pk"`
	SuperFamilyName string    `pg:"super_family_name"`

	SubOrderID uuid.UUID `pg:"sub_order_id"`
	SubOrder   *SubOrder `pg:"rel:has-one"`

	OrderID uuid.UUID `pg:"order_id"`
	Order   *Order    `pg:"rel:has-one"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
