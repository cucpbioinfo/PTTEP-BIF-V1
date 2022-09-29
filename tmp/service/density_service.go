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

func (densityService *DensityService) ListDensity(query types.DensityListQuery) ([]types.DensityDto, int, error) {
	densitys, total, err := densityService.DensityRepository.ListDensity(query)
	if err != nil {
		return make([]types.DensityDto, 0), 0, err
	}
	densityList := make([]types.DensityDto, len(densitys))
	for i, density := range densitys {
		densityList[i] = types.DensityDto{
			DensityID: density.DensityID,
			Density:   density.Density,
			Year:      density.DensityYear,
			Asset:     density.Asset.AssetName,
			Platform:  density.Platform.PlatformName,
			Station:   density.Station.StationName,
			// Identification: density.Identification.IdentificationName,
			// Method:         density.Method.MethodName,
			// MajorGroup:     density.MajorGroup.MajorGroupName,
			// Kingdom:        density.Kingdom.KingdomName,
			// Phylum:         density.Phylum.PhylumName,
			// Class:          density.Class.ClassName,
			// Order:          density.Order.OrderName,
			// Family:         density.Family.FamilyName,
			// Genus:          density.Genus.GenusName,
			Species: density.Species.SpeciesName,
		}
	}
	return densityList, total, nil
}
