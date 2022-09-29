package models

import (
	"time"

	"github.com/google/uuid"
)

type Density struct {
	tableName   struct{}  `pg:"density"`
	DensityID   uuid.UUID `pg:"density_id,pk"`
	Density     string    `pg:"density"`
	DensityYear string    `pg:"density_year"`

	AssetID    uuid.UUID `pg:"project_id"`
	Asset      *Project  `pg:"rel:has-one"`
	PlatformID uuid.UUID `pg:"platform_id"`
	Platform   *Platform `pg:"rel:has-one"`
	StationID  uuid.UUID `pg:"station_id"`
	Station    *Station  `pg:"rel:has-one"`

	IdentificationID uuid.UUID       `pg:"identification_id"`
	Identification   *Identification `pg:"rel:has-one"`
	MethodID         uuid.UUID       `pg:"method_id"`
	Method           *Method         `pg:"rel:has-one"`

	// MajorGroupID uuid.UUID   `pg:"major_group_id"`
	// MajorGroup   *MajorGroup `pg:"rel:has-one"`
	// KingdomID    uuid.UUID   `pg:"kingdom_id"`
	// Kingdom      *Kingdom    `pg:"rel:has-one"`
	// PhylumID     uuid.UUID   `pg:"phylum_id"`
	// Phylum       *Phylum     `pg:"rel:has-one"`
	// ClassID      uuid.UUID   `pg:"class_id"`
	// Class        *Class      `pg:"rel:has-one"`
	// OrderID      uuid.UUID   `pg:"order_id"`
	// Order        *Order      `pg:"rel:has-one"`
	// FamilyID     uuid.UUID   `pg:"family_id"`
	// Family       *Family     `pg:"rel:has-one"`
	// GenusID      uuid.UUID   `pg:"genus_id"`
	// Genus        *Genus      `pg:"rel:has-one"`
	SpeciesID uuid.UUID `pg:"species_id"`
	Species   *Species  `pg:"rel:has-one"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
