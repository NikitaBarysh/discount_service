package service

import (
	"github.com/NikitaBarysh/discount_service.git/internal/entity"
	"github.com/NikitaBarysh/discount_service.git/internal/repository"
)

//go:generate mockgen -source ${GOFILE} -destination mock.go -package ${GOPACKAGE}

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GenerateToken(userID int) (string, error)
	ParseToken(token string) (int, error)
	ValidateLogin(user entity.User) error
	CheckData(user entity.User) error
	GetUser(userData entity.User) (entity.User, error)
	GetUserIDByLogin(login string) (int, error)
}

type Order interface {
	LuhnAlgorithm(num int) bool
	CreateOrder(user entity.Order) error
	CheckNumber(number string) error
	GetOrders(userID int) ([]entity.Order, error)
}

type Withdraw interface {
	GetBalance(userID int) (entity.Balance, error)
	SetWithdraw(withdraw entity.Withdraw, userID int) error
	GetWithdraw(userID int) ([]entity.Withdraw, error)
}

type Service struct {
	Authorization
	Order
	Withdraw
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(rep),
		Order:         NewOrderService(rep),
		Withdraw:      NewWithdrawService(rep),
	}
}
