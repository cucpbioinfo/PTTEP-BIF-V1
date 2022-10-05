package models

import (
	"time"

	"github.com/google/uuid"
)

type Summary struct {
	tableName struct{}  `pg:"summary"`
	SummaryID uuid.UUID `pg:"summary_id,pk"`
	Year      string    `pg:"year"`

	MajorGroupID     uuid.UUID       `pg:"major_group_id"`
	MajorGroup       *MajorGroup     `pg:"rel:has-one"`
	IdentificationID uuid.UUID       `pg:"identification_id"`
	Identification   *Identification `pg:"rel:has-one"`
	AssetID          uuid.UUID       `pg:"asset_id"`
	Asset            *Asset          `pg:"rel:has-one"`
	PlatformID       uuid.UUID       `pg:"platform_id"`
	Platform         *Platform       `pg:"rel:has-one"`
	StationID        uuid.UUID       `pg:"station_id"`
	Station          *Station        `pg:"rel:has-one"`

	SurfaceShannon  string `pg:"surfaceshannon"`
	SurfaceNumber   string `pg:"surfacenumber"`
	SurfaceMax      string `pg:"surfacemax"`
	SurfaceEvenness string `pg:"surfaceevenness"`

	EuphoticZoneShannon  string `pg:"euphoticzoneshannon"`
	EuphoticZoneNumber   string `pg:"euphoticzonenumber"`
	EuphoticZoneMax      string `pg:"euphoticzonemax"`
	EuphoticZoneEvenness string `pg:"euphoticzoneevenness"`

	AverageShannon  string `pg:"averageshannon"`
	AverageNumber   string `pg:"averagenumber"`
	AverageMax      string `pg:"averagemax"`
	AverageEvenness string `pg:"averageevenness"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
