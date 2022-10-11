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

	if query.MajorGroupID != uuid.Nil {
		dbQuery.Where("summary.major_group_id = ?", query.MajorGroupID)
	}
	if query.IdentificationID != uuid.Nil {
		dbQuery.Where("summary.identification_id = ?", query.IdentificationID)
	}
	if query.AssetID != uuid.Nil {
		dbQuery.Where("summary.asset_id = ?", query.AssetID)
	}
	if query.PlatformID != uuid.Nil {
		dbQuery.Where("summary.platform_id = ?", query.PlatformID)
	}
	if query.StationID != uuid.Nil {
		dbQuery.Where("summary.station_id = ?", query.StationID)
	}

	err := dbQuery.
		Relation("MajorGroup").
		Relation("Identification").
		Relation("Asset").
		Relation("Platform").
		Relation("Station").
		//Limit(1).
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

	if query.MajorGroupID != uuid.Nil {
		dbQuery.Where("platformsummary.major_group_id = ?", query.MajorGroupID)
	}
	if query.IdentificationID != uuid.Nil {
		dbQuery.Where("platformsummary.identification_id = ?", query.IdentificationID)
	}
	if query.AssetID != uuid.Nil {
		dbQuery.Where("platformsummary.asset_id = ?", query.AssetID)
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

	if query.MajorGroupID != uuid.Nil {
		dbQuery.Where("assetsummary.major_group_id = ?", query.MajorGroupID)
	}
	if query.IdentificationID != uuid.Nil {
		dbQuery.Where("assetsummary.identification_id = ?", query.IdentificationID)
	}
	if query.AssetID != uuid.Nil {
		dbQuery.Where("assetsummary.asset_id = ?", query.AssetID)
	}

	err := dbQuery.
		Relation("MajorGroup").
		Relation("Identification").
		Relation("Asset").
		//Limit(1).
		Select()
	if err != nil {
		return make([]models.AssetSummary, 0), err
	}
	return assetsummary, nil
}
