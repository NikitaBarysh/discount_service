package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type OrderRequest struct {
	AccrualHost string
}

func NewOrderRequest(cfg string) *OrderRequest {
	return &OrderRequest{AccrualHost: cfg}
}

type OrderResponse struct {
	Order   string  `json:"order"`
	Status  string  `json:"status"`
	Accrual float64 `json:"accrual"`
}

func (s *OrderRequest) RequestToAccrual(number string) (OrderResponse, error) {
	fmt.Println("4")
	url := s.AccrualHost
	fmt.Println("url", url)
	time.Sleep(time.Second * 3)
	fmt.Println("http://localhost:8080/" + number)
	response, err := http.Get("http://localhost:8080/api/orders/" + number)
	if err != nil {
		return OrderResponse{}, fmt.Errorf("err to get reposnse from Accrual: %w", err)
	}

	if response.StatusCode == http.StatusInternalServerError {
		return OrderResponse{}, fmt.Errorf("problem with Accrual service: %w", err)
	}

	if response.StatusCode == http.StatusNotFound {
		return OrderResponse{}, fmt.Errorf("can't find number: %w", err)
	}

	if response.StatusCode == http.StatusTooManyRequests {
		return OrderResponse{}, fmt.Errorf("to many request: %w", err)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return OrderResponse{}, fmt.Errorf("err to read body: %w", err)
	}
	defer response.Body.Close()

	var res OrderResponse

	err = json.Unmarshal(body, &res)
	if err != nil {
		return OrderResponse{}, fmt.Errorf("err to unmarshal body: %w", err)
	}

	return res, nil
}
