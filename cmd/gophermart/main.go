package main

import (
	"context"
	"fmt"
	"github.com/NikitaBarysh/discount_service.git/internal/handler"

	"github.com/NikitaBarysh/discount_service.git/configs"
	"github.com/NikitaBarysh/discount_service.git/internal/app"
	"github.com/NikitaBarysh/discount_service.git/internal/repository"
	"github.com/NikitaBarysh/discount_service.git/internal/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	cfg := configs.NewConfig()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := repository.NewPostgresDB(ctx, cfg.DataBase)
	if err != nil {
		logrus.Fatalf("main: NewPostgresDB: %s", err.Error())
	}
	storage := repository.NewRepository(db)
	newService := service.NewService(storage)
	handlers := handler.NewHandler(newService)
	//
	work := service.NewWorkerPool(ctx, 6, storage.Order)

	go func() {
		service.NewOrderRequest(cfg.Accrual)
		fmt.Println("main: ", cfg.Accrual)
		work.Run(ctx)
	}()

	srv := new(app.Server)
	if err := srv.Run("8000", handlers.InitRouters()); err != nil {
		logrus.Fatalf("err while runnig server: %s", err.Error())
	}
}
