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
		dbQuery.Where("year = ?", query.Year)
		fmt.Println(query.Year)
	}
	if query.AssetID != uuid.Nil {
		dbQuery.Where("asset.asset_id = ?", query.AssetID)
		fmt.Println(query.AssetID)
	}
	if query.PlatformID != uuid.Nil {
		dbQuery.Where("platform.platform_id = ?", query.PlatformID)
		fmt.Println(query.PlatformID)
	}
	if query.StationID != uuid.Nil {
		dbQuery.Where("station.station_id = ?", query.StationID)
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
		Order("year DESC").
		//Order("station_name ASC").
		Limit(1).
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
		dbQuery.Where("year = ?", query.Year)
	}
	if query.AssetID != uuid.Nil {
		dbQuery.Where("asset.asset_id = ?", query.AssetID)
	}
	if query.PlatformID != uuid.Nil {
		dbQuery.Where("platform.platform_id = ?", query.PlatformID)
	}
	if query.SpeciesID != uuid.Nil {
		dbQuery.Where("species_id = ?", query.SpeciesID)
	}

	err := dbQuery.
		Relation("Asset").
		Relation("Platform").
		Order("year DESC").
		//Order("station_name ASC").
		Limit(1).
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
		dbQuery.Where("year = ?", query.Year)
	}
	//dbQuery.Where("assetdensity.year = ?", 2021)
	if query.AssetID != uuid.Nil {
		dbQuery.Where("asset.asset_id = ?", query.AssetID)
	}
	if query.SpeciesID != uuid.Nil {
		dbQuery.Where("species_id = ?", query.SpeciesID)
	}

	err := dbQuery.
		Relation("Asset").
		Order("year DESC").
		//Order("station_name ASC").
		Limit(1).
		Select()
	if err != nil {
		return make([]models.AssetDensity, 0), err
	}
	return assetdensity, nil
}

// YearAsset yearasset
type YearAssetDensityRepository struct {
	pg *pg.DB
}

func NewYearAssetDensityRepository(pg *pg.DB) *YearAssetDensityRepository {
	return &YearAssetDensityRepository{
		pg: pg,
	}
}

func (yearassetdensityRepository *YearAssetDensityRepository) ListYearAssetDensity(query types.ListYearAssetDensityQuery) ([]models.YearAssetDensity, error) {
	var yearassetdensity []models.YearAssetDensity
	dbQuery := yearassetdensityRepository.pg.Model(&yearassetdensity)

	//dbQuery.Where("assetdensity.year = ?", 2021)
	if query.AssetID != uuid.Nil {
		dbQuery.Where("asset_id = ?", query.AssetID)
	}
	if query.SpeciesID != uuid.Nil {
		dbQuery.Where("species_id = ?", query.SpeciesID)
	}

	err := dbQuery.
		//Relation("Asset").
		Order("year ASC").
		//Order("station_name ASC").
		Group("year").
		Select()
	if err != nil {
		return make([]models.YearAssetDensity, 0), err
	}
	return yearassetdensity, nil
}

// Year Platform year platform
type YearPlatformDensityRepository struct {
	pg *pg.DB
}

func NewYearPlatformDensityRepository(pg *pg.DB) *YearPlatformDensityRepository {
	return &YearPlatformDensityRepository{
		pg: pg,
	}
}

func (yearplatformdensityRepository *YearPlatformDensityRepository) ListYearPlatformDensity(query types.ListYearPlatformDensityQuery) ([]models.YearPlatformDensity, error) {
	var yearplatformdensity []models.YearPlatformDensity
	dbQuery := yearplatformdensityRepository.pg.Model(&yearplatformdensity)

	//dbQuery.Where("assetdensity.year = ?", 2021)
	if query.PlatformID != uuid.Nil {
		dbQuery.Where("platform_id = ?", query.PlatformID)
	}
	if query.AssetID != uuid.Nil {
		dbQuery.Where("asset_id = ?", query.AssetID)
	}
	if query.SpeciesID != uuid.Nil {
		dbQuery.Where("species_id = ?", query.SpeciesID)
	}

	err := dbQuery.
		//Relation("Asset").
		Order("year ASC").
		//Order("station_name ASC").
		Group("year").
		Select()
	if err != nil {
		return make([]models.YearPlatformDensity, 0), err
	}
	return yearplatformdensity, nil
}

// Year Station year station
type YearDensityRepository struct {
	pg *pg.DB
}

func NewYearDensityRepository(pg *pg.DB) *YearDensityRepository {
	return &YearDensityRepository{
		pg: pg,
	}
}

func (yeardensityRepository *YearDensityRepository) ListYearDensity(query types.ListYearDensityQuery) ([]models.YearDensity, error) {
	var yeardensity []models.YearDensity
	dbQuery := yeardensityRepository.pg.Model(&yeardensity)

	//dbQuery.Where("assetdensity.year = ?", 2021)
	if query.StationID != uuid.Nil {
		dbQuery.Where("station_id = ?", query.StationID)
	}
	if query.PlatformID != uuid.Nil {
		dbQuery.Where("platform_id = ?", query.PlatformID)
	}
	if query.AssetID != uuid.Nil {
		dbQuery.Where("asset_id = ?", query.AssetID)
	}
	if query.SpeciesID != uuid.Nil {
		dbQuery.Where("species_id = ?", query.SpeciesID)
	}

	err := dbQuery.
		//Relation("Asset").
		Order("year ASC").
		//Order("station_name ASC").
		Group("year").
		Select()
	if err != nil {
		return make([]models.YearDensity, 0), err
	}
	return yeardensity, nil
}
