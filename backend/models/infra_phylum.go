package models

import (
	"time"

	"github.com/google/uuid"
)

type InfraPhylum struct {
	tableName       struct{}  `pg:"infra_phylum"`
	InfraPhylumID   uuid.UUID `pg:"infra_phylum_id"`
	InfraPhylumName string    `pg:"infra_phylum_name"`

	PhylumID uuid.UUID `pg:"phylum_id"`
	Phylum   *Phylum   `pg:"rel:has-one"`

	SubPhylumID uuid.UUID  `pg:"sub_phylum_id"`
	SubPhylum   *SubPhylum `pg:"rel:has-one"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
