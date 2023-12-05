package repository

const (
	insertUser        = `INSERT INTO users (login, password) VALUES ($1, $2) RETURNING id`
	getUser           = `SELECT id FROM users WHERE  login=$1 AND password=$2`
	insertOrder       = `INSERT INTO orders (user_id, number, status, accrual) VALUES ($1, $2, $3, $4)`
	getOrder          = `SELECT id FROM orders WHERE number=$1`
	getOrders         = `SELECT number, status, accrual, uploaded_at FROM orders WHERE user_id=$1 ORDER BY uploaded_at`
	getBalance        = `SELECT current, withdraw FROM users WHERE id=$1`
	insertWithdraw    = `INSERT INTO withdraws (number, user_id, sum, status, uploaded_at) VALUES ($1, $2, $3, $4, $5)`
	getAllWithdraw    = `SELECT number, sum, uploaded_at FROM withdraws WHERE user_id=$1`
	getNewOrder       = `SELECT number, user_id FROM orders WHERE status='NEW'`
	updateUserBalance = `UPDATE  users SET current=current + $1 WHERE id=$2`
	updateOrderStatus = `UPDATE orders SET status=$1, accrual=$2 WHERE number=$3`
	getUserIDByLogin  = `SELECT id FROM users WHERE login=$1`
	getUserOrder      = `SELECT id FROM orders WHERE user_id=$1 AND number=$2`
)
