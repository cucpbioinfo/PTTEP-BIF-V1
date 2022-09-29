package services

import (
	"github.com/wichadak/eDNA/repositories"
	"github.com/wichadak/eDNA/types"
)

type MajorGroupService struct {
	MajorGroupRepository *repositories.MajorGroupRepository
}

func NewMajorGroupService(majorGroupRepository *repositories.MajorGroupRepository) *MajorGroupService {
	return &MajorGroupService{
		MajorGroupRepository: majorGroupRepository,
	}
}

func (majorGroupService *MajorGroupService) ListMajorGroup() ([]types.MajorGroupDto, error) {
	majorGroups, err := majorGroupService.MajorGroupRepository.ListMajorGroup()
	if err != nil {
		return make([]types.MajorGroupDto, 0), err
	}
	majorGroupList := make([]types.MajorGroupDto, len(majorGroups))
	for i, majorGroup := range majorGroups {
		majorGroupList[i] = types.MajorGroupDto{
			MajorGroupID:   majorGroup.MajorGroupID,
			MajorGroupName: majorGroup.MajorGroupName,
		}
	}
	return majorGroupList, nil
}
