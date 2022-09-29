package models

import (
	"time"

	"github.com/google/uuid"
)

type Platform struct {
	tableName    struct{}  `pg:"platform"`
	PlatformID   uuid.UUID `pg:"platform_id,pk"`
	PlatformName string    `pg:"platform_name"`

	AssetID uuid.UUID `pg:"asset_id"`
	Asset   *Asset    `pg:"rel:has-one"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
