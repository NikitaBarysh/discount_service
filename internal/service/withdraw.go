package service

import (
	"errors"
	"fmt"

	"github.com/NikitaBarysh/discount_service.git/internal/entity"
	"github.com/NikitaBarysh/discount_service.git/internal/repository"
)

type WithdrawService struct {
	rep repository.Withdraw
}

func NewWithdrawService(newRep *repository.Repository) *WithdrawService {
	return &WithdrawService{rep: newRep}
}

func (s *WithdrawService) GetBalance(userID int) (entity.Balance, error) {
	balance, err := s.rep.GetBalance(userID)
	if err != nil {
		return entity.Balance{}, fmt.Errorf("err to GetBalnca from DB: %w", err)
	}

	return balance, nil
}

func (s *WithdrawService) SetWithdraw(withdraw entity.Withdraw, userID int) error {
	err := s.rep.SetWithdraw(withdraw, userID)
	if err != nil {
		if errors.Is(err, entity.ErrNotEnoughMoney) {
			return entity.ErrNotEnoughMoney
		}
		return fmt.Errorf("err to Set in DB: %w", err)
	}

	return nil
}

func (s *WithdrawService) GetWithdraw(userID int) ([]entity.Withdraw, error) {
	res, err := s.rep.GetWithdraw(userID)
	if err != nil {
		return nil, fmt.Errorf("err GetWithdraw from DB: %w", err)
	}

	return res, nil
}
