package interfaces

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/models"
)

type QuizServiceInterface interface {
	Get(id int) (models.ChapterQuiz, error)
	Delete(id int) (definitions.GenericAPIMessage, error)
	Create(chapter models.ChapterQuiz) (definitions.GenericCreationMessage, error)
	Update(id int, chapter models.ChapterQuiz) (definitions.GenericAPIMessage, error)
}
