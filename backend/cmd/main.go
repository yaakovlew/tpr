package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"backend"
	"backend/pkg/config"
	"backend/pkg/handler"
	"backend/pkg/integrations"
	"backend/pkg/repository"
	"backend/pkg/service"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// @title backend
// @version 1.0
// @description API Sever for Mephi Application

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuthStudent
// @in header
// @name AuthorizationStudent

// @securityDefinitions.apikey ApiKeyAuthLecturer
// @in header
// @name AuthorizationLecturer

// @securityDefinitions.apikey ApiKeyAuthSeminarian
// @in header
// @name AuthorizationSeminarian

// @securityDefinitions.apikey ApiKeyAuthCommon
// @in header
// @name AuthorizationCommon
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := config.InitConfig(); err != nil {
		logrus.Fatalf("Failed to load configuration: %v", err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     config.AppConfig.DB.Host,
		Port:     config.AppConfig.DB.Port,
		Username: config.AppConfig.DB.Username,
		Password: config.AppConfig.DB.Password,
		DBName:   config.AppConfig.DB.DBName,
		SSLMode:  config.AppConfig.DB.SSLMode,
	})
	if err != nil {
		logrus.Fatalf("Fatal to connect to DB, because: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewController(services)

	integrations.CronRun()

	srv := new(backend.Server)

	go func() {
		if err := srv.Run(handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Problem with start server, because %s", err.Error())
			return
		}
	}()
	logrus.Println("backend started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Println("backend shutting down")

	if err := srv.ShutDown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
		return
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
		return
	}
}
