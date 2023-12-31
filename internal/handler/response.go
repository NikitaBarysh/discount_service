package handler

import "time"

type ResponseOrder struct {
	Number     string    `json:"number"`
	Status     string    `json:"status"`
	Accrual    float64   `json:"accrual,omitempty"`
	UploadedAt time.Time `json:"uploaded_at"`
}

type ResponseWithdraw struct {
	OrderNumber string    `json:"order"`
	Sum         float64   `json:"sum"`
	UploadedAt  time.Time `json:"uploaded_at"`
}

type ResponseBalance struct {
	Current  float64 `json:"current"`
	Withdraw float64 `json:"withdrawn"`
}

type InputWithdraw struct {
	Number string  `json:"order" binding:"required"`
	Sum    float64 `json:"sum" binding:"required"`
}
