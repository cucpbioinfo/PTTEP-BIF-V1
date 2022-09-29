package types

import (
	"time"

	"github.com/google/uuid"
)

type SpeciesListQuery struct {
	Keyword      string    `query:"keyword"`
	MajorGroupID uuid.UUID `query:"majorGroupId"`
	KingdomID    uuid.UUID `query:"kingdomId"`
	PhylumID     uuid.UUID `query:"phylumId"`
	ClassID      uuid.UUID `query:"classId"`
	OrderID      uuid.UUID `query:"orderId"`
	FamilyID     uuid.UUID `query:"familyId"`
	GenusID      uuid.UUID `query:"genusId"`
	PageNumber   int       `query:"pageNumber"`
	PageSize     int       `query:"pageSize"`
}

type SpeciesListDto struct {
	SpeciesID   uuid.UUID `json:"speciesId"`
	SpeciesName string    `json:"speciesName"`
	// KingdomName  string    `json:"kingdomName"`
	// PhylumName   string    `json:"phylumName"`
	FamilyName string `json:"familyName"`
	GenusName  string `json:"genusName"`
	//SpeciesImage string    `json:"speciesImage"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type SpeciesListMeta struct {
	SpeciesListQuery   `json:"query"`
	PaginationResponse `json:"pagination"`
}

type SpeciesListResponse struct {
	BaseResponse
	Data []SpeciesListDto   `json:"data"`
	Meta PaginationResponse `json:"meta"`
}

// type SpeciesCountGroupBySpeciesType struct {
// 	SpeciesType string `json:"speciesType"`
// 	Count       int    `json:"count"`
// }

// type SpeciesCountGroupBySpeciesTypeResponse struct {
// 	BaseResponse
// 	Data []SpeciesCountGroupBySpeciesType `json:"data"`
// }

type SpeciesDetailsDto struct {
	SpeciesID   uuid.UUID `json:"speciesId"`
	SpeciesName string    `json:"speciesName"`
	//ScientificName string    `json:"scientificName"`
	IdentificationID   uuid.UUID `json:"identificationId"`
	IdentificationName string    `json:"identificationName"`
	MethodID           uuid.UUID `json:"methodId"`
	MethodName         string    `json:"methodName"`
	MajorGroupID       uuid.UUID `json:"majorGroupId"`
	MajorGroupName     string    `json:"majorGroupName"`
	KingdomID          uuid.UUID `json:"kingdomId"`
	KingdomName        string    `json:"kingdomName"`
	PhylumID           uuid.UUID `json:"phylumId"`
	PhylumName         string    `json:"phylumName"`
	ClassID            uuid.UUID `json:"classId"`
	ClassName          string    `json:"className"`
	OrderID            uuid.UUID `json:"orderId"`
	OrderName          string    `json:"orderName"`
	FamilyID           uuid.UUID `json:"familyId"`
	FamilyName         string    `json:"familyName"`
	GenusID            uuid.UUID `json:"genusId"`
	GenusName          string    `json:"genusName"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
}

type SpeciesDetailsResponse struct {
	BaseResponse
	Data SpeciesDetailsDto `json:"data"`
}

// type SpeciesSearchQuery struct {
// 	Keyword string `json:"keyword"`
// 	Limit   int    `json:"limit"`
// }

// type SpeciesSearchResponse struct {
// 	BaseResponse
// 	Data []SpeciesListDto `json:"data"`
// }
