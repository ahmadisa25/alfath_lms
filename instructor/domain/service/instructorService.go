package service

import (
	"alfath_lms/instructor/domain/entity"
	"alfath_lms/api/definitions"
)

type InstructorServiceInterface interface {
	GetInstructor(id int) (entity.Instructor, error)
	CreateInstructor(instructor entity.Instructor) (definitions.GenericCreationMessage, error)
	UpdateInstructor(id int, instructor entity.Instructor) (definitions.GenericCreationMessage, error)
}
