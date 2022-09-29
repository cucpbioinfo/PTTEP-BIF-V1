package repositories

import (
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/wichadak/eDNA/models"
	"github.com/wichadak/eDNA/types"
)

type OrderRepository struct {
	pg *pg.DB
}

func NewOrderRepository(pg *pg.DB) *OrderRepository {
	return &OrderRepository{
		pg: pg,
	}
}

func (orderRepository *OrderRepository) ListOrder(query types.ListOrderQuery) ([]models.Order, error) {
	var orders []models.Order
	dbQuery := orderRepository.pg.Model(&orders)

	if query.ClassID != uuid.Nil {
		dbQuery.Where(`"order".class_id = ?`, query.ClassID)
	}

	err := dbQuery.Select()
	if err != nil {
		return make([]models.Order, 0), err
	}
	return orders, nil
}
