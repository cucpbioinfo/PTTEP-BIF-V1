package types

import (
	"time"

	"github.com/google/uuid"
)

type EvennessDto struct {
	EvennessID    uuid.UUID `json:"EvennessId"`
	Year          string    `json:"Year"`
	AssetName     string    `json:"AssetName"`
	PlatformName  string    `json:"PlatformName"`
	StationName   string    `json:"StationName"`
	Surface       string    `json:"Surface"`
	Euphotic_zone string    `json:"Euphotic_zone"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type EvennessListResponse struct {
	BaseResponse
	Data []EvennessDto `json:"data"`
}
