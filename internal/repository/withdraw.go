package repository

import (
	"fmt"
	"time"

	"github.com/NikitaBarysh/discount_service.git/internal/entity"
	"github.com/jmoiron/sqlx"
)

type WithdrawRepository struct {
	db *sqlx.DB
}

func NewWithdrawRepository(newDB *sqlx.DB) *WithdrawRepository {
	return &WithdrawRepository{db: newDB}
}

func (r *WithdrawRepository) GetUserBalance(userID int) (entity.Balance, error) {
	var balance entity.Balance

	row := r.db.QueryRow(getBalance, userID)
	err := row.Scan(&balance.Money, &balance.Bonus)
	if err != nil {
		return balance, err
	}

	return balance, nil
}

func (r *WithdrawRepository) SetWithdraw(withdraw entity.Withdraw, userID int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("err to begin tx: %w", err)
	}
	defer tx.Rollback()

	var enough bool

	row := tx.QueryRow(getWithdraw, withdraw.Sum, userID)

	if row.Err() != nil {
		return fmt.Errorf("err to do query: %w", row.Err())
	}

	err = row.Scan(&enough)

	if err != nil {
		return fmt.Errorf("err to do Scan: %w", err)
	}

	if !enough {
		return entity.ErrNotEnoughMoney
	}

	_, err = tx.Exec(updateBalance, withdraw.Sum, userID)

	if err != nil {
		return fmt.Errorf("err to update balance: %w", err)
	}

	_, err = tx.Exec(insertWithdraw, withdraw.Number, userID, withdraw.Sum, "PROCESSED", time.Now())

	if err != nil {
		return fmt.Errorf("err to insert Withdraw: %w", err)
	}

	return tx.Commit()
}

func (r *WithdrawRepository) GetWithdraw(userID int) ([]entity.Withdraw, error) {
	allWithdraw := make([]entity.Withdraw, 0)

	err := r.db.Select(&allWithdraw, getAllWithdraw, userID)

	if err != nil {
		return nil, fmt.Errorf("err to get Withdraw: %w", err)
	}

	return allWithdraw, nil
}
