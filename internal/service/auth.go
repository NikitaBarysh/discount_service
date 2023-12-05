package service

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"time"

	"github.com/NikitaBarysh/discount_service.git/internal/entity"
	"github.com/NikitaBarysh/discount_service.git/internal/repository"
	"github.com/golang-jwt/jwt/v4"
)

const (
	salt       = "32po34hf982v29"
	tokenTTL   = time.Hour * 24
	signingKey = "232okc0andha298rudf23r03uc"
)

type claims struct {
	jwt.RegisteredClaims
	UserID int `json:"user_id"`
}

type AuthService struct {
	rep repository.Authorization
}

func NewAuthService(newRep *repository.Repository) *AuthService {
	return &AuthService{rep: newRep}
}

func (s *AuthService) CreateUser(user entity.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.rep.CreateUser(user)
}

func (s *AuthService) GetUser(userData entity.User) (entity.User, error) {
	user, err := s.rep.GetUser(userData.Login, generatePasswordHash(userData.Password))
	if err != nil {
		return entity.User{}, fmt.Errorf("GetUser: %w", err)
	}
	return user, nil
}

func (s *AuthService) GenerateToken(userID int) (string, error) {
	//user, err := s.rep.GetUser(userData.Login, generatePasswordHash(userData.Password))
	//
	//if err != nil {
	//	return "", fmt.Errorf("GetUser: %w", err)
	//}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
		},
		UserID: userID,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) GetUserIDByLogin(login string) (int, error) {
	userID, err := s.rep.GetUserIDByLogin(login)

	if err != nil {
		return 0, fmt.Errorf("get ID from DB: %w", err)
	}
	return userID, nil
}

func (s *AuthService) ParseToken(authToken string) (int, error) {
	token, err := jwt.ParseWithClaims(authToken, &claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return 0, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, fmt.Errorf("err to parse token: %w", err)
	}

	claims, ok := token.Claims.(*claims)
	if !ok {
		return 0, errors.New("wrong type of token claims")
	}
	fmt.Println("claims user id", claims.UserID)
	return claims.UserID, nil
}

func (s *AuthService) ValidateLogin(user entity.User) error {
	userFromDB, err := s.rep.GetUser(user.Login, generatePasswordHash(user.Password))
	if err != nil {
		return fmt.Errorf("get user: %w", err)
	}

	if userFromDB.Login == "" {
		return nil
	}

	return entity.ErrNotUniqueLogin
}

func (s *AuthService) CheckData(user entity.User) error {
	_, err := s.rep.GetUser(user.Login, generatePasswordHash(user.Password))
	if err != nil {
		return fmt.Errorf("get user: %w", err)
	}

	return nil
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
