package repositories

import (
	"github.com/go-pg/pg/v10"
	"github.com/wichadak/eDNA/models"
)

type EvennessRepository struct {
	pg *pg.DB
}

func NewEvennessRepository(pg *pg.DB) *EvennessRepository {
	return &EvennessRepository{
		pg: pg,
	}
}

func (EvennessRepository *EvennessRepository) ListEvenness() ([]models.Evenness, error) {
	var Evennesses []models.Evenness
	dbquery := EvennessRepository.pg.Model(&Evennesses)
	// if query.Year != "" {
	// 	dbquery.Where("evenness.year = ?", query.Year)
	// }
	// if query.ProjectID != uuid.Nil {
	// 	dbquery.Where("evenness.project_id = ?", query.ProjectID)
	// }
	// if query.PlatformID != uuid.Nil {
	// 	dbquery.Where("evenness.platform_id = ?", query.PlatformID)
	// }
	err := dbquery.
		Relation("Asset").
		Relation("Platform").
		Relation("Station").
		Select()
	if err != nil {
		return make([]models.Evenness, 0), err
	}
	return Evennesses, nil
}

// func (EvennessRepository *EvennessRepository) ListEvennessYear(query types.EvennessListQuery) ([]models.EvennessYear, error) {
// 	var Evennesses []models.EvennessYear
// 	dbquery := EvennessRepository.pg.Model(&Evennesses)
// 	dbquery.Column("year")

// 	err := dbquery.
// 		Group("year").
// 		Select()
// 	if err != nil {
// 		return make([]models.EvennessYear, 0), err
// 	}
// 	return Evennesses, nil
// }

// func (EvennessRepository *EvennessRepository) ListEvennessProject(query types.EvennessListQuery) ([]models.Evenness, error) {
// 	var Evennesses []models.Evenness
// 	dbquery := EvennessRepository.pg.Model(&Evennesses)
// 	//dbquery.Column("year")
// 	//Column("project.project_name")
// 	//dbquery.ColumnExpr("distinct on(evenness.project_id) project_id")
// 	if query.Year != "" {
// 		dbquery.Where("evenness.year = ?", query.Year)
// 	}
// 	// if query.ProjectID != uuid.Nil {
// 	// 	dbquery.Where("evenness.project_id = ?", query.ProjectID)
// 	// }
// 	// if query.PlatformID != uuid.Nil {
// 	// 	dbquery.Where("evenness.platform_id = ?", query.PlatformID)
// 	// }
// 	// if query.StationID != uuid.Nil {
// 	// 	dbquery.Where("evenness.station_id = ?", query.StationID)
// 	// }
// 	err := dbquery.
// 		//Join("LEFT JOIN project ON project.project_id = evenness.project_id").
// 		Relation("Project").
// 		//Relation("Platform").
// 		//Relation("Station").
// 		//Group("Year").
// 		Select()
// 	if err != nil {
// 		return make([]models.Evenness, 0), err
// 	}
// 	return Evennesses, nil
// }
