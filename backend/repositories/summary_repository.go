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
