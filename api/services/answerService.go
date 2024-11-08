package services

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/deps/pagination"
	"alfath_lms/api/funcs"
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

func (answerSvc *AnswerService) GetAll(req definitions.PaginationRequest) (definitions.PaginationResult, error) {
	paginationParams := definitions.PaginationParam{
		Sql:          "with answers as (select ma.*, ms.name, ms.email, ms.mobile_phone, mc.id as quiz_id, mq.title from ms_quiz_questions mq inner join ms_quiz_answers ma on ma.quiz_question_id = mq.id inner join ms_chapter_quiz mc on mq.chapter_quiz_id = mc.id inner join ms_student ms on ma.student_id = ms.id) select -select- from answers -where- order by answers.student_id",
		SelectFields: []string{"answer", "name", "email", "mobile_phone"},
		SearchFields: map[string]string{},
		FilterFields: map[string]string{
			"name":             "answers.name",
			"email":            "answers.email",
			"mobile_phone":     "answers.mobile_phone",
			"quiz_id":          "answers.quiz_id",
			"quiz_question_id": "answers.quiz_question_id",
			"student_id":       "answers.student_id",
		},
		NullFilterFields: map[string]bool{
			"deleted_at": true,
		},
	}

	res := answerSvc.paginator.Paginate(req, paginationParams)
	return res, nil
}

func (answerSvc *AnswerService) GetAllDistinct(req definitions.PaginationRequest) (definitions.PaginationResult, error) {
	paginationParams := definitions.PaginationParam{
		Sql:          "select distinct name, email, ms.id as student_id, msq.final_grade from ms_quiz_answers ma inner join ms_quiz_questions mq on ma.quiz_question_id = mq.id inner join ms_student ms on ma.student_id = ms.id left join ms_student_quiz msq on ms.id = msq.student_id -where-",
		SelectFields: []string{"answer", "name", "email", "mobile_phone"},
		SearchFields: map[string]string{},
		FilterFields: map[string]string{
			"chapter_quiz_id": "chapter_quiz_id",
		},
		NullFilterFields: map[string]bool{
			"deleted_at": true,
		},
	}

	res := answerSvc.paginator.Paginate(req, paginationParams)
	return res, nil
}

func (answerSvc *AnswerService) Create(answer definitions.SubmittedQuizAnswer) definitions.GenericAPIMessage {
	/*delResult := answerSvc.db.Where("id = ?", id).Delete(&models.QuizAnswer{})
	if result.Error != nil {
		return definitions.GenericAPIMessage{}, result.Error
	}*/
	count := 0
	studentID := answer.StudentID
	var answerTemp models.QuizAnswer
	var questionIDs []int
	for k, _ := range funcs.JsonToMap(answer.Answer) {
		questionID := funcs.StringToPositiveInt(k)
		questionIDs = append(questionIDs, questionID)
	}
	err := answerSvc.db.Where("quiz_question_id IN (?) and student_id = ?", questionIDs, studentID).Delete({}models.QuizAnswer).Error
	if err != nil {
		return definitions.GenericAPIMessage{
			Status:  500,
			Message: "Failed to update answers!",
		}
	}
	for k, value := range funcs.JsonToMap(answer.Answer) {
		answerTemp = models.QuizAnswer{}
		answerTemp.Answer = value.(string)
		answerTemp.StudentID = answer.StudentID

		questionID := funcs.StringToPositiveInt(k)

		if questionID <= 0 {
			return definitions.GenericAPIMessage{
				Status:  400,
				Message: "Question not found!",
			}
		}

		answerTemp.QuizQuestionID = questionID
		answerTemp.CreatedAt = answer.CreatedAt
		result := answerSvc.db.Create(&answerTemp)
		if result.Error == nil {
			count++
		}
	}

	return definitions.GenericAPIMessage{
		Status:  200,
		Message: "Successfully inserted " + string(count) + " answers",
	}
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
