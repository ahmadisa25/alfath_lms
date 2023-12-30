package services

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/deps/pagination"
	"alfath_lms/api/funcs"
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

func (announcementSvc *AnnouncementService) GetAll(limit int, page int) (definitions.PaginationResult, error) {
	return definitions.PaginationResult{
		Data:    announcementSvc.mongo.Collection("announcement").find().skip(page*limit - limit).limit(limit),
		Page:    page,
		PerPage: limit,
		Total:   0,
		Status:  200,
	}, nil
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

func (announcementSvc *AnnouncementService) Get(id string) (definitions.GenericGetMessage[models.Announcement], error) {
	objID := funcs.StringToMongoOID(id)
	if objID == primitive.NilObjectID {
		return definitions.GenericGetMessage[models.Announcement]{
			Status: 500,
			Data:   models.Announcement{},
		}, nil
	}
	filter := bson.M{"_id": objID}
	searchResult := announcementSvc.mongo.Collection("announcement").FindOne(context.TODO(), filter)
	if searchResult.Err() == mongo.ErrNoDocuments {
		return definitions.GenericGetMessage[models.Announcement]{
			Status: 500,
			Data:   models.Announcement{},
		}, nil
	} else if searchResult.Err() != nil {
		return definitions.GenericGetMessage[models.Announcement]{
			Status: 500,
			Data:   models.Announcement{},
		}, nil
	} else {
		var existingAnnouncement models.Announcement
		searchResult.Decode(&existingAnnouncement)

		return definitions.GenericGetMessage[models.Announcement]{
			Status: 200,
			Data:   existingAnnouncement,
		}, nil
	}
}

func (announcementSvc *AnnouncementService) Update(id string, Updates []bson.E) (definitions.GenericAPIMessage, error) {
	objID := funcs.StringToMongoOID(id)
	if objID == primitive.NilObjectID {
		return definitions.GenericAPIMessage{
			Status:  400,
			Message: "Announcement not found",
		}, nil
	}
	filter := bson.M{"_id": objID}
	searchResult := announcementSvc.mongo.Collection("announcement").FindOne(context.TODO(), filter)
	if searchResult.Err() == mongo.ErrNoDocuments {
		return definitions.GenericAPIMessage{
			Status:  400,
			Message: "Announcement not found",
		}, nil
	} else if searchResult.Err() != nil {
		return definitions.GenericAPIMessage{
			Status:  500,
			Message: searchResult.Err().Error(),
		}, nil
	} else {
		_, err := announcementSvc.mongo.Collection("announcement").UpdateOne(context.TODO(), filter, bson.D{{"$set", Updates}})

		if err != nil {
			return definitions.GenericAPIMessage{
				Status:  500,
				Message: err.Error(),
			}, nil
		}

		return definitions.GenericAPIMessage{
			Status:  200,
			Message: "Announcement is successfully updated",
		}, nil
	}
}

func (announcementSvc *AnnouncementService) Delete(id string) (definitions.GenericAPIMessage, error) {
	//delete here means soft delete
	objID := funcs.StringToMongoOID(id)
	if objID == primitive.NilObjectID {
		return definitions.GenericAPIMessage{
			Status:  400,
			Message: "Announcement not found",
		}, nil
	}
	filter := bson.M{"_id": objID}

	searchResult := announcementSvc.mongo.Collection("announcement").FindOne(context.TODO(), filter)
	if searchResult.Err() == mongo.ErrNoDocuments {
		return definitions.GenericAPIMessage{
			Status:  400,
			Message: "Announcement not found",
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

/*func (materialSvc *MaterialService) Get(id int) (models.ChapterMaterial, error) {
	var material models.ChapterMaterial

	result := &material
	materialSvc.db.First(result, "id = ?", id)

	return *result, nil

}*/
