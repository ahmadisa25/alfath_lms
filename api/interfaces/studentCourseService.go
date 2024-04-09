package interfaces

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/models"
)

type StudentCourseServiceInterface interface {
	Create(stdQuiz models.StudentCourse) (definitions.GenericCreationMessage, error)
	Delete(studentId int) (definitions.GenericAPIMessage, error)
}
