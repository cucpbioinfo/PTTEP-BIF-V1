package types

import (
	"github.com/google/uuid"
)

type MajorGroupDto struct {
	MajorGroupID   uuid.UUID `json:"majorGroupId"`
	MajorGroupName string    `json:"majorGroupName"`
}

type ListMajorGroupResponse struct {
	BaseResponse
	Data []MajorGroupDto `json:"data"`
}
