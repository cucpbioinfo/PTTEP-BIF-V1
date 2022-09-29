package repositories

import (
	"github.com/go-pg/pg/v10"
	"github.com/wichadak/eDNA/models"
)

type AssetRepository struct {
	pg *pg.DB
}

func NewAssetRepository(pg *pg.DB) *AssetRepository {
	return &AssetRepository{
		pg: pg,
	}
}

func (assetRepository *AssetRepository) ListAsset() ([]models.Asset, error) {
	var assets []models.Asset
	dbQuery := assetRepository.pg.Model(&assets)

	err := dbQuery.Select()
	if err != nil {
		return make([]models.Asset, 0), err
	}
	return assets, nil
}
