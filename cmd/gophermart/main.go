package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/NikitaBarysh/discount_service.git/cmd/gophermart/config"
	"github.com/NikitaBarysh/discount_service.git/internal/app"
	"github.com/NikitaBarysh/discount_service.git/internal/handler"
	"github.com/NikitaBarysh/discount_service.git/internal/repository"
	"github.com/NikitaBarysh/discount_service.git/internal/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	cfg := config.NewServer()
	logrus.Info("project config: ", cfg)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	m, err := repository.RunMigration(cfg.DataBase)
	if err != nil && !m {
		log.Fatal(err)
	}

	db, err := repository.NewPostgresDB(ctx, cfg.DataBase)
	logrus.Info("database path: ", cfg.DataBase)
	if err != nil {

		logrus.Error("main: NewPostgresDB: %w", err)
	}
	storage := repository.NewRepository(db)
	newService := service.NewService(storage)
	handlers := handler.NewHandler(newService)

	work := service.NewWorkerPool(6, storage.Order, cfg.Accrual)

	go func() {
		work.Run(ctx)
	}()

	srv := new(app.Server)
	go func() {
		if err := srv.Run(":8000", handlers.InitRouters()); err != nil {
			logrus.Error("err while running server: ", err)
		}
	}()
	logrus.Info("server started with port: ", cfg.Endpoint)

	termSig := make(chan os.Signal, 1)
	signal.Notify(termSig, syscall.SIGTERM, syscall.SIGINT)
	<-termSig

	logrus.Info("Shutting Down")

	if err := srv.Shutdown(ctx); err != nil {
		logrus.Error("err to shut down")
	}

}
