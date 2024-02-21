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
		return funcs.CorsedDataResponse(questionController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select a question!",
		}))
	}

	intID, err := strconv.Atoi(req.Params["id"])
	//PrintError(err)

	if intID <= 0 {
		return funcs.CorsedDataResponse(questionController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select a question!",
		}))
	}

	question, err := questionController.questionService.Get(intID)
	if err != nil {
		return funcs.CorsedDataResponse(questionController.responder.Data(definitions.GenericAPIMessage{
			Status:  500,
			Message: "We cannot process your request. Please try again or contact support!",
		}))
	}

	if question.ID <= 0 {
		return funcs.CorsedDataResponse(questionController.responder.Data(definitions.GenericAPIMessage{
			Status:  404,
			Message: "question Not Found!",
		}))
	}

	return funcs.CorsedDataResponse(questionController.responder.Data(GetQuestionResponse{
		Status: 200,
		Data:   question,
	}))
}

func (questionController *QuestionController) Update(ctx context.Context, req *web.Request) web.Result {
	if req.Params["id"] == "" {
		return funcs.CorsedDataResponse(questionController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select a question!",
		}))
	}

	intID, err := strconv.Atoi(req.Params["id"])
	if err != nil {
		return funcs.CorsedResponse(questionController.responder.HTTP(500, strings.NewReader(err.Error())))
	}
	//PrintError(err)

	if intID <= 0 {
		return funcs.CorsedDataResponse(questionController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select a question!",
		}))
	}

	question, err := questionController.questionService.Get(intID)
	if err != nil {
		return funcs.CorsedDataResponse(questionController.responder.Data(definitions.GenericAPIMessage{
			Status:  500,
			Message: "We cannot process your request. Please try again or contact support!",
		}))
	}

	if question.ID <= 0 {
		return funcs.CorsedDataResponse(questionController.responder.Data(definitions.GenericAPIMessage{
			Status:  404,
			Message: "question Not Found!",
		}))
	}

	formError := req.Request().ParseForm()
	if formError != nil {
		return funcs.CorsedResponse(questionController.responder.HTTP(400, strings.NewReader(formError.Error())))
	}

	form := req.Request().Form

	questionData := &models.QuizQuestion{
		Title:         funcs.ValidateStringFormKeys("Title", form, "string").(string),
		ChapterQuizID: funcs.ValidateStringFormKeys("ChapterQuizID", form, "int").(int),
		Type:          funcs.ValidateStringFormKeys("Type", form, "string").(string),
		Length:        funcs.ValidateStringFormKeys("Length", form, "int").(int),
		Choices:       funcs.ValidateStringFormKeys("Choices", form, "string").(string),
		CreatedAt:     time.Now(),
	}

	if questionData.Type != "" {
		isTypeExist := false
		for _, val := range questionController.questionTypes {
			if question.Type == val {
				isTypeExist = true
			}
		}

		if !isTypeExist {
			return funcs.CorsedDataResponse(questionController.responder.Data(definitions.GenericAPIMessage{
				Status:  400,
				Message: "Question type doesn't exist!",
			}))
		}
	}

	//fmt.Printf("validator: %+v\n", quizController.validator.validate)
	validateError := questionController.customValidator.Validate.Struct(questionData)
	if validateError != nil {
		errorResponse := funcs.ErrorPackagingForMaps(questionController.customValidator.TranslateError(validateError))
		errorResponse, packError := funcs.ErrorPackaging(errorResponse, 400)
		if packError != nil {
			return funcs.CorsedResponse(questionController.responder.HTTP(500, strings.NewReader(packError.Error())))
		}
		return funcs.CorsedResponse(questionController.responder.HTTP(400, strings.NewReader(errorResponse)))
	}

	result, err := questionController.questionService.Update(intID, *questionData)
	if err != nil {
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return funcs.CorsedResponse(questionController.responder.HTTP(500, strings.NewReader(packError.Error())))
		}
		return funcs.CorsedResponse(questionController.responder.HTTP(500, strings.NewReader(errorResponse)))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return funcs.CorsedResponse(questionController.responder.HTTP(400, strings.NewReader(resErr.Error())))
	}

	return funcs.CorsedResponse(questionController.responder.HTTP(uint(result.Status), strings.NewReader(string(res))))
}

func (questionController *QuestionController) Delete(ctx context.Context, req *web.Request) web.Result {
	if req.Params["id"] == "" {
		return funcs.CorsedDataResponse(questionController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select a question!",
		}))
	}

	intID, err := strconv.Atoi(req.Params["id"])
	//PrintError(err)

	if intID <= 0 {
		return funcs.CorsedDataResponse(questionController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select a question!",
		}))
	}

	question, err := questionController.questionService.Get(intID)
	if err != nil {
		return funcs.CorsedDataResponse(questionController.responder.Data(definitions.GenericAPIMessage{
			Status:  500,
			Message: "We cannot process your request. Please try again or contact support!",
		}))
	}

	if question.ID <= 0 {
		return funcs.CorsedDataResponse(questionController.responder.Data(definitions.GenericAPIMessage{
			Status:  404,
			Message: "question Not Found!",
		}))
	}

	result, err := questionController.questionService.Delete(intID)
	if err != nil {
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return funcs.CorsedResponse(questionController.responder.HTTP(500, strings.NewReader(packError.Error())))
		}
		return funcs.CorsedResponse(questionController.responder.HTTP(500, strings.NewReader(errorResponse)))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return funcs.CorsedResponse(questionController.responder.HTTP(400, strings.NewReader(resErr.Error())))
	}

	return funcs.CorsedResponse(questionController.responder.HTTP(uint(result.Status), strings.NewReader(string(res))))
}

func (questionController *QuestionController) Create(ctx context.Context, req *web.Request) web.Result {
	formError := req.Request().ParseForm()
	if formError != nil {
		return funcs.CorsedResponse(questionController.responder.HTTP(400, strings.NewReader(formError.Error())))
	}

	form := req.Request().Form

	question := &models.QuizQuestion{
		Title:         funcs.ValidateStringFormKeys("Title", form, "string").(string),
		ChapterQuizID: funcs.ValidateStringFormKeys("ChapterQuizID", form, "int").(int),
		Type:          funcs.ValidateStringFormKeys("Type", form, "string").(string),
		Length:        funcs.ValidateStringFormKeys("Length", form, "int").(int),
		Choices:       funcs.ValidateStringFormKeys("Choices", form, "string").(string),
		CreatedAt:     time.Now(),
	}

	if question.Type != "" {
		isTypeExist := false
		for _, val := range questionController.questionTypes {
			if question.Type == val {
				isTypeExist = true
			}
		}
		if !isTypeExist {
			return funcs.CorsedDataResponse(questionController.responder.Data(definitions.GenericAPIMessage{
				Status:  400,
				Message: "Question type doesn't exist!",
			}))
		}
	}

	//fmt.Printf("validator: %+v\n", instructorController.validator.validate)
	validateError := questionController.customValidator.Validate.Struct(question)
	if validateError != nil {
		errorResponse := funcs.ErrorPackagingForMaps(questionController.customValidator.TranslateError(validateError))
		errorResponse, packError := funcs.ErrorPackaging(errorResponse, 400)
		if packError != nil {
			return funcs.CorsedResponse(questionController.responder.HTTP(500, strings.NewReader(packError.Error())))
		}
		return funcs.CorsedResponse(questionController.responder.HTTP(400, strings.NewReader(errorResponse)))
	}

	result, err := questionController.questionService.Create(*question)
	if err != nil {
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return funcs.CorsedResponse(questionController.responder.HTTP(500, strings.NewReader(packError.Error())))
		}
		return funcs.CorsedResponse(questionController.responder.HTTP(500, strings.NewReader(errorResponse)))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return funcs.CorsedResponse(questionController.responder.HTTP(400, strings.NewReader(resErr.Error())))
	}

	return funcs.CorsedResponse(questionController.responder.HTTP(uint(result.Status), strings.NewReader(string(res))))

}
