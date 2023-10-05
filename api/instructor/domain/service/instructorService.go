package service

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/instructor/domain/entity"
)

type InstructorServiceInterface interface {
	GetAllInstructors(req definitions.PaginationRequest) (definitions.PaginationResult, error)
	GetInstructor(id int) (entity.Instructor, error)
	DeleteInstructor(id int) (definitions.GenericAPIMessage, error)
	CreateInstructor(instructor entity.Instructor) (definitions.GenericCreationMessage, error)
	UpdateInstructor(id int, instructor entity.Instructor) (definitions.GenericAPIMessage, error)
}
