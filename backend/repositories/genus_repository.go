package repositories

import (
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/wichadak/eDNA/models"
	"github.com/wichadak/eDNA/types"
)

type GenusRepository struct {
	pg *pg.DB
}

func NewGenusRepository(pg *pg.DB) *GenusRepository {
	return &GenusRepository{
		pg: pg,
	}
}

func (genusRepository *GenusRepository) ListGenus(query types.ListGenusQuery) ([]models.Genus, error) {
	var genus []models.Genus
	dbQuery := genusRepository.pg.Model(&genus)

	if query.FamilyID != uuid.Nil {
		dbQuery.Where("genus.family_id = ?", query.FamilyID)
	}

	err := dbQuery.Select()
	if err != nil {
		return make([]models.Genus, 0), err
	}
	return genus, nil
}
