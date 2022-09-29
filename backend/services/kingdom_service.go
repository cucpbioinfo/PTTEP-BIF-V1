package services

import (
	"github.com/wichadak/eDNA/repositories"
	"github.com/wichadak/eDNA/types"
)

type KingdomService struct {
	KingdomRepository *repositories.KingdomRepository
}

func NewKingdomService(kingdomRepository *repositories.KingdomRepository) *KingdomService {
	return &KingdomService{
		KingdomRepository: kingdomRepository,
	}
}

func (kingdomService *KingdomService) ListKingdom() ([]types.KingdomDto, error) {
	kingdoms, err := kingdomService.KingdomRepository.ListKingdom()
	if err != nil {
		return make([]types.KingdomDto, 0), err
	}
	kingdomList := make([]types.KingdomDto, len(kingdoms))
	for i, kingdom := range kingdoms {
		kingdomList[i] = types.KingdomDto{
			KingdomID:   kingdom.KingdomID,
			KingdomName: kingdom.KingdomName,
		}
	}
	return kingdomList, nil
}
