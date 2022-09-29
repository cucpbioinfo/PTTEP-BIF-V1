package services

import (
	"github.com/wichadak/eDNA/repositories"
	"github.com/wichadak/eDNA/types"
)

type OrderService struct {
	OrderRepository *repositories.OrderRepository
}

func NewOrderService(orderRepository *repositories.OrderRepository) *OrderService {
	return &OrderService{
		OrderRepository: orderRepository,
	}
}

func (orderService *OrderService) ListOrder(query types.ListOrderQuery) ([]types.OrderDto, error) {
	orders, err := orderService.OrderRepository.ListOrder(query)
	if err != nil {
		return make([]types.OrderDto, 0), err
	}
	orderList := make([]types.OrderDto, len(orders))
	for i, order := range orders {
		orderList[i] = types.OrderDto{
			OrderID:   order.OrderID,
			OrderName: order.OrderName,
		}
	}
	return orderList, nil
}
