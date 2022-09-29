package models

import (
	"time"

	"github.com/google/uuid"
)

type FieldLocation struct {
	tableName         struct{}  `pg:"field_location"`
	FieldLocationID   uuid.UUID `pg:"field_location_id,pk"`
	FieldLocationName string    `pg:"field_location_name"`
	FieldLocationLat  string    `pg:"field_location_lat"`
	FieldLocationLong string    `pg:"field_location_long"`
	Replication       int32     `pg:"replication"`
	MainStationID     uuid.UUID `pg:"main_station_id"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
