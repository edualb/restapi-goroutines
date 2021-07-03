package models

import (
	"context"

	"github.com/edualb/restapi-goroutines/db"
)

const walkDogCollection = "col_walk_dog_service"

type WalkDogServiceDB interface {
	CreateWalkDogService(w WalkDogService) error
}

type WalkDogServiceDBImpl struct {
}

func NewWalkDogServiceDB() WalkDogServiceDB {
	return &WalkDogServiceDBImpl{}
}

func (wdb *WalkDogServiceDBImpl) CreateWalkDogService(w WalkDogService) error {
	client, err := db.GetMongoClient()
	if err != nil {
		return err
	}
	collection := client.Database((db.DB)).Collection(walkDogCollection)
	_, err = collection.InsertOne(context.TODO(), w)
	if err != nil {
		return err
	}
	return nil
}
