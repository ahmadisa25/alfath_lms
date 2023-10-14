package interfaces

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/models"
)

type CourseServiceInterface interface {
	GetAll(req definitions.PaginationRequest) (definitions.PaginationResult, error)
	Get(id int) (models.Course, error)
	Delete(id int) (definitions.GenericAPIMessage, error)
	Create(Student models.Course) (definitions.GenericCreationMessage, error)
	Update(id int, Course models.Course, existingCourse models.Course) (definitions.GenericAPIMessage, error)
}
