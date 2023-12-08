package entity

import "time"

type Order struct {
	ID         int       `json:"id,omitempty"`
	UserID     int       `json:"user_id,omitempty"`
	Number     string    `json:"number"`
	Status     string    `json:"status"`
	Accrual    int       `json:"accrual,omitempty"`
	UploadedAt time.Time `json:"uploaded_at" db:"uploaded_at"`
}
