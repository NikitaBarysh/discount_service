package service

import (
	"github.com/NikitaBarysh/discount_service.git/internal/entity"
	"github.com/NikitaBarysh/discount_service.git/internal/repository"
)

//go:generate mockgen -source ${GOFILE} -destination mock.go -package ${GOPACKAGE}

type Authorization interface {
	CreateUser(user entity.User) error
	GenerateToken(user entity.User) (string, error)
	ParseToken(token string) (int, error)
	ValidateLogin(user entity.User) error
	CheckData(user entity.User) error
}

type Service struct {
	Authorization
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(rep),
	}
}
