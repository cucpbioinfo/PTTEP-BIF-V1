package models

import (
	"time"

	"github.com/google/uuid"
)

type PlatformDensity struct {
	tableName struct{}  `pg:"platformdensity"`
	DensityID uuid.UUID `pg:"density_id,pk"`
	Year      string    `pg:"year"`

	AssetID     uuid.UUID `pg:"asset_id"`
	Asset       *Asset    `pg:"rel:has-one"`
	PlatformID  uuid.UUID `pg:"platform_id"`
	Platform    *Platform `pg:"rel:has-one"`
	SpeciesID   uuid.UUID `pg:"species_id"`
	SpeciesName string    `pg:"species_name"`

	Surface       string `pg:"surface"`
	Euphotic_zone string `pg:"euphotic_zone"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}