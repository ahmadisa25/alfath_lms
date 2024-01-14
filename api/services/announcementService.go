package services

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/deps/pagination"
	"alfath_lms/api/funcs"
	"alfath_lms/api/models"
	"context"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AnnouncementService struct {
	mongo     *mongo.Database
	paginator *pagination.Paginator
}

func (announcementSvc *AnnouncementService) Inject(mongo *mongo.Database, paginator *pagination.Paginator) {
	announcementSvc.mongo = mongo
	announcementSvc.paginator = paginator
}

func (announcementSvc *AnnouncementService) GetAll(req definitions.PaginationRequest) (definitions.PaginationResult, error) {
	var mapResult []models.Announcement
	filter := bson.M{}
	if req.Page == "" || req.Page == "0" {
		req.Page = "1"
	}

	if req.PerPage == "" || req.PerPage == "0" {
		req.PerPage = "10"
	}

	limit, err := strconv.Atoi(req.PerPage)

	if err != nil {
		return definitions.PaginationResult{}, err
	}

	page, err := strconv.Atoi(req.Page)

	if err != nil {
		return definitions.PaginationResult{}, err
	}

	if req.Filter != "" {
		//conditions := []bson.M{}
		filters := strings.Split(req.Filter, ",")
		for _, value := range filters {
			filterKey := strings.Split(value, ":")

			filter[filterKey[0]] = bson.M{"$regex": "^" + filterKey[1], "$options": "i"}
		}

		//filter = bson.M{"$or": conditions}
	}

	//fmt.Println(filter)

	findOptions := options.Find()
	findOptions.SetSkip(int64(page*limit - limit))
	findOptions.SetLimit(int64(limit))
	ctx := context.TODO()
	anouncementCursor, err := announcementSvc.mongo.Collection("announcement").Find(ctx, filter, findOptions)

	if err != nil {
		return definitions.PaginationResult{}, err
	}

	defer anouncementCursor.Close(ctx)

	i := 0
	for anouncementCursor.Next(ctx) {
		var res models.Announcement
		anouncementCursor.Decode(&res)
		mapResult = append(mapResult, res)
		i++
	}

	return definitions.PaginationResult{
		Data:    mapResult,
		Page:    page,
		PerPage: limit,
		Total:   int64(i),
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
