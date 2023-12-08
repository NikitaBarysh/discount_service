package entity

type Status struct {
	UserID  int    `json:"user_id" db:"user_id"`
	Order   string `json:"order" db:"number"`
	Status  string `json:"status"`
	Accrual int    `json:"accrual" db:"current"`
}
