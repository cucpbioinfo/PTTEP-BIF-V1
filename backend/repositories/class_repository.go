package repositories

import (
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/wichadak/eDNA/models"
	"github.com/wichadak/eDNA/types"
)

type ClassRepository struct {
	pg *pg.DB
}

func NewClassRepository(pg *pg.DB) *ClassRepository {
	return &ClassRepository{
		pg: pg,
	}
}

func (classRepository *ClassRepository) ListClass(query types.ListClassQuery) ([]models.Class, error) {
	var classes []models.Class
	dbQuery := classRepository.pg.Model(&classes)

	if query.PhylumID != uuid.Nil {
		dbQuery.Where("class.phylum_id = ?", query.PhylumID)
	}

	err := dbQuery.Select()
	if err != nil {
		return make([]models.Class, 0), err
	}
	return classes, nil
}
