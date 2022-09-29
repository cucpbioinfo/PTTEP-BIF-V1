package repositories

import (
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/wichadak/eDNA/models"
	"github.com/wichadak/eDNA/types"
)

type SpeciesVideoRepository struct {
	pg *pg.DB
}

func NewSpeciesVideoRepository(pg *pg.DB) *SpeciesVideoRepository {
	return &SpeciesVideoRepository{
		pg: pg,
	}
}

func (speciesVideoRepo *SpeciesVideoRepository) CreateSpeciesVideo(speciesVideo *types.SpeciesVideoDto) (uuid.UUID, error) {
	speciesVideoModel := &models.SpeciesVideo{
		SpeciesVideoID:  speciesVideo.SpeciesVideoID,
		SpeciesID:       speciesVideo.SpeciesID,
		FieldLocationID: speciesVideo.FieldLocationID,
		VideoSource:     speciesVideo.VideoSource,
		TakenAt:         speciesVideo.TakenAt,
	}
	_, err := speciesVideoRepo.pg.Model(speciesVideoModel).Returning("species_video_id").Insert()
	if err != nil {
		return uuid.Nil, err
	}
	return speciesVideoModel.SpeciesVideoID, nil
}
