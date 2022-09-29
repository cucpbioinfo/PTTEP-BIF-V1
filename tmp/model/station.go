package models

import (
	"time"

	"github.com/google/uuid"
)

type Station struct {
	tableName   struct{}  `pg:"station"`
	StationID   uuid.UUID `pg:"station_id,pk"`
	StationName string    `pg:"station_name"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
