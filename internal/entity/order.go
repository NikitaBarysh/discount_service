package entity

import "time"

type Order struct {
	Id         int       `json:"id,omitempty"`
	UserId     int       `json:"user_id,omitempty"`
	Number     string    `json:"number"`
	Status     string    `json:"status"`
	Accrual    float64   `json:"accrual,omitempty"`
	UploadedAt time.Time `json:"uploaded_at" db:"uploaded_at"`
}
