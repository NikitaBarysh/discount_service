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
	"log"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	cfg := configs.ParseServerConfig()
	fmt.Println(cfg)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	m, err := repository.RunMigration(cfg.DatabaseDSN)
	if err != nil && !m {
		log.Fatal(err)
	}

	db, err := repository.NewPostgresDB(ctx, cfg.DatabaseDSN)
	if err != nil {
		fmt.Println("main: postgres")
		logrus.Error("main: NewPostgresDB: %w", err)
	}
	storage := repository.NewRepository(db)
	newService := service.NewService(storage)
	handlers := handler.NewHandler(newService)

	work := service.NewWorkerPool(ctx, 6, storage.Order, cfg.AccrualSystemAddr)

	go func() {
		work.Run(ctx)
	}()

	srv := new(app.Server)
	if err := srv.Run(cfg.RunAddr, handlers.InitRouters()); err != nil {
		fmt.Println("main: run", cfg.RunAddr)
		logrus.Error("err while running server: %w", err)
	}
}
