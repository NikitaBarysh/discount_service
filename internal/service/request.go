package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/NikitaBarysh/discount_service.git/internal/entity"
)

type OrderRequest struct {
	AccrualHost string
}

type OrderResponse struct {
	Order   string  `json:"order"`
	Status  string  `json:"status"`
	Accrual float64 `json:"accrual"`
}

func RequestToAccrual(number, accrual string) (OrderResponse, error) {
	url := fmt.Sprintf("%s/api/orders/%s", accrual, number)

	response, err := http.Get(url)
	if err != nil {
		return OrderResponse{}, fmt.Errorf("err to get reposnse from Accrual: %w", err)
	}

	if response.StatusCode == http.StatusInternalServerError {
		return OrderResponse{}, fmt.Errorf("problem with Accrual service: %w", err)
	}

	if response.StatusCode == http.StatusNotFound || response.StatusCode == http.StatusNoContent {
		return OrderResponse{
			Order:  number,
			Status: "INVALID",
		}, nil
	}

	if response.StatusCode == http.StatusTooManyRequests {
		return OrderResponse{}, entity.ErrTooManyRequest
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

	if res.Status == "REGISTERED" || res.Status == "PROCESSING" {
		return OrderResponse{}, fmt.Errorf("order on process")
	}

	return res, nil
}
