package models

import (
	"time"

	"github.com/google/uuid"
)

type InfraKingdom struct {
	tableName        struct{}  `pg:"infra_kingdom"`
	InfraKingdomID   uuid.UUID `pg:"infra_kingdom_id"`
	InfraKingdomName string    `pg:"infra_kingdom_name"`

	KingdomID uuid.UUID `pg:"kingdom_id"`
	Kingdom   *Kingdom  `pg:"rel:has-one"`

	SubKingdomID uuid.UUID   `pg:"sub_kingdom_id"`
	SubKingdom   *SubKingdom `pg:"rel:has-one"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
