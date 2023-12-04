package service

import (
	"fmt"

	"github.com/NikitaBarysh/discount_service.git/internal/entity"
	"github.com/NikitaBarysh/discount_service.git/internal/repository"
)

type OrderService struct {
	rep repository.Order
}

func NewOrderService(newRep *repository.Repository) *OrderService {
	return &OrderService{rep: newRep}
}

func (s *OrderService) CreateOrder(order entity.Order) error {
	err := s.rep.CreateOrder(order)
	fmt.Println("create order: ", err)
	if err != nil {
		return fmt.Errorf("create order in DB: %s", err)
	}
	return nil
}

func (s *OrderService) GetUserIDByLogin(login string) (int, error) {
	userID, err := s.rep.GetUserIDByLogin(login)
	if err != nil {
		return 0, fmt.Errorf("get ID from DB: %w", err)
	}
	return userID, nil
}

func (s *OrderService) GetOrders(userID int) ([]entity.Order, error) {
	orders, err := s.rep.GetOrders(userID)
	if err != nil {
		return nil, fmt.Errorf("GetOrders from DB: %w", err)
	}

	return orders, nil
}

func (s *OrderService) CheckNumber(number string) error {

	numDB := s.rep.CheckNumber(number)
	if numDB != 0 {
		return fmt.Errorf("order already exist")
	}

	return nil
}

func (s *OrderService) LuhnAlgorithm(number int) bool {
	return (number%10+checksum(number/10))%10 == 0
}

func checksum(number int) int {
	var luhn int

	for i := 0; number > 0; i++ {
		cur := number % 10

		if i%2 == 0 {
			cur = cur * 2
			if cur > 9 {
				cur = cur%10 + cur/10
			}
		}

		luhn += cur
		number = number / 10
	}
	return luhn % 10
}
