package models

import (
	"time"

	"github.com/google/uuid"
)

type Asset struct {
	tableName struct{}  `pg:"asset"`
	AssetID   uuid.UUID `pg:"asset_id,pk"`
	AssetName string    `pg:"asset_name"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
