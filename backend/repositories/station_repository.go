package repositories

import (
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/wichadak/eDNA/models"
	"github.com/wichadak/eDNA/types"
)

type StationRepository struct {
	pg *pg.DB
}

func NewStationRepository(pg *pg.DB) *StationRepository {
	return &StationRepository{
		pg: pg,
	}
}

func (stationRepository *StationRepository) ListStation(query types.ListStationQuery) ([]models.Station, error) {
	var station []models.Station
	dbQuery := stationRepository.pg.Model(&station)

	if query.PlatformID != uuid.Nil {
		dbQuery.Where("station.platform_id = ?", query.PlatformID)
	}

	err := dbQuery.Select()
	if err != nil {
		return make([]models.Station, 0), err
	}
	return station, nil
}
