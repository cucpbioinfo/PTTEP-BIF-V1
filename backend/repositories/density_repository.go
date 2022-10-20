package repositories

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/wichadak/eDNA/models"
	"github.com/wichadak/eDNA/types"
)

type DensityRepository struct {
	pg *pg.DB
}

func NewDensityRepository(pg *pg.DB) *DensityRepository {
	return &DensityRepository{
		pg: pg,
	}
}

func (densityRepository *DensityRepository) ListDensity(query types.ListDensityQuery) ([]models.Density, error) {
	var density []models.Density
	dbQuery := densityRepository.pg.Model(&density)

	if query.Year != "" {
		dbQuery.Where("density.year = ?", query.Year)
		fmt.Println(query.Year)
	}
	if query.AssetID != uuid.Nil {
		dbQuery.Where("density.asset_id = ?", query.AssetID)
		fmt.Println(query.AssetID)
	}
	if query.PlatformID != uuid.Nil {
		dbQuery.Where("density.platform_id = ?", query.PlatformID)
		fmt.Println(query.PlatformID)
	}
	if query.StationID != uuid.Nil {
		dbQuery.Where("density.station_id = ?", query.StationID)
		fmt.Println(query.StationID)
	}

	// dbQuery.Where("density.station_id = ?", "0ffad568-4644-4318-ba70-04d9b1051c75")

	if query.SpeciesID != uuid.Nil {
		dbQuery.Where("density.species_id = ?", query.SpeciesID)
		fmt.Println(query.SpeciesID)
	}

	err := dbQuery.
		Relation("Asset").
		Relation("Platform").
		Relation("Station").
		Order("year ASC").
		Order("station_name ASC").
		Select()
	if err != nil {
		return make([]models.Density, 0), err
	}
	return density, nil
}

// Platform platform
type PlatformDensityRepository struct {
	pg *pg.DB
}

func NewPlatformDensityRepository(pg *pg.DB) *PlatformDensityRepository {
	return &PlatformDensityRepository{
		pg: pg,
	}
}

func (platformdensityRepository *PlatformDensityRepository) ListPlatformDensity(query types.ListPlatformDensityQuery) ([]models.PlatformDensity, error) {
	var platformdensity []models.PlatformDensity
	dbQuery := platformdensityRepository.pg.Model(&platformdensity)

	if query.Year != "" {
		dbQuery.Where("platformdensity.year = ?", query.Year)
	}
	if query.AssetID != uuid.Nil {
		dbQuery.Where("platformdensity.asset_id = ?", query.AssetID)
	}
	if query.PlatformID != uuid.Nil {
		dbQuery.Where("platformdensity.platform_id = ?", query.PlatformID)
	}
	if query.SpeciesID != uuid.Nil {
		dbQuery.Where("platformdensity.species_id = ?", query.SpeciesID)
	}

	err := dbQuery.
		Relation("Asset").
		Relation("Platform").
		Select()
	if err != nil {
		return make([]models.PlatformDensity, 0), err
	}
	return platformdensity, nil
}

// Asset asset
type AssetDensityRepository struct {
	pg *pg.DB
}

func NewAssetDensityRepository(pg *pg.DB) *AssetDensityRepository {
	return &AssetDensityRepository{
		pg: pg,
	}
}

func (assetdensityRepository *AssetDensityRepository) ListAssetDensity(query types.ListAssetDensityQuery) ([]models.AssetDensity, error) {
	var assetdensity []models.AssetDensity
	dbQuery := assetdensityRepository.pg.Model(&assetdensity)

	if query.Year != "" {
		dbQuery.Where("assetdensity.year = ?", query.Year)
	}
	if query.AssetID != uuid.Nil {
		dbQuery.Where("assetdensity.asset_id = ?", query.AssetID)
	}
	if query.SpeciesID != uuid.Nil {
		dbQuery.Where("assetdensity.species_id = ?", query.SpeciesID)
	}

	err := dbQuery.
		Relation("Asset").
		Select()
	if err != nil {
		return make([]models.AssetDensity, 0), err
	}
	return assetdensity, nil
}
