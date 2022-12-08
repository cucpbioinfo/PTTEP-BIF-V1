package repositories

import (
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/wichadak/eDNA/models"
	"github.com/wichadak/eDNA/types"
)

type LocationRepository struct {
	pg *pg.DB
}

func NewLocationRepository(pg *pg.DB) *LocationRepository {
	return &LocationRepository{
		pg: pg,
	}
}

func (locationRepository *LocationRepository) ListLocation(query types.ListLocationQuery) ([]models.Location, error) {
	var location []models.Location
	dbQuery := locationRepository.pg.Model(&location)

	if query.AssetID != uuid.Nil {
		dbQuery.Where("location.asset_id = ?", query.AssetID)
	}
	if query.PlatformID != uuid.Nil {
		dbQuery.Where("location.platform_id = ?", query.PlatformID)
	}
	if query.StationID != uuid.Nil {
		dbQuery.Where("location.station_id = ?", query.StationID)
	}

	err := dbQuery.
		Relation("Asset").
		Relation("Platform").
		Relation("Station").
		Select()
	if err != nil {
		return make([]models.Location, 0), err
	}
	return location, nil
}

// Continued Here
func (locationRepository *LocationRepository) ListLocationAsset(query types.ListLocationQuery) ([]models.Location, error) {
	var location []models.Location
	dbQuery := locationRepository.pg.Model(&location)

	err := dbQuery.
		Relation("Asset").
		Select()
	if err != nil {
		return make([]models.Location, 0), err
	}
	return location, nil
}

func (locationRepository *LocationRepository) ListLocationSelect(query types.ListLocationQuery) ([]models.Location, error) {
	var location []models.Location
	dbQuery := locationRepository.pg.Model(&location)

	if query.AssetID != uuid.Nil {
		dbQuery.Where("location.asset_id = ?", query.AssetID)
	}
	if query.PlatformID != uuid.Nil {
		dbQuery.Where("location.platform_id = ?", query.PlatformID)
	}
	if query.StationID != uuid.Nil {
		dbQuery.Where("location.station_id = ?", query.StationID)
	}

	err := dbQuery.
		Relation("Asset").
		Relation("Platform").
		Relation("Station").
		Limit(1).
		Select()
	if err != nil {
		return make([]models.Location, 0), err
	}
	return location, nil
}
