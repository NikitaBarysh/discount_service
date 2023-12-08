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
	tx, err := r.db.Begin()

	if err != nil {
		return fmt.Errorf("err to begin TX: %w", err)
	}

	_, errInsert := tx.Exec(`INSERT INTO orders (user_id, number, status, accrual) VALUES ($1, $2, $3, $4)`,
		order.UserID, order.Number, order.Status, order.Accrual)

	if errInsert != nil {
		tx.Rollback()
		return fmt.Errorf("err to do insert into order db")
	}

	return tx.Commit()
}

func (r *OrderRepository) GetOrders(userID int) ([]entity.Order, error) {
	orderSlice := make([]entity.Order, 0)
	err := r.db.Select(&orderSlice,
		`SELECT number, status, accrual, uploaded_at FROM orders WHERE user_id=$1 ORDER BY uploaded_at`,
		userID)

	if err != nil {
		return nil, fmt.Errorf("can't get data")
	}

	return orderSlice, nil
}

func (r *OrderRepository) CheckNumber(number string) int {
	var order entity.Order
	r.db.Get(&order.ID, `SELECT id FROM orders WHERE number=$1`, number)

	return order.ID
}

func (r *OrderRepository) CheckUserOrder(userID int, number string) int {
	var orderID int
	r.db.Get(&orderID,
		`SELECT id FROM orders WHERE user_id=$1 AND number=$2`,
		userID, number)

	return orderID
}

func (r *OrderRepository) GetNewOrder() ([]entity.Status, error) {
	number := make([]entity.Status, 0)
	err := r.db.Select(&number,
		`SELECT number, user_id FROM orders WHERE status='NEW'`)

	if err != nil {
		return nil, fmt.Errorf("err to get number: %w", err)
	}

	return number, nil
}

func (r *OrderRepository) UpdateStatus(response entity.Status) error {
	tx, err := r.db.Begin()

	if err != nil {
		return fmt.Errorf("err to begin TX: %w", err)
	}

	_, err = tx.Exec(`UPDATE  users SET current=current + $1 WHERE id=$2`,
		response.Accrual, response.UserID)

	if err != nil {
		tx.Rollback()
		return fmt.Errorf("err to update user balance: %w", err)
	}

	accrual := response.Accrual
	_, err = tx.Exec(`UPDATE orders SET status=$1, accrual=$2 WHERE number=$3`,
		response.Status, accrual, response.Order)

	if err != nil {
		tx.Rollback()
		return fmt.Errorf("err to update order stattus")
	}

	return tx.Commit()
}
