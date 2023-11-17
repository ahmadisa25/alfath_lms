package services

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/models"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	mongo *mongo.Database
}

func (userSvc *UserService) Inject(mongo *mongo.Database) {
	userSvc.mongo = mongo
}

func (userSvc *UserService) Create(User models.User) (definitions.GenericMongoCreationMessage, error) {
	//return definitions.GenericCreationMessage{}, nil
	insertResult, err := userSvc.mongo.Collection("users").InsertOne(context.Background(), User)

	if err != nil {
		return definitions.GenericMongoCreationMessage{}, nil
	}

	return definitions.GenericMongoCreationMessage{
		Status:     200,
		InstanceID: insertResult.InsertedID.(primitive.ObjectID),
	}, nil
}
