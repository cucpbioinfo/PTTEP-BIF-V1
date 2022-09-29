package types

import (
	"github.com/google/uuid"
)

type KingdomCreateBody struct {
	KingdomName string `json:"kingdomName"`
}

type KingdomDto struct {
	KingdomID   uuid.UUID `json:"kingdomId"`
	KingdomName string    `json:"kingdomName"`
}

type ListKingdomResponse struct {
	BaseResponse
	Data []KingdomDto `json:"data"`
}
