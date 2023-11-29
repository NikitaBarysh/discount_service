package service

import (
	"context"
	"fmt"
	"github.com/NikitaBarysh/discount_service.git/internal/entity"
	"github.com/NikitaBarysh/discount_service.git/internal/repository"
	"sync"
	"time"
)

type WorkerPool struct {
	ctx     context.Context
	workers int
	inputCH chan entity.UpdateStatus
	storage repository.Order
	request OrderRequest
}

func NewWorkerPool(ctx context.Context, workers int, rep repository.Order) *WorkerPool {
	return &WorkerPool{
		ctx:     ctx,
		workers: workers,
		inputCH: make(chan entity.UpdateStatus, 6),
		storage: rep,
	}
}

func (s *WorkerPool) Run(ctx context.Context) {
	fmt.Println("11")
	var wg sync.WaitGroup
	fmt.Println("22")
	for i := 0; i <= s.workers; i++ {
		fmt.Println("33")
		wg.Add(1)
		fmt.Println("44")
		go func() {
			for {
				fmt.Println("55")
				select {
				case update := <-s.inputCH:
					fmt.Println("2")
					err := s.storage.UpdateStatus(update)
					if err != nil {
						fmt.Println("err to do request into Accrual: ", err)
					}
				case <-ctx.Done():
					fmt.Println("66")
					return
				}
			}
			wg.Done()
		}()
	}
	fmt.Println("123123")
	sch := s.scheduler(ctx)
	wg.Wait()
	close(s.inputCH)
	defer sch.Stop()
}

func (s *WorkerPool) scheduler(ctx context.Context) *time.Ticker {
	fmt.Println("6")
	//ticker := time.NewTicker(time.Second * time.Duration(3))
	//	for {
	//		fmt.Println("7")
	//		select {
	//		case <-ticker.C:
	//			fmt.Println("sched")
	//			s.GetRequest()
	//		case <-ctx.Done():
	//			return nil
	//		}
	//	}

	ticker := time.NewTicker(time.Second * 5)
	go func() {
		for {
			fmt.Println(7)
			select {
			case <-ticker.C:
				fmt.Println("sched")
				s.GetRequest()
			case <-ctx.Done():
				fmt.Println("done")
				return
			}
		}
	}()
	return ticker
}

func (s *WorkerPool) set(res entity.UpdateStatus) {
	fmt.Println("set")
	s.inputCH <- res
}

func (s *WorkerPool) GetRequest() error {
	fmt.Println("get order")
	numbers, err := s.storage.GetNewOrder()
	if err != nil {
		return fmt.Errorf("err to get new order: %w", err)
	}

	for _, v := range numbers {
		res, err := s.request.RequestToAccrual(v.Order)
		if err != nil {
			fmt.Println(fmt.Errorf("err to do request: %w", err))
			continue
		}
		response := entity.UpdateStatus{
			UserID:  v.UserID,
			Order:   res.Order,
			Status:  res.Status,
			Accrual: res.Accrual,
		}

		s.set(response)
	}

	return nil
}
