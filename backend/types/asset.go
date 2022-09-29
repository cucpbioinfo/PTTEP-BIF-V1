package types

import (
	"github.com/google/uuid"
)

type AssetDto struct {
	AssetID   uuid.UUID `json:"assetId"`
	AssetName string    `json:"assetName"`
}

type ListAssetResponse struct {
	BaseResponse
	Data []AssetDto `json:"data"`
}
