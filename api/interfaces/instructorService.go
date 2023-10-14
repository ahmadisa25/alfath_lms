package interfaces

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/models"
)

type InstructorServiceInterface interface {
	GetAllInstructors(req definitions.PaginationRequest) (definitions.PaginationResult, error)
	GetInstructor(id int) (models.Instructor, error)
	DeleteInstructor(id int) (definitions.GenericAPIMessage, error)
	CreateInstructor(instructor models.Instructor) (definitions.GenericCreationMessage, error)
	UpdateInstructor(id int, instructor models.Instructor, existingInstructor models.Instructor) (definitions.GenericAPIMessage, error)
}
