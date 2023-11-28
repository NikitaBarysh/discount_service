package entity

import "time"

type Withdraw struct {
	Number     string    `json:"order" binding:"required"`
	Sum        string    `json:"sum" binding:"required"`
	Status     string    `json:"status,omitempty"`
	UploadedAt time.Time `json:"uploaded_at,omitempty" db:"uploaded_at"`
}
