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

	_, err = tx.Exec(`INSERT INTO users (login, password) VALUES ($1, $2) RETURNING id`,
		user.Login, user.Password)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return 0, fmt.Errorf("error to do rollback: %w", err)
		}
		return 0, err
	}

	row := tx.QueryRow(`SELECT id FROM users WHERE login=$1`,
		user.Login)

	if row.Err() != nil {
		tx.Rollback()
		return 0, fmt.Errorf("err to get id: %w", row.Err())
	}

	err = row.Scan(&id)

	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return 0, fmt.Errorf("error to do rollback: %w", err)
		}
		return 0, fmt.Errorf("err to scan: %w", err)
	}

	return id, tx.Commit()
}

func (r *AuthPostgres) GetUserIDByLogin(login string) (int, error) {
	var userID int

	err := r.db.Get(&userID,
		`SELECT id FROM users WHERE login=$1`,
		login)

	if err != nil {
		return 0, fmt.Errorf("err to get id: %w", err)
	}
	return userID, nil
}

func (r *AuthPostgres) GetUser(login, password string) (int, error) {
	var id int
	tx, err := r.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("err to beginTx: %w", err)
	}

	row := r.db.QueryRow(`SELECT id FROM users WHERE  login=$1 AND password=$2`,
		login, password)

	if row.Err() != nil {
		tx.Rollback()
		return 0, fmt.Errorf("err to query: %w", row.Err())
	}
	err = row.Scan(&id)

	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("err to scan: %w", err)
	}

	return id, tx.Commit()
}
