package repositories

import (
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/wichadak/eDNA/models"
	"github.com/wichadak/eDNA/types"
)

type FieldLocationRepository struct {
	pg *pg.DB
}

func NewFieldLocationRepository(pg *pg.DB) *FieldLocationRepository {
	return &FieldLocationRepository{
		pg: pg,
	}
}

func (fieldLocationRepository *FieldLocationRepository) ListFieldLocation(query *types.FieldLocationListQuery) ([]models.FieldLocation, error) {
	var fieldLocations []models.FieldLocation

	dbQuery := fieldLocationRepository.pg.Model(&fieldLocations)

	if query.MainStationID != uuid.Nil {
		dbQuery.Where("main_station_id = ?", query.MainStationID)
	}

	err := dbQuery.Select()
	if err != nil {
		return make([]models.FieldLocation, 0), err
	}
	return fieldLocations, nil
}
