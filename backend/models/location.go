package models

import (
	"time"

	"github.com/google/uuid"
)

type Location struct {
	tableName  struct{}  `pg:"location"`
	LocationID uuid.UUID `pg:"location_id,pk"`
	Latitude   string    `pg:"latitude"`
	Longitude  string    `pg:"longitude"`

	AssetID    uuid.UUID `pg:"asset_id"`
	Asset      *Asset    `pg:"rel:has-one"`
	PlatformID uuid.UUID `pg:"platform_id"`
	Platform   *Platform `pg:"rel:has-one"`
	StationID  uuid.UUID `pg:"station_id"`
	Station    *Station  `pg:"rel:has-one"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}

type AssetLocation struct {
	tableName  struct{}  `pg:"assetlocation"`
	LocationID uuid.UUID `pg:"assetlocation_id,pk"`
	Latitude   string    `pg:"latitude"`
	Longitude  string    `pg:"longitude"`

	AssetID uuid.UUID `pg:"asset_id"`
	Asset   *Asset    `pg:"rel:has-one"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
