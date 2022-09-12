package dao

import (
	"context"
	"errors"
	"log"

	"student-management/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StudentDAO struct {
	Server     string
	Database   string
	Collection string
}


var Collection *mongo.Collection
var ctx = context.TODO()

//Basically this method establish connection between mongodb with go through mongodb driver.
//And we get the Collection of connected db
func (e *StudentDAO) Connect() {
	clientOptions := options.Client().ApplyURI(e.Server)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	Collection = client.Database(e.Database).Collection(e.Collection)
}

func (e *StudentDAO) Insert(student model.Student) error {
	_, err := Collection.InsertOne(ctx, student)

	if err != nil {
		return errors.New("unable to create new record")
	}

	return nil
}