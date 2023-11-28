package repository

import (
	"fmt"

	"github.com/NikitaBarysh/discount_service.git/internal/entity"
	"github.com/jmoiron/sqlx"
)

type OrderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(newDB *sqlx.DB) *OrderRepository {
	return &OrderRepository{db: newDB}
}

func (r *OrderRepository) CreateOrder(order entity.Order) error {
	_, errInsert := r.db.Exec(insertOrder, order.UserId, order.Number, order.Status, order.Accrual)

	if errInsert != nil {
		return fmt.Errorf("err to do insert into order db")
	}

	return nil
}

func (r *OrderRepository) GetOrders(userID int) ([]entity.Order, error) {
	orderSlice := make([]entity.Order, 0)
	err := r.db.Select(&orderSlice, getOrders, userID)

	if err != nil {
		return nil, fmt.Errorf("can't get data")
	}

	return orderSlice, nil
}

func (r *OrderRepository) CheckNumber(number string) bool {
	var order entity.Order
	r.db.Get(&order.UserId, getOrder, number)

	if order.UserId == 0 {
		return true
	}

	return false
}
