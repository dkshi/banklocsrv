package repository

import (
	"github.com/dkshi/banklocsrv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	officesCollectionName = "offices"
	usersCollectionName   = "users"
	atmsCollectionName    = "atms"
)

type Authorization interface {
	CreateUser(user banklocsrv.User) (string, error)
	GetUser(username, password string) (banklocsrv.User, error)
}

type Atms interface {
	GetAtms(*map[string]interface{}) (*banklocsrv.AtmsData, error)
	FillAtms() error
}

type Offices interface {
	GetOffices(filter *map[string]interface{}) (*banklocsrv.OfficesData, error)
	FillOffices() error
}

// Load imitation
type OfficesLoad interface {
	ImitateLoad()
	UpdateRow(office bson.M, amount int) error
}

type Repository struct {
	Authorization
	Atms
	Offices
	OfficesLoad
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db.Collection(usersCollectionName)),
		Atms:          NewAtmsMongo(db.Collection(atmsCollectionName)),
		Offices:       NewOfficesMongo(db.Collection(officesCollectionName)),
		OfficesLoad:   NewOfficeLoadImitator(db.Collection(officesCollectionName)),
	}
}

func (repo *Repository) FillCollections() error {
	if err := repo.FillAtms(); err != nil {
		return err
	}

	if err := repo.FillOffices(); err != nil {
		return err
	}

	return nil
}
