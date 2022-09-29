package types

import (
	"github.com/google/uuid"
)

type FamilyCreateBody struct {
	FamilyName string `json:"familyName"`
}

type FamilyDto struct {
	FamilyID   uuid.UUID `json:"familyId"`
	FamilyName string    `json:"familyName"`
}

type ListFamilyQuery struct {
	OrderID uuid.UUID `json:"orderId"`
}

type ListFamilyResponse struct {
	BaseResponse
	Data []FamilyDto `json:"data"`
}
