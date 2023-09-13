package service

import (
	"alfath_lms/instructor/domain/entity"
)

type InstructorServiceInterface interface {
	GetInstructor(id int) (entity.Instructor, error)
}
