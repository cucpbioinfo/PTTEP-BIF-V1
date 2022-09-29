package models

import (
	"time"

	"github.com/google/uuid"
)

type Evenness struct {
	tableName     struct{}  `pg:"evenness"`
	EvennessID    uuid.UUID `pg:"evenness_id,pk"`
	Year          string    `pg:"year"`
	Surface       string    `pg:"surface"`
	Euphotic_zone string    `pg:"euphotic_zone"`

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
