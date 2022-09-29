package services

import (
	"github.com/wichadak/eDNA/repositories"
	"github.com/wichadak/eDNA/types"
)

type FieldLocationService struct {
	FieldLocationRepository *repositories.FieldLocationRepository
}

func NewFieldLocationService(fieldLocationRepository *repositories.FieldLocationRepository) *FieldLocationService {
	return &FieldLocationService{
		FieldLocationRepository: fieldLocationRepository,
	}
}

func (fieldLocationService *FieldLocationService) ListFieldLocation(query *types.FieldLocationListQuery) ([]types.FieldLocationDto, error) {
	fieldLocations, err := fieldLocationService.FieldLocationRepository.ListFieldLocation(query)
	if err != nil {
		return make([]types.FieldLocationDto, 0), err
	}
	fieldLocationList := make([]types.FieldLocationDto, len(fieldLocations))
	for i, fieldLocation := range fieldLocations {
		fieldLocationList[i] = types.FieldLocationDto{
			FieldLocationID:   fieldLocation.FieldLocationID,
			FieldLocationName: fieldLocation.FieldLocationName,
			Location: &types.Location{
				Latitude:  fieldLocation.FieldLocationLat,
				Longitude: fieldLocation.FieldLocationLong,
			},
			Replication: fieldLocation.Replication,
			CreatedAt:   fieldLocation.CreatedAt,
			UpdatedAt:   fieldLocation.UpdatedAt,
		}
	}
	return fieldLocationList, nil
}
