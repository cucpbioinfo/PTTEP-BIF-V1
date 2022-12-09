package repositories

import (
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/wichadak/eDNA/models"
	"github.com/wichadak/eDNA/types"
)

type PlatformRepository struct {
	pg *pg.DB
}

func NewPlatformRepository(pg *pg.DB) *PlatformRepository {
	return &PlatformRepository{
		pg: pg,
	}
}

func (platformRepository *PlatformRepository) ListPlatform(query types.ListPlatformQuery) ([]models.Platform, error) {
	var platform []models.Platform
	dbQuery := platformRepository.pg.Model(&platform)

	if query.AssetID != uuid.Nil {
		dbQuery.Where("platform.asset_id = ?", query.AssetID)
	}

	err := dbQuery.Order("platform_name ASC").Select()
	if err != nil {
		return make([]models.Platform, 0), err
	}
	return platform, nil
}
