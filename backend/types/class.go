package types

import (
	"github.com/google/uuid"
)

type ClassCreateBody struct {
	ClassName string `json:"className"`
}

type ClassDto struct {
	ClassID   uuid.UUID `json:"classId"`
	ClassName string    `json:"className"`
}

type ListClassQuery struct {
	PhylumID uuid.UUID `json:"phylumId"`
}

type ListClassResponse struct {
	BaseResponse
	Data []ClassDto `json:"data"`
}
