package services

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/deps/pagination"
	"alfath_lms/api/models"

	"gorm.io/gorm"
)

type QuizService struct {
	db        *gorm.DB
	paginator *pagination.Paginator
}

func (quizSvc *QuizService) Inject(db *gorm.DB, paginator *pagination.Paginator) {
	quizSvc.db = db
	quizSvc.paginator = paginator
}

func (quizSvc *QuizService) Create(quiz models.ChapterQuiz) (definitions.GenericCreationMessage, error) {
	result := quizSvc.db.Create(&quiz)
	if result.Error != nil {
		return definitions.GenericCreationMessage{}, result.Error
	}
	return definitions.GenericCreationMessage{
		Status:     201,
		InstanceID: quiz.ID,
	}, nil
}

func (quizSvc *QuizService) Update(id int, quiz models.ChapterQuiz) (definitions.GenericAPIMessage, error) {
	var quizTemp models.ChapterQuiz
	result := quizSvc.db.Model(&quizTemp).Where("id = ?", id).Updates(&quiz)
	if result.Error != nil {
		return definitions.GenericAPIMessage{}, result.Error
	}
	return definitions.GenericAPIMessage{
		Status:  200,
		Message: "quiz is successfully updated",
	}, nil
}

func (quizSvc *QuizService) Delete(id int) (definitions.GenericAPIMessage, error) {
	//delete here means soft delete
	result := quizSvc.db.Where("id = ?", id).Delete(&models.ChapterQuiz{})
	if result.Error != nil {
		return definitions.GenericAPIMessage{}, result.Error
	}
	return definitions.GenericAPIMessage{
		Status:  200,
		Message: "quiz has been deleted successfully",
	}, nil
}

func (quizSvc *QuizService) Get(id int) (models.ChapterQuiz, error) {
	var quiz models.ChapterQuiz

	result := &quiz
	quizSvc.db.Preload("Questions").First(result, "id = ?", id)

	return *result, nil

}
