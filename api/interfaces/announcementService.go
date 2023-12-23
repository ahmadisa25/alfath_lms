package interfaces

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AnnouncementServiceInterface interface {
	GetAll(req definitions.PaginationRequest) (definitions.PaginationResult, error)
	Get(id int) (models.Announcement, error)
	Delete(id primitive.ObjectID) (definitions.GenericAPIMessage, error)
	Create(Announcement models.Announcement) (definitions.GenericMongoCreationMessage, error)
	Update(id int, Student models.Announcement, existingAnnouncement models.Announcement) (definitions.GenericAPIMessage, error)
}
