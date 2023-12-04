package repository

import (
	"fmt"

	"github.com/NikitaBarysh/discount_service.git/internal/entity"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(newDB *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: newDB}
}

func (r *AuthPostgres) CreateUser(user entity.User) (int, error) {
	var id int
	tx, err := r.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("err to beginTx: %w", err)
	}

	_, err = tx.Exec(insertUser, user.Login, user.Password)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return 0, fmt.Errorf("error to do rollback: %w", err)
		}
		return 0, err
	}

	row := tx.QueryRow(getUserIDByLogin, user.Login)
	err = row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("err to scan: %w", err)
	}

	return id, tx.Commit()
}

func (r *AuthPostgres) GetUserIDByLogin(login string) (int, error) {
	var userID int
	//SELECT id FROM users WHERE login=$1
	fmt.Println("rep login: ", login)
	err := r.db.Get(&userID, getUserIDByLogin, login)
	fmt.Println("rep err: ", err)
	fmt.Println("userID id: ", userID)
	if err != nil {
		return 0, fmt.Errorf("err to get id: %w", err)
	}
	return userID, nil
}

func (r *AuthPostgres) GetUser(login, password string) (entity.User, error) {
	var user entity.User
	err := r.db.Get(&user, getUser, login, password)
	if err != nil {
		return entity.User{}, fmt.Errorf("err to get user form db: %w", err)
	}

	return user, nil
}
