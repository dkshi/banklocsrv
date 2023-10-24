package repository

import (
	"context"
	"encoding/json"
	"os"

	"github.com/dkshi/banklocsrv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	atmsFilePath       = "data/atms.json"
)

type AtmsMongo struct {
	atmsCollection *mongo.Collection
}

func NewAtmsMongo(ac *mongo.Collection) *AtmsMongo {
	return &AtmsMongo{atmsCollection: ac}
}

func (a *AtmsMongo) GetAtms(filter *map[string]interface{}) (*banklocsrv.AtmsData, error) {
	var atms banklocsrv.AtmsData
	cursor, err := a.atmsCollection.Find(context.TODO(), *filter)

	for cursor.Next(context.TODO()) {
		var atm banklocsrv.Atm
		err = bson.Unmarshal(cursor.Current, &atm)
		if err != nil {
			return &banklocsrv.AtmsData{}, err
		}
		atms.Atms = append(atms.Atms, atm)
	}

	if err != nil {
		return &banklocsrv.AtmsData{}, err
	}

	return &atms, nil
}

func (a *AtmsMongo) FillAtms() error {
	cursor, err := a.atmsCollection.Find(context.TODO(), bson.M{})

	if cursor.RemainingBatchLength() != 0 {
		return nil
	}

	var atms banklocsrv.AtmsData

	atmsJson, err := os.ReadFile(atmsFilePath)

	if err != nil {
		return err
	}

	json.Unmarshal(atmsJson, &atms)

	for _, atm := range atms.Atms {
		_, err := a.atmsCollection.InsertOne(context.TODO(), atm)

		if err != nil {
			return err
		}
	}

	return nil
}
