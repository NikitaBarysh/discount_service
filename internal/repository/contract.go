package repository

import (
	"github.com/NikitaBarysh/discount_service.git/internal/entity"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GetUser(login, password string) (int, error)
	GetUserIDByLogin(login string) (int, error)
}

type Order interface {
	CreateOrder(order entity.Order) error
	CheckNumber(number string) int
	GetOrders(userID int) ([]entity.Order, error)
	GetNewOrder() ([]entity.UpdateStatus, error)
	UpdateStatus(response entity.UpdateStatus) error
	CheckUserOrder(userID int, number string) int
}

type Withdraw interface {
	GetBalance(userID int) (entity.Balance, error)
	SetWithdraw(withdraw entity.Withdraw, userID int) error
	GetWithdraw(userID int) ([]entity.Withdraw, error)
}

type Repository struct {
	Authorization
	Order
	Withdraw
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Order:         NewOrderRepository(db),
		Withdraw:      NewWithdrawRepository(db),
	}
}
