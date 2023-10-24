package interfaces

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/models"
)

type CourseServiceInterface interface {
	GetAll(req definitions.PaginationRequest) (definitions.PaginationResult, error)
	Get(id int) (models.Course, error)
	Delete(id int) (definitions.GenericAPIMessage, error)
	Create(Course models.Course, instructorList string) (definitions.GenericCreationMessage, error)
	Update(Course models.Course, instructorList string) (definitions.GenericAPIMessage, error)
}
