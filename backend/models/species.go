package models

import (
	"time"

	"github.com/google/uuid"
)

type Species struct {
	tableName   struct{}  `pg:"species"`
	SpeciesID   uuid.UUID `pg:"species_id,pk"`
	SpeciesName string    `pg:"species_name"`
	// ScientificName string    `pg:"scientific_name"`
	// CommonNameEn   string    `pg:"common_name_en"`
	// CommonNameTh   string    `pg:"common_name_th"`
	// Synonym        string    `pg:"synonym"`
	// DNA            string    `pg:"dna"`

	IdentificationID uuid.UUID       `pg:"identification_id"`
	Identification   *Identification `pg:"rel:has-one"`

	MethodID uuid.UUID `pg:"method_id"`
	Method   *Method   `pg:"rel:has-one"`

	MajorGroupID uuid.UUID   `pg:"major_group_id"`
	MajorGroup   *MajorGroup `pg:"rel:has-one"`

	KingdomID uuid.UUID `pg:"kingdom_id"`
	Kingdom   *Kingdom  `pg:"rel:has-one"`

	// SubKingdomID uuid.UUID   `pg:"sub_kingdom_id"`
	// SubKingdom   *SubKingdom `pg:"rel:has-one"`

	// InfraKingdomID uuid.UUID     `pg:"infra_kingdom_id"`
	// InfraKingdom   *InfraKingdom `pg:"rel:has-one"`

	// SuperPhylumID uuid.UUID    `pg:"super_phylum_id"`
	// SuperPhylum   *SuperPhylum `pg:"rel:has-one"`

	PhylumID uuid.UUID `pg:"phylum_id"`
	Phylum   *Phylum   `pg:"rel:has-one"`

	// SubPhylumID uuid.UUID  `pg:"sub_phylum_id"`
	// SubPhylum   *SubPhylum `pg:"rel:has-one"`

	// InfraPhylumID uuid.UUID    `pg:"infra_phylum_id"`
	// InfraPhylum   *InfraPhylum `pg:"rel:has-one"`

	// SuperClassID uuid.UUID   `pg:"super_class_id"`
	// SuperClass   *SuperClass `pg:"rel:has-one"`

	ClassID uuid.UUID `pg:"class_id"`
	Class   *Class    `pg:"rel:has-one"`

	// SubClassID uuid.UUID `pg:"sub_class_id"`
	// SubClass   *SubClass `pg:"rel:has-one"`

	// InfraClassID uuid.UUID   `pg:"infra_class_id"`
	// InfraClass   *InfraClass `pg:"rel:has-one"`

	// SuperOrderID uuid.UUID   `pg:"super_order_id"`
	// SuperOrder   *SuperOrder `pg:"rel:has-one"`

	OrderID uuid.UUID `pg:"order_id"`
	Order   *Order    `pg:"rel:has-one"`

	// SubOrderID uuid.UUID `pg:"sub_order_id"`
	// SubOrder   *SubOrder `pg:"rel:has-one"`

	// InfraOrderID uuid.UUID   `pg:"infra_order_id"`
	// InfraOrder   *InfraOrder `pg:"rel:has-one"`

	// SuperFamilyID uuid.UUID    `pg:"super_family_id"`
	// SuperFamily   *SuperFamily `pg:"rel:has-one"`

	FamilyID uuid.UUID `pg:"family_id"`
	Family   *Family   `pg:"rel:has-one"`

	// SubFamilyID uuid.UUID  `pg:"sub_family_id"`
	// SubFamily   *SubFamily `pg:"rel:has-one"`

	// SuperGenusID uuid.UUID   `pg:"super_genus_id"`
	// SuperGenus   *SuperGenus `pg:"rel:has-one"`

	GenusID uuid.UUID `pg:"genus_id"`
	Genus   *Genus    `pg:"rel:has-one"`

	// SubGenusID uuid.UUID `pg:"sub_genus_id"`
	// SubGenus   *SubGenus `pg:"rel:has-one"`

	// SpeciesImages []*SpeciesImage `pg:"rel:has-many"`
	// SpeciesVideos []*SpeciesVideo `pg:"rel:has-many"`

	// SpeciesFieldLocations []*SpeciesFieldLocation `pg:"rel:has-many"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
