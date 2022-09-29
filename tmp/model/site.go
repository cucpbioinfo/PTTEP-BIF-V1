package models

import (
	"time"

	"github.com/google/uuid"
)

type Site struct {
	tableName struct{}  `pg:"site"`
	SiteID    uuid.UUID `pg:"site_id,pk"`
	SiteName  string    `pg:"site_name"`

	MainStations []*MainStation `pg:"rel:has-many"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
