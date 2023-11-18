package services

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/models"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	mongo *mongo.Database
}

func (userSvc *UserService) Inject(mongo *mongo.Database) {
	userSvc.mongo = mongo
}

func (userSvc *UserService) Create(User models.User, Role string) (definitions.GenericMongoCreationMessage, error) {
	filter := bson.D{{"email", User.Email}}
	searchResult := userSvc.mongo.Collection("users").FindOne(context.TODO(), filter)
	if searchResult.Err() == mongo.ErrNoDocuments {
		filter = bson.D{{"name", Role}}
		roleSearch := userSvc.mongo.Collection("roles").FindOne(context.TODO(), filter)

		if roleSearch.Err() == mongo.ErrNoDocuments {
			return definitions.GenericMongoCreationMessage{}, errors.New("role doesn't exist")
		} else if roleSearch.Err() != nil {
			return definitions.GenericMongoCreationMessage{}, roleSearch.Err()
		}
		var role models.Role
		roleSearch.Decode(&role)
		User.Role = role

		insertResult, err := userSvc.mongo.Collection("users").InsertOne(context.TODO(), User)

		if err != nil {
			return definitions.GenericMongoCreationMessage{}, nil
		}

		return definitions.GenericMongoCreationMessage{
			Status:     200,
			InstanceID: insertResult.InsertedID.(primitive.ObjectID),
		}, nil
	} else if searchResult.Err() != nil {
		return definitions.GenericMongoCreationMessage{}, searchResult.Err()
	} else {
		return definitions.GenericMongoCreationMessage{}, errors.New("user with that email already exists")
	}

}
