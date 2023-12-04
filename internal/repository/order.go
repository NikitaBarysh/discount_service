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
	_, errInsert := r.db.Exec(insertOrder, order.UserID, order.Number, order.Status, order.Accrual)

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

func (r *OrderRepository) CheckNumber(number string) int {
	var order entity.Order
	r.db.Get(&order.ID, getOrder, number)

	return order.ID
}

func (r *OrderRepository) CheckUserOrder(userID int, number string) int {
	var orderID int
	r.db.Get(&orderID, getUserOrder, userID, number)

	return orderID
}

func (r *OrderRepository) GetUserIDByLogin(login string) (int, error) {
	var userID int
	err := r.db.Get(&userID, getUserIDByLogin, login)
	if err != nil {
		return 0, fmt.Errorf("err to get id: %w", err)
	}

	return userID, nil
}

func (r *OrderRepository) GetNewOrder() ([]entity.UpdateStatus, error) {
	number := make([]entity.UpdateStatus, 0)
	err := r.db.Select(&number, getNewOrder)

	if err != nil {
		return nil, fmt.Errorf("err to get number: %w", err)
	}

	return number, nil
}

func (r *OrderRepository) UpdateStatus(response entity.UpdateStatus) error {
	tx, err := r.db.Begin()

	if err != nil {
		return fmt.Errorf("err to begin TX: %w", err)
	}

	_, err = tx.Exec(updateUserBalance, response.Accrual, response.UserID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("err to update user balance: %w", err)
	}

	_, err = tx.Exec(updateOrderStatus, response.Status, response.Order)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("err to update order stattus")
	}

	return tx.Commit()
}
