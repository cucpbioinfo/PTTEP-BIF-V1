package types

import (
	"github.com/google/uuid"
)

type OrderCreateBody struct {
	OrderName string `json:"orderName"`
}

type OrderDto struct {
	OrderID   uuid.UUID `json:"orderId"`
	OrderName string    `json:"orderName"`
}

type ListOrderQuery struct {
	ClassID uuid.UUID `json:"classId"`
}

type ListOrderResponse struct {
	BaseResponse
	Data []OrderDto `json:"data"`
}
