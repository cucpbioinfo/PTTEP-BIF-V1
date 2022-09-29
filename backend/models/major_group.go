package models

import (
	"github.com/google/uuid"
	"time"
)

type MajorGroup struct {
	tableName      struct{}  `pg:"major_group"`
	MajorGroupID   uuid.UUID `pg:"major_group_id,pk"`
	MajorGroupName string    `pg:"major_group_name"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
