package services

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/deps/pagination"
	"alfath_lms/api/models"

	"gorm.io/gorm"
)

type StudentQuizService struct {
	db        *gorm.DB
	paginator *pagination.Paginator
}

func (studQuizSvc *StudentQuizService) Inject(db *gorm.DB, paginator *pagination.Paginator) {
	studQuizSvc.db = db
	studQuizSvc.paginator = paginator
}

func (studQuizSvc *StudentQuizService) Create(stdQuiz models.StudentQuiz) (definitions.GenericCreationMessage, error) {
	result := studQuizSvc.db.Create(&stdQuiz)
	if result.Error != nil {
		return definitions.GenericCreationMessage{}, result.Error
	}
	return definitions.GenericCreationMessage{
		Status:     201,
		InstanceID: stdQuiz.ID,
	}, nil
}
