package repositories

import (
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/wichadak/eDNA/models"
	"github.com/wichadak/eDNA/types"
)

type PhylumRepository struct {
	pg *pg.DB
}

func NewPhylumRepository(pg *pg.DB) *PhylumRepository {
	return &PhylumRepository{
		pg: pg,
	}
}

func (phylumRepository *PhylumRepository) ListPhylum(query types.ListPhylumQuery) ([]models.Phylum, error) {
	var phylums []models.Phylum
	dbQuery := phylumRepository.pg.Model(&phylums)

	if query.KingdomID != uuid.Nil {
		dbQuery.Where("phylum.kingdom_id = ?", query.KingdomID)
	}

	err := dbQuery.Select()
	if err != nil {
		return make([]models.Phylum, 0), err
	}
	return phylums, nil
}
