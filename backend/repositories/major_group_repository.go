package repositories

import (
	"github.com/go-pg/pg/v10"
	"github.com/wichadak/eDNA/models"
)

type MajorGroupRepository struct {
	pg *pg.DB
}

func NewMajorGroupRepository(pg *pg.DB) *MajorGroupRepository {
	return &MajorGroupRepository{
		pg: pg,
	}
}

func (majorGroupRepository *MajorGroupRepository) ListMajorGroup() ([]models.MajorGroup, error) {
	var majorGroups []models.MajorGroup
	dbQuery := majorGroupRepository.pg.Model(&majorGroups)

	err := dbQuery.Select()
	if err != nil {
		return make([]models.MajorGroup, 0), err
	}
	return majorGroups, nil
}
