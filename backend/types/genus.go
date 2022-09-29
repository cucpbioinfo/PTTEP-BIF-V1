package types

import (
	"github.com/google/uuid"
)

type GenusCreateBody struct {
	GenusName string `json:"genusName"`
}

type GenusDto struct {
	GenusID   uuid.UUID `json:"genusId"`
	GenusName string    `json:"genusName"`
}

type ListGenusQuery struct {
	FamilyID uuid.UUID `json:"familyId"`
}

type ListGenusResponse struct {
	BaseResponse
	Data []GenusDto `json:"data"`
}
