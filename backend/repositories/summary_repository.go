package repositories

import (
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/wichadak/eDNA/models"
	"github.com/wichadak/eDNA/types"
)

type SummaryRepository struct {
	pg *pg.DB
}

func NewSummaryRepository(pg *pg.DB) *SummaryRepository {
	return &SummaryRepository{
		pg: pg,
	}
}

func (summaryRepository *SummaryRepository) ListSummary(query types.ListSummaryQuery) ([]models.Summary, error) {
	var summary []models.Summary
	dbQuery := summaryRepository.pg.Model(&summary)

	if query.Year != "" {
		dbQuery.Where("year = ?", query.Year)
	}
	if query.MajorGroupID != uuid.Nil {
		dbQuery.Where("major_group.major_group_id = ?", query.MajorGroupID)
	}
	if query.IdentificationID != uuid.Nil {
		dbQuery.Where("identification.identification_id = ?", query.IdentificationID)
	}
	if query.AssetID != uuid.Nil {
		dbQuery.Where("asset.asset_id = ?", query.AssetID)
	}
	if query.PlatformID != uuid.Nil {
		dbQuery.Where("platform.platform_id = ?", query.PlatformID)
	}
	if query.StationID != uuid.Nil {
		dbQuery.Where("station.station_id = ?", query.StationID)
	}

	err := dbQuery.
		Relation("MajorGroup").
		Relation("Identification").
		Relation("Asset").
		Relation("Platform").
		Relation("Station").
		Limit(1).
		Select()
	if err != nil {
		return make([]models.Summary, 0), err
	}
	return summary, nil
}

// all summary
func (summaryRepository *SummaryRepository) ListAllSummary(query types.ListSummaryQuery) ([]models.Summary, error) {
	var summary []models.Summary
	dbQuery := summaryRepository.pg.Model(&summary)

	if query.Year != "" {
		dbQuery.Where("year = ?", query.Year)
	}
	if query.MajorGroupID != uuid.Nil {
		dbQuery.Where("major_group.major_group_id = ?", query.MajorGroupID)
	}
	if query.IdentificationID != uuid.Nil {
		dbQuery.Where("identification.identification_id = ?", query.IdentificationID)
	}
	if query.AssetID != uuid.Nil {
		dbQuery.Where("asset.asset_id = ?", query.AssetID)
	}
	if query.PlatformID != uuid.Nil {
		dbQuery.Where("platform.platform_id = ?", query.PlatformID)
	}
	if query.StationID != uuid.Nil {
		dbQuery.Where("station.station_id = ?", query.StationID)
	}

	err := dbQuery.
		Relation("MajorGroup").
		Relation("Identification").
		Relation("Asset").
		Relation("Platform").
		Relation("Station").
		Select()
	if err != nil {
		return make([]models.Summary, 0), err
	}
	return summary, nil
}

// Platform platform
type PlatformSummaryRepository struct {
	pg *pg.DB
}

func NewPlatformSummaryRepository(pg *pg.DB) *PlatformSummaryRepository {
	return &PlatformSummaryRepository{
		pg: pg,
	}
}

func (platformsummaryRepository *PlatformSummaryRepository) ListPlatformSummary(query types.ListPlatformSummaryQuery) ([]models.PlatformSummary, error) {
	var platformsummary []models.PlatformSummary
	dbQuery := platformsummaryRepository.pg.Model(&platformsummary)

	if query.Year != "" {
		dbQuery.Where("year = ?", query.Year)
	}
	if query.MajorGroupID != uuid.Nil {
		dbQuery.Where("major_group_id = ?", query.MajorGroupID)
	}
	if query.IdentificationID != uuid.Nil {
		dbQuery.Where("identification_id = ?", query.IdentificationID)
	}
	if query.AssetID != uuid.Nil {
		dbQuery.Where("asset_id = ?", query.AssetID)
	}

	err := dbQuery.
		Relation("MajorGroup").
		Relation("Identification").
		Relation("Asset").
		Relation("Platform").
		//Limit(1).
		Select()
	if err != nil {
		return make([]models.PlatformSummary, 0), err
	}
	return platformsummary, nil
}

// Asset asset
type AssetSummaryRepository struct {
	pg *pg.DB
}

func NewAssetSummaryRepository(pg *pg.DB) *AssetSummaryRepository {
	return &AssetSummaryRepository{
		pg: pg,
	}
}

func (assetsummaryRepository *AssetSummaryRepository) ListAssetSummary(query types.ListAssetSummaryQuery) ([]models.AssetSummary, error) {
	var assetsummary []models.AssetSummary
	dbQuery := assetsummaryRepository.pg.Model(&assetsummary)

	if query.Year != "" {
		dbQuery.Where("year = ?", query.Year)
	}
	if query.MajorGroupID != uuid.Nil {
		dbQuery.Where("major_group.major_group_id = ?", query.MajorGroupID)
	}
	if query.IdentificationID != uuid.Nil {
		dbQuery.Where("identification.identification_id = ?", query.IdentificationID)
	}
	if query.AssetID != uuid.Nil {
		dbQuery.Where("asset.asset_id = ?", query.AssetID)
	}

	err := dbQuery.
		Relation("MajorGroup").
		Relation("Identification").
		Relation("Asset").
		Limit(1).
		Select()
	if err != nil {
		return make([]models.AssetSummary, 0), err
	}
	return assetsummary, nil
}

// yearsummary Year year
type YearSummaryRepository struct {
	pg *pg.DB
}

func NewYearSummaryRepository(pg *pg.DB) *YearSummaryRepository {
	return &YearSummaryRepository{
		pg: pg,
	}
}
func (yearsummaryRepository *YearSummaryRepository) ListYearSummary(query types.ListYearSummaryQuery) ([]models.YearSummary, error) {
	var yearsummary []models.YearSummary
	dbQuery := yearsummaryRepository.pg.Model(&yearsummary)

	if query.MajorGroupID != uuid.Nil {
		dbQuery.Where("major_group_id = ?", query.MajorGroupID)
	}
	if query.IdentificationID != uuid.Nil {
		dbQuery.Where("identification_id = ?", query.IdentificationID)
	}
	if query.AssetID != uuid.Nil {
		dbQuery.Where("asset_id = ?", query.AssetID)
	}
	if query.PlatformID != uuid.Nil {
		dbQuery.Where("platform_id = ?", query.PlatformID)
	}
	if query.StationID != uuid.Nil {
		dbQuery.Where("station_id = ?", query.StationID)
	}

	err := dbQuery.
		Order("year ASC").
		Group("year").
		Select()
	if err != nil {
		return make([]models.YearSummary, 0), err
	}
	return yearsummary, nil
}
