package interfaces

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/models"
)

type MaterialServiceInterface interface {
	Get(id int) (models.ChapterMaterial, error)
	Delete(id int) (definitions.GenericAPIMessage, error)
	Create(chapter models.ChapterMaterial) (definitions.GenericCreationMessage, error)
	Update(id int, chapter models.ChapterMaterial) (definitions.GenericAPIMessage, error)
}
