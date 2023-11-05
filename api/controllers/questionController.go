package controllers

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/deps/validator"
	"alfath_lms/api/funcs"
	"alfath_lms/api/interfaces"
	"alfath_lms/api/models"
	"context"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"flamingo.me/flamingo/v3/framework/web"
)

type (
	QuestionController struct {
		responder       *web.Responder
		customValidator *validator.CustomValidator
		questionService interfaces.QuestionServiceInterface
		questionTypes   []string
	}

	GetQuestionResponse struct {
		Status int
		Data   models.QuizQuestion
	}
)

func (questionController *QuestionController) Inject(
	responder *web.Responder,
	customValidator *validator.CustomValidator,
	questionService interfaces.QuestionServiceInterface,
) {
	questionController.responder = responder
	questionController.customValidator = customValidator
	questionController.questionService = questionService
	questionController.questionTypes = []string{"single-text", "multiple-text", "single-choices", "multiple-choices"}
}

func (questionController *QuestionController) Get(ctx context.Context, req *web.Request) web.Result {
	if req.Params["id"] == "" {
		return questionController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select a question!",
		})
	}

	intID, err := strconv.Atoi(req.Params["id"])
	//PrintError(err)

	if intID <= 0 {
		return questionController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select a question!",
		})
	}

	question, err := questionController.questionService.Get(intID)
	if err != nil {
		return questionController.responder.Data(definitions.GenericAPIMessage{
			Status:  500,
			Message: "We cannot process your request. Please try again or contact support!",
		})
	}

	if question.ID <= 0 {
		return questionController.responder.Data(definitions.GenericAPIMessage{
			Status:  404,
			Message: "question Not Found!",
		})
	}

	return questionController.responder.Data(GetQuestionResponse{
		Status: 200,
		Data:   question,
	})
}

func (questionController *QuestionController) Update(ctx context.Context, req *web.Request) web.Result {
	if req.Params["id"] == "" {
		return questionController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select a question!",
		})
	}

	intID, err := strconv.Atoi(req.Params["id"])
	if err != nil {
		return questionController.responder.HTTP(500, strings.NewReader(err.Error()))
	}
	//PrintError(err)

	if intID <= 0 {
		return questionController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select a question!",
		})
	}

	question, err := questionController.questionService.Get(intID)
	if err != nil {
		return questionController.responder.Data(definitions.GenericAPIMessage{
			Status:  500,
			Message: "We cannot process your request. Please try again or contact support!",
		})
	}

	if question.ID <= 0 {
		return questionController.responder.Data(definitions.GenericAPIMessage{
			Status:  404,
			Message: "question Not Found!",
		})
	}

	formError := req.Request().ParseForm()
	if formError != nil {
		return questionController.responder.HTTP(400, strings.NewReader(formError.Error()))
	}

	form := req.Request().Form

	questionData := &models.QuizQuestion{
		Answer:         funcs.ValidateStringFormKeys("Answer", form, "string").(string),
		QuizQuestionID: funcs.ValidateStringFormKeys("QuizQuestion", form, "int").(int),
		StudentID:      funcs.ValidateStringFormKeys("Student", form, "int").(int),
		CreatedAt:      time.Now(),
	}

	//fmt.Printf("validator: %+v\n", quizController.validator.validate)
	validateError := answerController.customValidator.Validate.Struct(answerData)
	if validateError != nil {
		errorResponse := funcs.ErrorPackagingForMaps(answerController.customValidator.TranslateError(validateError))
		errorResponse, packError := funcs.ErrorPackaging(errorResponse, 400)
		if packError != nil {
			return answerController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return answerController.responder.HTTP(400, strings.NewReader(errorResponse))
	}

	result, err := answerController.answerService.Update(intID, *answerData)
	if err != nil {
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return answerController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return answerController.responder.HTTP(500, strings.NewReader(errorResponse))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return answerController.responder.HTTP(400, strings.NewReader(resErr.Error()))
	}

	return answerController.responder.HTTP(uint(result.Status), strings.NewReader(string(res)))
}

func (questionController *QuestionController) Delete(ctx context.Context, req *web.Request) web.Result {
	if req.Params["id"] == "" {
		return questionController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select a question!",
		})
	}

	intID, err := strconv.Atoi(req.Params["id"])
	//PrintError(err)

	if intID <= 0 {
		return questionController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select a question!",
		})
	}

	question, err := questionController.questionService.Get(intID)
	if err != nil {
		return questionController.responder.Data(definitions.GenericAPIMessage{
			Status:  500,
			Message: "We cannot process your request. Please try again or contact support!",
		})
	}

	if question.ID <= 0 {
		return questionController.responder.Data(definitions.GenericAPIMessage{
			Status:  404,
			Message: "question Not Found!",
		})
	}

	result, err := questionController.questionService.Delete(intID)
	if err != nil {
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return questionController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return questionController.responder.HTTP(500, strings.NewReader(errorResponse))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return questionController.responder.HTTP(400, strings.NewReader(resErr.Error()))
	}

	return questionController.responder.HTTP(uint(result.Status), strings.NewReader(string(res)))
}

func (questionController *QuestionController) Create(ctx context.Context, req *web.Request) web.Result {
	formError := req.Request().ParseForm()
	if formError != nil {
		return questionController.responder.HTTP(400, strings.NewReader(formError.Error()))
	}

	form := req.Request().Form

	question := &models.QuizQuestion{
		Title:         funcs.ValidateStringFormKeys("Title", form, "string").(string),
		ChapterQuizID: funcs.ValidateStringFormKeys("ChapterQuizID", form, "int").(int),
		Type:          funcs.ValidateStringFormKeys("Type", form, "string").(string),
		Length:        funcs.ValidateStringFormKeys("Length", form, "int").(int),
		CreatedAt:     time.Now(),
	}

	if question.Type != "" {

	}

	//fmt.Printf("validator: %+v\n", instructorController.validator.validate)
	validateError := questionController.customValidator.Validate.Struct(question)
	if validateError != nil {
		errorResponse := funcs.ErrorPackagingForMaps(questionController.customValidator.TranslateError(validateError))
		errorResponse, packError := funcs.ErrorPackaging(errorResponse, 400)
		if packError != nil {
			return questionController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return questionController.responder.HTTP(400, strings.NewReader(errorResponse))
	}

	result, err := questionController.questionService.Create(*question)
	if err != nil {
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return questionController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return questionController.responder.HTTP(500, strings.NewReader(errorResponse))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return questionController.responder.HTTP(400, strings.NewReader(resErr.Error()))
	}

	return questionController.responder.HTTP(uint(result.Status), strings.NewReader(string(res)))

}