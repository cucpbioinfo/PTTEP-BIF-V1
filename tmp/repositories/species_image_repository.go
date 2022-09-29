package repositories

import (
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/wichadak/eDNA/models"
	"github.com/wichadak/eDNA/types"
)

type SpeciesImageRepository struct {
	pg *pg.DB
}

func NewSpeciesImageRepository(pg *pg.DB) *SpeciesImageRepository {
	return &SpeciesImageRepository{
		pg: pg,
	}
}

func (speciesImageRepo *SpeciesImageRepository) CreateSpeciesImage(speciesImage *types.SpeciesImageDto) (uuid.UUID, error) {
	speciesImageModel := &models.SpeciesImage{
		SpeciesImageID:  speciesImage.SpeciesImageID,
		SpeciesID:       speciesImage.SpeciesID,
		FieldLocationID: speciesImage.FieldLocationID,
		ImageSource:     speciesImage.ImageSource,
		TakenAt:         speciesImage.TakenAt,
	}
	_, err := speciesImageRepo.pg.Model(speciesImageModel).Returning("species_image_id").Insert()

	if err != nil {
		return uuid.Nil, err
	}
	return speciesImageModel.SpeciesImageID, nil
}
