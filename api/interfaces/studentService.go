package interfaces

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/models"
)

type StudentServiceInterface interface {
	GetAllStudents(req definitions.PaginationRequest) (definitions.PaginationResult, error)
	GetStudent(id int) (models.Student, error)
	DeleteStudent(id int) (definitions.GenericAPIMessage, error)
	CreateStudent(Student models.Student) (definitions.GenericCreationMessage, error)
	UpdateStudent(id int, Student models.Student) (definitions.GenericAPIMessage, error)
}
