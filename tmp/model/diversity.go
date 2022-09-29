package models

import (
	"time"

	"github.com/google/uuid"
)

type Diversity struct {
	tableName   struct{}  `pg:"diversity"`
	DiversityID uuid.UUID `pg:"diversity_id,pk"`
	Year        string    `pg:"year"`
	Diversity   string    `pg:"diversity"`

	ProjectID  uuid.UUID `pg:"project_id"`
	Project    *Project  `pg:"rel:has-one"`
	PlatformID uuid.UUID `pg:"platform_id"`
	Platform   *Platform `pg:"rel:has-one"`
	StationID  uuid.UUID `pg:"station_id"`
	Station    *Station  `pg:"rel:has-one"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
