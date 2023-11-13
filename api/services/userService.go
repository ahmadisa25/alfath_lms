package services

import (
	"alfath_lms/api/definitions"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	mongo *mongo.Database
}

func (userSvc *UserService) Inject(mongo *mongo.Database) {
	userSvc.mongo = mongo
}

func (userSvc *UserService) Create(data map[string]interface{}) (definitions.GenericCreationMessage, error) {

	insertResult, err := userSvc.mongo.Collection("users").InsertOne(context.Background(), bson.M(data))

	if err != nil {
		return definitions.GenericCreationMessage{}, nil
	}

	return definitions.GenericCreationMessage{
		Status:     200,
		InstanceID: insertResult.InsertedID.(int),
	}, nil
}
