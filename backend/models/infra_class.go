package models

import (
	"time"

	"github.com/google/uuid"
)

type InfraClass struct {
	tableName      struct{}  `pg:"infra_class"`
	InfraClassID   uuid.UUID `pg:"infra_class_id"`
	InfraClassName string    `pg:"infra_class_name"`

	ClassID uuid.UUID `pg:"class_id"`
	Class   *Class    `pg:"rel:has-one"`

	SubClassID uuid.UUID `pg:"sub_class_id"`
	SubClass   *SubClass `pg:"rel:has-one"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
