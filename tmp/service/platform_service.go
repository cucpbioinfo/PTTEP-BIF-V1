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

func (platformService *PlatformService) ListPlatform() ([]types.PlatformDto, error) {
	platforms, err := platformService.PlatformRepository.ListPlatform()
	if err != nil {
		return make([]types.PlatformDto, 0), err
	}
	platformList := make([]types.PlatformDto, len(platforms))
	for i, platform := range platforms {
		platformList[i] = types.PlatformDto{
			PlatformID:   platform.PlatformID,
			PlatformName: platform.PlatformName,
			ProjectID:    platform.ProjectID,
			ProjectName:  platform.Project.ProjectName,
			Latitude:     platform.PlatformLat,
			Longitude:    platform.PlatformLong,
		}
	}
	return platformList, nil
}
