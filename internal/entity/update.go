package entity

type UpdateStatus struct {
	UserID  int     `json:"user_id" db:"user_id"`
	Order   string  `json:"order" db:"number"`
	Status  string  `json:"status"`
	Accrual float64 `json:"accrual" db:"accrual"`
}
