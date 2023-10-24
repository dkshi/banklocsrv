# Bank Location Server
A server that provides information about VTB Bank branches. Developed for the MORE Tech 5.0 Hackathon.

The bot works with MongoDB, so it's necessary for the request body to contain valid JSON that MongoDB can process. For example:
```
{
    "services.blind.serviceActivity": "AVAILABLE"
}
```

**Server functions:***

- User authentication
- Automatic population of MongoDB collections
- Simulating the number of people in offices

**Available routes:** 

/auth
- POST /auth/sign-up - takes JSON with username, password and returns new user id
- POST /auth/sign-in - takes JSON with username and password and returns JWT-token if authentication was successful

/departments (requires JWT-token)
- POST /departments/atms - takes JSON with filters you want and returns correct atms (body can be empty)
- POST /departments/offices - takes JSON with filters you want and returns correct offices (body can be empty)

## Getting started

Clone this repository and edit configs/config.yml and .env:
```yaml
port: "8080"

db:
  authsource: "admin"
  username: "mongo"
  host: "localhost"
  port: "27017"
  dbname: "bankloc-db"
```

```.env
DB_PASSWORD=YOUR_DB_PASSWORD
```

And use go run command on main.go file:
```bash
$ go run cmd/main.go
```

## Load imitation

This script works with goroutines and allows simulating the number of people in offices at the current moment.
Every 1-30 seconds, a random selection of 20 offices is made, and one person enters each of them.
Each person will leave after a random number of minutes (between 5 and 10).

```.go
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
		return
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
				return
			}

			for _, office := range offices {
				go func(office bson.M) {
					if err := o.UpdateRow(office, int(office["load"].(int32)+1)); err != nil {
						logrus.Printf("error incrementing row: %s", err.Error())
						return
					}

					time.Sleep(time.Duration(rand.Intn(6)+5) * time.Minute)

					if err := o.UpdateRow(office, int(office["load"].(int32)-1)); err != nil {
						logrus.Printf("error decrementing row: %s", err.Error())
						return
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

```
