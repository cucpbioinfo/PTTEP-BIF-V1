package repositories

import (
	"github.com/go-pg/pg/v10"
	"github.com/wichadak/eDNA/models"
)

type KingdomRepository struct {
	pg *pg.DB
}

func NewKingdomRepository(pg *pg.DB) *KingdomRepository {
	return &KingdomRepository{
		pg: pg,
	}
}

func (kingdomRepository *KingdomRepository) ListKingdom() ([]models.Kingdom, error) {
	var kingdoms []models.Kingdom
	err := kingdomRepository.pg.Model(&kingdoms).Select()
	if err != nil {
		return make([]models.Kingdom, 0), err
	}
	return kingdoms, nil
}
