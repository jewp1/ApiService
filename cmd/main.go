package main

import (
	"ApiService/internal/api"
	"ApiService/internal/config"
	"ApiService/internal/logger"
	"ApiService/internal/repo"
	"ApiService/internal/service"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	if err := godotenv.Load(config.EnvPath); err != nil {
		log.Fatal("Ошибка загрузки env файла:", err)
	}

	var cfg config.Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(errors.Wrap(err, "fail to load config"))
	}

	newRepository, err := repo.NewRepo()
	if err != nil {
		log.Fatal(errors.Wrap(err, "fail to init repository"))
	}

	logger, err := logger.NewLogger(cfg.LogLevel)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error init logger"))
	}

	newService := service.NewService(newRepository, logger)

	app := api.NewRouter(&api.Router{Service: newService}, cfg.Token)

	go func() {
		logger.Infof("Starting http server on %s", cfg.ListenAddr)
		if err := app.Listen(cfg.ListenAddr); err != nil {
			logger.Fatal(err, "error while starting http server")
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

	logger.Info("Shutting down gracefully...")

}
