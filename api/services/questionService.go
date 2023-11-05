package services

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/deps/pagination"
	"alfath_lms/api/models"

	"gorm.io/gorm"
)

type QuestionService struct {
	db        *gorm.DB
	paginator *pagination.Paginator
}

func (questionSvc *QuestionService) Inject(db *gorm.DB, paginator *pagination.Paginator) {
	questionSvc.db = db
	questionSvc.paginator = paginator
}

func (questionSvc *QuestionService) Create(question models.QuizQuestion) (definitions.GenericCreationMessage, error) {
	result := questionSvc.db.Create(&question)
	if result.Error != nil {
		return definitions.GenericCreationMessage{}, result.Error
	}
	return definitions.GenericCreationMessage{
		Status:     201,
		InstanceID: question.ID,
	}, nil
}

func (questionSvc *QuestionService) Update(id int, question models.QuizQuestion) (definitions.GenericAPIMessage, error) {
	var questionTemp models.QuizQuestion
	result := questionSvc.db.Model(&questionTemp).Where("id = ?", id).Updates(&question)
	if result.Error != nil {
		return definitions.GenericAPIMessage{}, result.Error
	}
	return definitions.GenericAPIMessage{
		Status:  200,
		Message: "question is successfully updated",
	}, nil
}

func (questionSvc *QuestionService) Delete(id int) (definitions.GenericAPIMessage, error) {
	//delete here means soft delete
	result := questionSvc.db.Where("id = ?", id).Delete(&models.QuizQuestion{})
	if result.Error != nil {
		return definitions.GenericAPIMessage{}, result.Error
	}
	return definitions.GenericAPIMessage{
		Status:  200,
		Message: "question has been deleted successfully",
	}, nil
}

func (questionSvc *QuestionService) Get(id int) (models.QuizQuestion, error) {
	var question models.QuizQuestion

	result := &question
	questionSvc.db.First(result, "id = ?", id)

	return *result, nil

}
