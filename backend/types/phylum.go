package types

import (
	"github.com/google/uuid"
)

type PhylumDto struct {
	PhylumID   uuid.UUID `json:"phylumId"`
	PhylumName string    `json:"phylumName"`
}

type ListPhylumQuery struct {
	KingdomID uuid.UUID `json:"kingdomId"`
}

type ListPhylumResponse struct {
	BaseResponse
	Data []PhylumDto `json:"data"`
}
