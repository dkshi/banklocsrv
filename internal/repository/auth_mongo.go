package repository

import (
	"context"

	"github.com/dkshi/banklocsrv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthMongo struct {
	usersCollection *mongo.Collection
}

func NewAuthPostgres(uc *mongo.Collection) *AuthMongo {
	return &AuthMongo{usersCollection: uc}
}

func (a *AuthMongo) CreateUser(user banklocsrv.User) (string, error) {
	userBSON, err := bson.Marshal(user)

	if err != nil {
		return "", err
	}

	cursor, err := a.usersCollection.InsertOne(context.TODO(), userBSON)

	if err != nil {
		return "", err
	}

	return cursor.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (a *AuthMongo) GetUser(username, password string) (banklocsrv.User, error) {
	var user banklocsrv.User
	err := a.usersCollection.FindOne(context.TODO(), bson.M{"username": username, "password": password}).Decode(&user)

	if err != nil {
		return banklocsrv.User{}, err
	}

	return user, nil
}
