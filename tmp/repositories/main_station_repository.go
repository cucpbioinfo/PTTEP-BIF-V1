package repositories

import (
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/wichadak/eDNA/models"
	"github.com/wichadak/eDNA/types"
)

type MainStationRepository struct {
	pg *pg.DB
}

func NewMainStationRepository(pg *pg.DB) *MainStationRepository {
	return &MainStationRepository{
		pg: pg,
	}
}

func (mainStationRepository *MainStationRepository) ListMainStation(query *types.MainStationListQuery) ([]models.MainStation, error) {
	var mainStations []models.MainStation

	dbQuery := mainStationRepository.pg.Model(&mainStations)

	if query.SiteID != uuid.Nil {
		dbQuery.Where("site_id = ?", query.SiteID)
	}

	err := dbQuery.
		Relation("FieldLocations").
		Select()
	if err != nil {
		return make([]models.MainStation, 0), err
	}
	return mainStations, nil
}
