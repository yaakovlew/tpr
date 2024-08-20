package main

import (
	"backend"
	"backend/pkg/handler"
	"backend/pkg/integrations"
	"backend/pkg/repository"
	"backend/pkg/service"
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
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
	if err := initConfig(); err != nil {
		logrus.Fatalf("error intializing config: %s", err.Error())
		return
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env varibles: %s", err.Error())
		return
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("Fatal to connect to DB, because: %s", err.Error())
		return
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

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
