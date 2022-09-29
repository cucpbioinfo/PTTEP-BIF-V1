package repositories

import (
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/wichadak/eDNA/models"
	"github.com/wichadak/eDNA/types"
)

type FamilyRepository struct {
	pg *pg.DB
}

func NewFamilyRepository(pg *pg.DB) *FamilyRepository {
	return &FamilyRepository{
		pg: pg,
	}
}

func (familyRepository *FamilyRepository) ListFamily(query types.ListFamilyQuery) ([]models.Family, error) {
	var families []models.Family
	dbQuery := familyRepository.pg.Model(&families)

	if query.OrderID != uuid.Nil {
		dbQuery.Where("family.order_id = ?", query.OrderID)
	}

	err := dbQuery.Select()
	if err != nil {
		return make([]models.Family, 0), err
	}
	return families, nil
}
