package interfaces

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/models"
)

type StudentQuizServiceInterface interface {
	Create(stdQuiz models.StudentQuiz) (definitions.GenericCreationMessage, error)
}
