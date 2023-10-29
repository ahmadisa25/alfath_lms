package interfaces

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/models"
)

type ChapterServiceInterface interface {
	Get(id int) (models.CourseChapter, error)
	Delete(id int) (definitions.GenericAPIMessage, error)
	Create(chapter models.CourseChapter) (definitions.GenericCreationMessage, error)
	Update(id int, chapter models.CourseChapter) (definitions.GenericAPIMessage, error)
}
