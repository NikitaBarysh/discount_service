package migrations

//
//import (
//	"database/sql"
//	"fmt"
//
//	"github.com/pressly/goose"
//)
//
//func init() {
//	goose.AddMigration(upUserTable, downUserTable)
//	goose.AddMigration(upOrderTable, downOrderTable)
//	goose.AddMigration(upWithdrawsTable, downWithdrawsTable)
//
//}
//
//func upUserTable(tx *sql.Tx) error {
//	query := `CREATE TABLE users(
//    id  SERIAL PRIMARY KEY ,
//    login varchar(50) NOT NULL UNIQUE ,
//    password varchar NOT NULL,
//    current FLOAT,
//    withdraw INT);`
//	_, err := tx.Exec(query)
//	if err != nil {
//		return fmt.Errorf("migrations: upTable: %w", err)
//	}
//	return nil
//}
//
//func downUserTable(tx *sql.Tx) error {
//	query := `DROP TABLE users`
//	_, err := tx.Exec(query)
//	if err != nil {
//		return fmt.Errorf("migrations: downTable: %w", err)
//	}
//	return nil
//}
//
//func upOrderTable(tx *sql.Tx) error {
//	query := `CREATE TABLE orders(
//    id SERIAL PRIMARY KEY ,
//    user_id INT REFERENCES users(id)  ,
//    number VARCHAR UNIQUE,
//    status VARCHAR,
//    uploaded_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
//    accrual FLOAT DEFAULT 0
//);`
//	_, err := tx.Exec(query)
//	if err != nil {
//		return fmt.Errorf("migrations: upTable: %w", err)
//	}
//	return nil
//}
//
//func downOrderTable(tx *sql.Tx) error {
//	query := `DROP TABLE orders`
//	_, err := tx.Exec(query)
//	if err != nil {
//		return fmt.Errorf("migrations: downTable: %w", err)
//	}
//	return nil
//}
//
//func upWithdrawsTable(tx *sql.Tx) error {
//	query := `CREATE TABLE withdraws(
//    id SERIAL PRIMARY KEY,
//    user_id INT REFERENCES users(id),
//    number VARCHAR(50) NOT NULL UNIQUE,
//    status VARCHAR(30) DEFAULT 'NEW',
//    sum FLOAT,
//    uploaded_at TIMESTAMP
//);`
//	_, err := tx.Exec(query)
//	if err != nil {
//		return fmt.Errorf("migrations: upTable: %w", err)
//	}
//	return nil
//}
//
//func downWithdrawsTable(tx *sql.Tx) error {
//	query := `DROP TABLE withdraws`
//	_, err := tx.Exec(query)
//	if err != nil {
//		return fmt.Errorf("migrations: downTable: %w", err)
//	}
//	return nil
//}
