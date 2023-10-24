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
	officesFilePath       = "data/offices.json"
)

type OfficesMongo struct {
	officesCollection *mongo.Collection
}

func NewOfficesMongo(oc *mongo.Collection) *OfficesMongo {
	return &OfficesMongo{officesCollection: oc}
}

func (o *OfficesMongo) GetOffices(filter *map[string]interface{}) (*banklocsrv.OfficesData, error) {
	var offices banklocsrv.OfficesData
	cursor, err := o.officesCollection.Find(context.TODO(), *filter)

	for cursor.Next(context.TODO()) {
		var office banklocsrv.Office
		err = bson.Unmarshal(cursor.Current, &office)
		if err != nil {
			return &banklocsrv.OfficesData{}, err
		}
		offices.Offices = append(offices.Offices, office)
	}

	if err != nil {
		return &banklocsrv.OfficesData{}, err
	}

	return &offices, nil
}

func (o *OfficesMongo) FillOffices() error {
	cursor, err := o.officesCollection.Find(context.TODO(), bson.M{})

	if cursor.RemainingBatchLength() != 0 {
		return nil
	}

	var offices banklocsrv.OfficesData

	officesJson, err := os.ReadFile(officesFilePath)

	if err != nil {
		return err
	}

	json.Unmarshal(officesJson, &offices)

	for _, office := range offices.Offices {
		_, err := o.officesCollection.InsertOne(context.TODO(), office)

		if err != nil {
			return err
		}
	}

	return nil
}
