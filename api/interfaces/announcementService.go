package interfaces

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/models"

	"go.mongodb.org/mongo-driver/bson"
)

type AnnouncementServiceInterface interface {
	GetAll(req definitions.PaginationRequest) (definitions.PaginationResult, error)
	Get(id string) (definitions.GenericGetMessage[models.Announcement], error)
	Delete(id string) (definitions.GenericAPIMessage, error)
	Create(Announcement models.Announcement) (definitions.GenericMongoCreationMessage, error)
	Update(id string, Updates []bson.E) (definitions.GenericAPIMessage, error)
}
