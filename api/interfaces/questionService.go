package interfaces

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/models"
)

type QuestionServiceInterface interface {
	Get(id int) (models.QuizQuestion, error)
	Delete(id int) (definitions.GenericAPIMessage, error)
	Create(chapter models.QuizQuestion) (definitions.GenericCreationMessage, error)
	Update(id int, chapter models.QuizQuestion) (definitions.GenericAPIMessage, error)
}
