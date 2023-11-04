package services

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/deps/pagination"
	"alfath_lms/api/models"

	"gorm.io/gorm"
)

type AnswerService struct {
	db        *gorm.DB
	paginator *pagination.Paginator
}

func (answerSvc *AnswerService) Inject(db *gorm.DB, paginator *pagination.Paginator) {
	answerSvc.db = db
	answerSvc.paginator = paginator
}

func (answerSvc *AnswerService) Create(answer models.QuizAnswer) (definitions.GenericCreationMessage, error) {
	result := answerSvc.db.Create(&answer)
	if result.Error != nil {
		return definitions.GenericCreationMessage{}, result.Error
	}
	return definitions.GenericCreationMessage{
		Status:     201,
		InstanceID: answer.ID,
	}, nil
}

func (answerSvc *AnswerService) Update(id int, answer models.QuizAnswer) (definitions.GenericAPIMessage, error) {
	var answerTemp models.QuizAnswer
	result := answerSvc.db.Model(&answerTemp).Where("id = ?", id).Updates(&answer)
	if result.Error != nil {
		return definitions.GenericAPIMessage{}, result.Error
	}
	return definitions.GenericAPIMessage{
		Status:  200,
		Message: "answer is successfully updated",
	}, nil
}

func (answerSvc *AnswerService) Delete(id int) (definitions.GenericAPIMessage, error) {
	//delete here means soft delete
	result := answerSvc.db.Where("id = ?", id).Delete(&models.QuizAnswer{})
	if result.Error != nil {
		return definitions.GenericAPIMessage{}, result.Error
	}
	return definitions.GenericAPIMessage{
		Status:  200,
		Message: "answer has been deleted successfully",
	}, nil
}

func (answerSvc *AnswerService) Get(id int) (models.QuizAnswer, error) {
	var answer models.QuizAnswer

	result := &answer
	answerSvc.db.First(result, "id = ?", id)

	return *result, nil

}
