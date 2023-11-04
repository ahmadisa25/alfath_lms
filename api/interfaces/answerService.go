package interfaces

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/models"
)

type AnswerServiceInterface interface {
	Get(id int) (models.QuizAnswer, error)
	Delete(id int) (definitions.GenericAPIMessage, error)
	Create(chapter models.QuizAnswer) (definitions.GenericCreationMessage, error)
	Update(id int, chapter models.QuizAnswer) (definitions.GenericAPIMessage, error)
}
