package services

import (
	"github.com/wichadak/eDNA/repositories"
	"github.com/wichadak/eDNA/types"
)

type AssetService struct {
	AssetRepository *repositories.AssetRepository
}

func NewAssetService(assetRepository *repositories.AssetRepository) *AssetService {
	return &AssetService{
		AssetRepository: assetRepository,
	}
}

func (assetService *AssetService) ListAsset() ([]types.AssetDto, error) {
	assets, err := assetService.AssetRepository.ListAsset()
	if err != nil {
		return make([]types.AssetDto, 0), err
	}
	assetList := make([]types.AssetDto, len(assets))
	for i, asset := range assets {
		assetList[i] = types.AssetDto{
			AssetID:   asset.AssetID,
			AssetName: asset.AssetName,
		}
	}
	return assetList, nil
}
