package repositories

import (
	"github.com/go-pg/pg/v10"
	"github.com/wichadak/eDNA/models"
)

type SiteRepository struct {
	pg *pg.DB
}

func NewSiteRepository(pg *pg.DB) *SiteRepository {
	return &SiteRepository{
		pg: pg,
	}
}

func (siteRepository *SiteRepository) ListSite() ([]models.Site, error) {
	var sites []models.Site
	dbQuery := siteRepository.pg.Model(&sites)

	err := dbQuery.
		Relation("MainStations").
		Relation("MainStations.FieldLocations").
		Select()
	if err != nil {
		return make([]models.Site, 0), err
	}
	return sites, nil
}
