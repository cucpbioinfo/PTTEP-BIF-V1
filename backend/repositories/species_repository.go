package repositories

import (
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/wichadak/eDNA/models"
	"github.com/wichadak/eDNA/types"
	"github.com/wichadak/eDNA/utils"
)

type SpeciesRepository struct {
	pg *pg.DB
}

func NewSpeciesRepository(pg *pg.DB) *SpeciesRepository {
	return &SpeciesRepository{
		pg: pg,
	}
}

func (speciesRepository *SpeciesRepository) ListSpecies(query types.SpeciesListQuery) ([]models.Species, int, error) {
	var species []models.Species
	dbQuery := speciesRepository.pg.Model(&species).
		Where("species.scientific_name ILIKE ?", "%"+query.Keyword+"%").
		WhereOr("species.common_name_th ILIKE ?", "%"+query.Keyword+"%").
		WhereOr("species.common_name_en ILIKE ?", "%"+query.Keyword+"%").
		WhereOr("species.species_name ILIKE ?", "%"+query.Keyword+"%")

	limit, offset := utils.GetLimitAndOffset(query.PageNumber, query.PageSize)

	if query.MajorGroupID != uuid.Nil {
		dbQuery.Where("species.major_group_id = ?", query.MajorGroupID)
	}
	if query.KingdomID != uuid.Nil {
		dbQuery.Where("species.kingdom_id = ?", query.KingdomID)
	}
	if query.PhylumID != uuid.Nil {
		dbQuery.Where("species.phylum_id = ?", query.PhylumID)
	}
	if query.ClassID != uuid.Nil {
		dbQuery.Where("species.class_id = ?", query.ClassID)
	}
	if query.OrderID != uuid.Nil {
		dbQuery.Where("species.order_id = ?", query.OrderID)
	}
	if query.FamilyID != uuid.Nil {
		dbQuery.Where("species.family_id = ?", query.FamilyID)
	}
	if query.GenusID != uuid.Nil {
		dbQuery.Where("species.genus_id = ?", query.GenusID)
	}

	total, err := dbQuery.
		Relation("Family").
		Relation("Genus").
		//Relation("SpeciesImages").
		Limit(limit).
		Offset(offset).
		SelectAndCount()
	if err != nil {
		return make([]models.Species, 0), 0, &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}
	return species, total, nil
}

// func (speciesRepository *SpeciesRepository) GetSpeciesCountGroupBySpeciesType() ([]types.SpeciesCountGroupBySpeciesType, error) {

// 	var result []types.SpeciesCountGroupBySpeciesType
// 	err := speciesRepository.pg.Model((*models.Species)(nil)).
// 		Column("species_type").
// 		ColumnExpr("count(*) AS count").
// 		Group("species_type").
// 		Select(&result)
// 	if err != nil {
// 		return make([]types.SpeciesCountGroupBySpeciesType, 0), err
// 	}
// 	return result, nil
// }

func (speciesRepository *SpeciesRepository) GetSpeciesById(speciesId uuid.UUID) (models.Species, error) {
	var species models.Species
	err := speciesRepository.pg.Model(&species).
		Where("species_id = ?", speciesId).
		Relation("Identification").
		Relation("Method").
		Relation("MajorGroup").
		Relation("Kingdom").
		Relation("Phylum").
		Relation("Class").
		Relation("Order").
		Relation("Family").
		Relation("Genus").
		// Relation("SpeciesImages").
		// Relation("SpeciesVideos").
		Select()
	if err != nil {
		return models.Species{}, err
	}
	return species, nil
}

// func (speciesRepository *SpeciesRepository) SearchSpecies(query types.SpeciesSearchQuery) ([]models.Species, error) {
// 	var species []models.Species
// 	dbQuery := speciesRepository.pg.Model(&species)
// 	// if query.Keyword == "" {
// 	// 	return make([]models.Species, 0), nil
// 	// }

// 	err := dbQuery.
// 		Where("species.scientific_name ILIKE ?", "%"+query.Keyword+"%").
// 		WhereOr("species.common_name_th ILIKE ?", "%"+query.Keyword+"%").
// 		WhereOr("species.common_name_en ILIKE ?", "%"+query.Keyword+"%").
// 		Relation("Kingdom").
// 		Relation("SuperPhylum").
// 		Relation("Phylum").
// 		Relation("SubPhylum").
// 		Relation("SuperClass").
// 		Relation("Class").
// 		Relation("SubClass").
// 		Relation("SuperOrder").
// 		Relation("Order").
// 		Relation("SubOrder").
// 		Relation("SuperFamily").
// 		Relation("Family").
// 		Relation("SubFamily").
// 		Relation("SuperGenus").
// 		Relation("Genus").
// 		Relation("SubGenus").
// 		Relation("SubSpecies").
// 		// Relation("SpeciesImages").
// 		// Relation("SpeciesVideos").
// 		// Relation("SpeciesFieldLocations").
// 		// Relation("SpeciesImages").
// 		Limit(query.Limit).
// 		Select()

// 	if err != nil {
// 		return make([]models.Species, 0), &fiber.Error{
// 			Code:    500,
// 			Message: err.Error(),
// 		}
// 	}

// 	return species, nil
// }
