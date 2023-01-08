package services

import (
	"github.com/wichadak/eDNA/repositories"
	"github.com/wichadak/eDNA/types"
)

type PlatformService struct {
	PlatformRepository *repositories.PlatformRepository
}

func NewPlatformService(platformRepository *repositories.PlatformRepository) *PlatformService {
	return &PlatformService{
		PlatformRepository: platformRepository,
	}
}

func (platformService *PlatformService) ListPlatform(query types.ListPlatformQuery) ([]types.PlatformDto, error) {
	platform, err := platformService.PlatformRepository.ListPlatform(query)
	if err != nil {
		return make([]types.PlatformDto, 0), err
	}
	platformList := make([]types.PlatformDto, len(platform))
	for i, platform := range platform {
		platformList[i] = types.PlatformDto{
			PlatformID:   platform.PlatformID,
			PlatformName: platform.PlatformName,
			Latitude:     platform.Latitude,
			Longitude:    platform.Longitude,
			Type:         platform.Type,
		}
	}
	return platformList, nil
}

func (platformService *PlatformService) ListRefPlatform(query types.ListPlatformQuery) ([]types.PlatformDto, error) {
	platform, err := platformService.PlatformRepository.ListRefPlatform(query)
	if err != nil {
		return make([]types.PlatformDto, 0), err
	}
	platformList := make([]types.PlatformDto, len(platform))
	for i, platform := range platform {
		platformList[i] = types.PlatformDto{
			PlatformID:   platform.PlatformID,
			PlatformName: platform.PlatformName,
			Latitude:     platform.Latitude,
			Longitude:    platform.Longitude,
			Type:         platform.Type,
		}
	}
	return platformList, nil
}
