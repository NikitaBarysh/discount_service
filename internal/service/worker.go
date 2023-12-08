package service

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/NikitaBarysh/discount_service.git/internal/entity"
	"github.com/NikitaBarysh/discount_service.git/internal/repository"
)

type WorkerPool struct {
	workers int
	inputCH chan entity.Status
	storage repository.Order
	Accrual string
}

func NewWorkerPool(workers int, rep repository.Order, accrual string) *WorkerPool {
	return &WorkerPool{
		workers: workers,
		inputCH: make(chan entity.Status, 6),
		storage: rep,
		Accrual: accrual,
	}
}

func (s *WorkerPool) Run(ctx context.Context) {

	var wg sync.WaitGroup

	for i := 0; i <= s.workers; i++ {

		wg.Add(1)

		go func() {
		out:
			for {
				select {
				case update := <-s.inputCH:
					err := s.storage.UpdateStatus(update)
					if err != nil {
						fmt.Println("err to do request into Accrual: ", err)
					}
					continue
				case <-ctx.Done():
					break out
				}
			}
			wg.Done()
		}()
	}
	sch := s.scheduler(ctx)
	wg.Wait()
	close(s.inputCH)
	defer sch.Stop()
}

func (s *WorkerPool) scheduler(ctx context.Context) *time.Ticker {

	ticker := time.NewTicker(time.Second * 5)
	go func() {
		for {

			select {
			case <-ticker.C:
				err := s.GetRequest()
				if err != nil {
					if errors.Is(err, entity.ErrTooManyRequest) {
						ticker.Reset(time.Second * 60)
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return ticker
}

func (s *WorkerPool) set(res entity.Status) {
	s.inputCH <- res
}

func (s *WorkerPool) GetRequest() error {
	numbers, err := s.storage.GetNewOrder()
	if err != nil {
		return fmt.Errorf("err to get new order: %w", err)
	}

	for _, v := range numbers {
		res, err := RequestToAccrual(v.Order, s.Accrual)

		if err != nil {
			if errors.Is(err, entity.ErrTooManyRequest) {
				return err
			}
			fmt.Println(fmt.Errorf("err to do request: %w", err))
			continue
		}
		response := entity.Status{
			UserID:  v.UserID,
			Order:   res.Order,
			Status:  res.Status,
			Accrual: int(res.Accrual * 100),
		}

		s.set(response)
	}

	return nil
}
