package types

import (
	"github.com/google/uuid"
)

type UserCreateBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserCreateResponse struct {
	BaseResponse
}

type UserGetResponse struct {
	UserID uuid.UUID `json:"userId"`
}
