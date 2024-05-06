package interfaces

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/models"
)

type AnswerServiceInterface interface {
	GetAll(req definitions.PaginationRequest) (definitions.PaginationResult, error)
	GetAllDistinct(req definitions.PaginationRequest) (definitions.PaginationResult, error)
	Get(id int) (models.QuizAnswer, error)
	Delete(id int) (definitions.GenericAPIMessage, error)
	Create(answer models.QuizAnswer) (definitions.GenericAPIMessage, error)
	Update(id int, chapter models.QuizAnswer) (definitions.GenericAPIMessage, error)
}
