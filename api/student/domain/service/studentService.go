package service

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/student/domain/entity"
)

type StudentServiceInterface interface {
	GetAllStudents(req definitions.PaginationRequest) (definitions.PaginationResult, error)
	GetStudent(id int) (entity.Student, error)
	DeleteStudent(id int) (definitions.GenericAPIMessage, error)
	CreateStudent(Student entity.Student) (definitions.GenericCreationMessage, error)
	UpdateStudent(id int, Student entity.Student) (definitions.GenericAPIMessage, error)
}
