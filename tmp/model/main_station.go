package models

import (
	"time"

	"github.com/google/uuid"
)

type MainStation struct {
	tableName       struct{}  `pg:"main_station"`
	MainStationID   uuid.UUID `pg:"main_station_id,pk"`
	MainStationName string    `pg:"main_station_name"`
	RadiusFromSite  string    `pg:"radius_from_site"`
	Easting         int32     `pg:"easting"`
	Northing        int32     `pg:"northing"`
	MainStationLat  string    `pg:"main_station_lat"`
	MainStationLong string    `pg:"main_station_long"`
	SiteID          uuid.UUID `pg:"site_id"`

	FieldLocations []*FieldLocation `pg:"rel:has-many"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
