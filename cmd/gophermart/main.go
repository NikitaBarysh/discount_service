package main

import (
	"context"
	"fmt"
	"github.com/NikitaBarysh/discount_service.git/configs"
	"github.com/NikitaBarysh/discount_service.git/internal/app"
	"github.com/NikitaBarysh/discount_service.git/internal/handler"
	"github.com/NikitaBarysh/discount_service.git/internal/repository"
	"github.com/NikitaBarysh/discount_service.git/internal/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	cfg := configs.NewServer()
	fmt.Println(cfg)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := repository.NewPostgresDB(ctx, cfg.DataBase)
	if err != nil {
		logrus.Error("main: NewPostgresDB: %s", err.Error())
	}
	storage := repository.NewRepository(db)
	newService := service.NewService(storage)
	handlers := handler.NewHandler(newService)
	service.NewOrderRequest(cfg.Accrual)
	work := service.NewWorkerPool(ctx, 6, storage.Order)

	go func() {
		work.Run(ctx)
	}()

	srv := new(app.Server)
	if err := srv.Run(cfg.Endpoint, handlers.InitRouters()); err != nil {
		logrus.Error("err while running server: %s", err.Error())
	}
}
