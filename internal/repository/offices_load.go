package repository

import (
	"context"
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OfficeLoadImitator struct {
	officesCollection *mongo.Collection
}

func NewOfficeLoadImitator(oc *mongo.Collection) *OfficeLoadImitator {
	return &OfficeLoadImitator{officesCollection: oc}
}

func (o *OfficeLoadImitator) ImitateLoad() {
	if err := o.ClearLoad(); err != nil {
		logrus.Printf("error cleaning collection before starting app: %s", err.Error())
	}

	for {
		go func() {
			cursor, err := o.officesCollection.Aggregate(context.TODO(), mongo.Pipeline{
				bson.D{{Key: "$sample", Value: bson.D{{Key: "size", Value: 20}}}},
			})

			if err != nil {
				logrus.Printf("error selecting random documents from collection, can't imitate load: %s", err.Error())
				return
			}

			var offices []bson.M
			if err := cursor.All(context.TODO(), &offices); err != nil {
				logrus.Printf("error decoding bson: %s", err.Error())
			}

			for _, office := range offices {
				go func(office bson.M) {
					if err := o.UpdateRow(office, int(office["load"].(int32)+1)); err != nil {
						logrus.Printf("error incrementing row: %s", err.Error())
					}

					time.Sleep(time.Duration(rand.Intn(6)+5) * time.Minute)

					if err := o.UpdateRow(office, int(office["load"].(int32)-1)); err != nil {
						logrus.Printf("error decrementing row: %s", err.Error())
					}
				}(office)
			}
		}()
		time.Sleep(time.Duration(rand.Intn(30)+1) * time.Second)
	}
}

// Id is a filter like this: bson.M{"_id": "<any id>"}
func (o *OfficeLoadImitator) UpdateRow(office bson.M, amount int) error {
	filter := bson.M{"_id": office["_id"]}
	update := bson.M{"$set": bson.M{"load": amount}}

	if _, err := o.officesCollection.UpdateOne(context.TODO(), filter, update); err != nil {
		return err
	}

	return nil
}

// Clear load before starting application
func (o *OfficeLoadImitator) ClearLoad() error {
	cursor, err := o.officesCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		return err
	}

	var offices []bson.M
	cursor.All(context.TODO(), &offices)

	for _, office := range offices {
		if err := o.UpdateRow(office, 0); err != nil {
			return err
		}
	}

	return nil
}
