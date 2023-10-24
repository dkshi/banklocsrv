package main

import (
	"os"

	"github.com/dkshi/banklocsrv"
	"github.com/dkshi/banklocsrv/internal/handler"
	"github.com/dkshi/banklocsrv/internal/repository"
	"github.com/dkshi/banklocsrv/internal/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading environment: %s", err.Error())
	}

	db, err := repository.NewMongoDB(repository.Config{
		AuthSource: viper.GetString("db.authsource"),
		Host:       viper.GetString("db.host"),
		Port:       viper.GetString("db.port"),
		Username:   viper.GetString("db.username"),
		Password:   os.Getenv("DB_PASSWORD"),
		DBName:     viper.GetString("db.dbname"),
	})

	if err != nil {
		logrus.Fatalf("error connecting database: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	err = repo.FillCollections()
	if err != nil {
		logrus.Fatalf("error while filling collections: %s", err.Error())
	}

	go repo.ImitateLoad()

	srv := new(banklocsrv.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error while running app: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
