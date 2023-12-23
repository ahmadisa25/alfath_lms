package services

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/deps/pagination"
	"alfath_lms/api/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AnnouncementService struct {
	mongo     *mongo.Database
	paginator *pagination.Paginator
}

func (announcementSvc *AnnouncementService) Inject(mongo *mongo.Database, paginator *pagination.Paginator) {
	announcementSvc.mongo = mongo
	announcementSvc.paginator = paginator
}

func (announcementSvc *AnnouncementService) Create(Announcement models.Announcement) (definitions.GenericMongoCreationMessage, error) {
	insertResult, err := announcementSvc.mongo.Collection("announcement").InsertOne(context.TODO(), Announcement)

	if err != nil {
		return definitions.GenericMongoCreationMessage{}, nil
	}
	return definitions.GenericMongoCreationMessage{
		Status:     200,
		InstanceID: insertResult.InsertedID.(primitive.ObjectID),
	}, nil
}

func (materialSvc *MaterialService) Update(id int, material models.ChapterMaterial) (definitions.GenericAPIMessage, error) {
	var materialTemp models.ChapterMaterial
	result := materialSvc.db.Model(&materialTemp).Where("id = ?", id).Updates(&material)
	if result.Error != nil {
		return definitions.GenericAPIMessage{}, result.Error
	}
	return definitions.GenericAPIMessage{
		Status:  200,
		Message: "material is successfully updated",
	}, nil
}

func (announcementSvc *AnnouncementService) Delete(id primitive.ObjectID) (definitions.GenericAPIMessage, error) {
	//delete here means soft delete
	filter := bson.M{"_id": id}
	searchResult := announcementSvc.mongo.Collection("announcement").FindOne(context.TODO(), filter)
	if searchResult.Err() == mongo.ErrNoDocuments {
		return definitions.GenericAPIMessage{
			Status:  400,
			Message: "User not found",
		}, nil
	} else if searchResult.Err() != nil {
		return definitions.GenericAPIMessage{
			Status:  500,
			Message: searchResult.Err().Error(),
		}, nil
	} else {
		res, err := announcementSvc.mongo.Collection("announcement").DeleteOne(context.TODO(), filter)

		if err != nil {
			return definitions.GenericAPIMessage{
				Status:  500,
				Message: err.Error(),
			}, nil
		}

		if res.DeletedCount == 0 {
			return definitions.GenericAPIMessage{
				Status:  500,
				Message: "Failed to delete that announcement. Please try again or contact support",
			}, nil
		}

		return definitions.GenericAPIMessage{
			Status:  200,
			Message: "Announcement is successfully deleted",
		}, nil
	}
}

func (materialSvc *MaterialService) Get(id int) (models.ChapterMaterial, error) {
	var material models.ChapterMaterial

	result := &material
	materialSvc.db.First(result, "id = ?", id)

	return *result, nil

}
