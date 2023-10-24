package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	usersTable = "users"
)

type Config struct {
	AuthSource string
	Host       string
	Port       string
	Username   string
	Password   string
	DBName     string
}

func NewMongoDB(cfg Config) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", cfg.Host, cfg.Port))

	credential := options.Credential{
		AuthSource: cfg.AuthSource,
		Username:   cfg.Username,
		Password:   cfg.Password,
	}

	clientOptions.SetAuth(credential)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return &mongo.Database{}, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return &mongo.Database{}, err
	}

	return client.Database(cfg.DBName), nil
}

