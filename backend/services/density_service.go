package services

import (
	"github.com/wichadak/eDNA/repositories"
	"github.com/wichadak/eDNA/types"
)

type DensityService struct {
	DensityRepository *repositories.DensityRepository
}

func NewDensityService(densityRepository *repositories.DensityRepository) *DensityService {
	return &DensityService{
		DensityRepository: densityRepository,
	}
}

func (densityService *DensityService) ListDensity(query types.ListDensityQuery) ([]types.DensityDto, error) {
	density, err := densityService.DensityRepository.ListDensity(query)
	if err != nil {
		return make([]types.DensityDto, 0), err
	}
	densityList := make([]types.DensityDto, len(density))
	for i, density := range density {
		densityList[i] = types.DensityDto{
			DensityID:     density.DensityID,
			Year:          density.Year,
			AssetName:     density.Asset.AssetName,
			PlatformName:  density.Platform.PlatformName,
			StationName:   density.Station.StationName,
			Surface:       density.Surface,
			Euphotic_zone: density.Euphotic_zone,
		}
	}
	return densityList, nil
}
