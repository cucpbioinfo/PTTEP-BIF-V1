package repositories

import (
	"github.com/go-pg/pg/v10"
	"github.com/wichadak/eDNA/models"
)

type DensityRepository struct {
	pg *pg.DB
}

func NewDensityRepository(pg *pg.DB) *DensityRepository {
	return &DensityRepository{
		pg: pg,
	}
}

func (DensityRepository *DensityRepository) ListDensity() ([]models.Density, error) {
	var Densitys []models.Density
	err := DensityRepository.pg.Model(&Densitys).Select()
	if err != nil {
		return make([]models.Density, 0), err
	}
	return Densitys, nil
}

// func (DensityRepository *DensityRepository) ListDensity2(query types.DensityListQuery) ([]models.Density, int, error) {
// 	var Densitys []models.Density
// 	dbquery := DensityRepository.pg.Model(&Densitys)
// 	//limit, offset := utils.GetLimitAndOffset(query.PageNumber, query.PageSize)
// 	// if query.MajorGroupID != uuid.Nil {
// 	// 	dbquery.Where("density.major_group_id = ?", query.MajorGroupID)
// 	// }
// 	// if query.KingdomID != uuid.Nil {
// 	// 	dbquery.Where("density.kingdom_id = ?", query.KingdomID)
// 	// }
// 	// if query.PhylumID != uuid.Nil {
// 	// 	dbquery.Where("density.phylum_id = ?", query.PhylumID)
// 	// }
// 	// if query.ClassID != uuid.Nil {
// 	// 	dbquery.Where("density.class_id = ?", query.ClassID)
// 	// }
// 	// if query.OrderID != uuid.Nil {
// 	// 	dbquery.Where("density.order_id = ?", query.OrderID)
// 	// }
// 	// if query.FamilyID != uuid.Nil {
// 	// 	dbquery.Where("density.family_id = ?", query.FamilyID)
// 	// }
// 	// if query.GenusID != uuid.Nil {
// 	// 	dbquery.Where("density.genus_id = ?", query.GenusID)
// 	// }
// 	// if query.SpeciesID != uuid.Nil {
// 	// 	dbquery.Where("density.species_id = ?", query.SpeciesID)
// 	// }
// 	total, err := dbquery.
// 		Relation("Asset").
// 		Relation("Platform").
// 		Relation("Station").
// 		// Relation("Identification").
// 		// Relation("Method").
// 		// Relation("MajorGroup").
// 		// Relation("Kingdom").
// 		// Relation("Phylum").
// 		// Relation("Class").
// 		// Relation("Order").
// 		// Relation("Family").
// 		// Relation("Genus").
// 		Relation("Species").
// 		// Limit(limit).
// 		// Offset(offset).
// 		//SelectAndCount().
// 		Select()
// 	if err != nil {
// 		return make([]models.Density, 0), 0, err
// 	}
// 	return Densitys, total, nil
// }

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
