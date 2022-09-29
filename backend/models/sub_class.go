package models

import (
	"github.com/google/uuid"
	"time"
)

type SubClass struct {
	tableName    struct{}  `pg:"sub_class"`
	SubClassID   uuid.UUID `pg:"sub_class_id,pk"`
	SubClassName string    `pg:"sub_class_name"`

	ClassID uuid.UUID `pg:"class_id"`
	Class   *Class    `pg:"rel:has-one"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
