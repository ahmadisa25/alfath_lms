package interfaces

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/models"
)

type AnnouncementServiceInterface interface {
	GetAll(req definitions.PaginationRequest) (definitions.PaginationResult, error)
	Get(id int) (models.Announcement, error)
	Delete(id int) (definitions.GenericAPIMessage, error)
	Create(Student models.Announcement) (definitions.GenericCreationMessage, error)
	Update(id int, Student models.Announcement, existingAnnouncement models.Announcement) (definitions.GenericAPIMessage, error)
}
