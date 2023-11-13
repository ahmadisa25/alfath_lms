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
	QuizController struct {
		responder       *web.Responder
		customValidator *validator.CustomValidator
		quizService     interfaces.QuizServiceInterface
	}

	GetQuizResponse struct {
		Status int
		Data   models.ChapterQuiz
	}
)

func (quizController *QuizController) Inject(
	responder *web.Responder,
	customValidator *validator.CustomValidator,
	quizService interfaces.QuizServiceInterface,
) {
	quizController.responder = responder
	quizController.customValidator = customValidator
	quizController.quizService = quizService
}

func (quizController *QuizController) Get(ctx context.Context, req *web.Request) web.Result {
	if req.Params["id"] == "" {
		return quizController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an quiz!",
		})
	}

	intID, err := strconv.Atoi(req.Params["id"])
	//PrintError(err)

	if intID <= 0 {
		return quizController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an quiz!",
		})
	}

	quiz, err := quizController.quizService.Get(intID)
	if err != nil {
		return quizController.responder.Data(definitions.GenericAPIMessage{
			Status:  500,
			Message: "We cannot process your request. Please try again or contact support!",
		})
	}

	if quiz.ID <= 0 {
		return quizController.responder.Data(definitions.GenericAPIMessage{
			Status:  404,
			Message: "quiz Not Found!",
		})
	}

	return quizController.responder.Data(GetQuizResponse{
		Status: 200,
		Data:   quiz,
	})
}

func (quizController *QuizController) Update(ctx context.Context, req *web.Request) web.Result {
	if req.Params["id"] == "" {
		return quizController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an quiz!",
		})
	}

	intID, err := strconv.Atoi(req.Params["id"])
	if err != nil {
		return quizController.responder.HTTP(500, strings.NewReader(err.Error()))
	}
	//PrintError(err)

	if intID <= 0 {
		return quizController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an quiz!",
		})
	}

	quiz, err := quizController.quizService.Get(intID)
	if err != nil {
		return quizController.responder.Data(definitions.GenericAPIMessage{
			Status:  500,
			Message: "We cannot process your request. Please try again or contact support!",
		})
	}

	if quiz.ID <= 0 {
		return quizController.responder.Data(definitions.GenericAPIMessage{
			Status:  404,
			Message: "quiz Not Found!",
		})
	}

	formError := req.Request().ParseForm()
	if formError != nil {
		return quizController.responder.HTTP(400, strings.NewReader(formError.Error()))
	}

	form := req.Request().Form

	quizData := &models.ChapterQuiz{
		Name:            funcs.ValidateStringFormKeys("Name", form, "string").(string),
		Description:     funcs.ValidateStringFormKeys("Description", form, "string").(string),
		Duration:        funcs.ValidateStringFormKeys("Duration", form, "int").(int),
		CourseChapterID: funcs.ValidateStringFormKeys("CourseChapterID", form, "int").(int),
		CreatedAt:       time.Now(),
	}

	//fmt.Printf("validator: %+v\n", quizController.validator.validate)
	validateError := quizController.customValidator.Validate.Struct(quizData)
	if validateError != nil {
		errorResponse := funcs.ErrorPackagingForMaps(quizController.customValidator.TranslateError(validateError))
		errorResponse, packError := funcs.ErrorPackaging(errorResponse, 400)
		if packError != nil {
			return quizController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return quizController.responder.HTTP(400, strings.NewReader(errorResponse))
	}

	result, err := quizController.quizService.Update(intID, *quizData)
	if err != nil {
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return quizController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return quizController.responder.HTTP(500, strings.NewReader(errorResponse))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return quizController.responder.HTTP(400, strings.NewReader(resErr.Error()))
	}

	return quizController.responder.HTTP(uint(result.Status), strings.NewReader(string(res)))
}

func (quizController *QuizController) Delete(ctx context.Context, req *web.Request) web.Result {
	if req.Params["id"] == "" {
		return quizController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an quiz!",
		})
	}

	intID, err := strconv.Atoi(req.Params["id"])
	//PrintError(err)

	if intID <= 0 {
		return quizController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an quiz!",
		})
	}

	quiz, err := quizController.quizService.Get(intID)
	if err != nil {
		return quizController.responder.Data(definitions.GenericAPIMessage{
			Status:  500,
			Message: "We cannot process your request. Please try again or contact support!",
		})
	}

	if quiz.ID <= 0 {
		return quizController.responder.Data(definitions.GenericAPIMessage{
			Status:  404,
			Message: "quiz Not Found!",
		})
	}

	result, err := quizController.quizService.Delete(intID)
	if err != nil {
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return quizController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return quizController.responder.HTTP(500, strings.NewReader(errorResponse))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return quizController.responder.HTTP(400, strings.NewReader(resErr.Error()))
	}

	return quizController.responder.HTTP(uint(result.Status), strings.NewReader(string(res)))
}

func (quizController *QuizController) Create(ctx context.Context, req *web.Request) web.Result {
	formError := req.Request().ParseForm()
	if formError != nil {
		return quizController.responder.HTTP(400, strings.NewReader(formError.Error()))
	}

	form := req.Request().Form

	quiz := &models.ChapterQuiz{
		Name:            funcs.ValidateStringFormKeys("Name", form, "string").(string),
		Description:     funcs.ValidateStringFormKeys("Description", form, "string").(string),
		Duration:        funcs.ValidateStringFormKeys("Duration", form, "int").(int),
		CourseChapterID: funcs.ValidateStringFormKeys("CourseChapterID", form, "int").(int),
		CreatedAt:       time.Now(),
	}

	//fmt.Printf("validator: %+v\n", instructorController.validator.validate)
	validateError := quizController.customValidator.Validate.Struct(quiz)
	if validateError != nil {
		errorResponse := funcs.ErrorPackagingForMaps(quizController.customValidator.TranslateError(validateError))
		errorResponse, packError := funcs.ErrorPackaging(errorResponse, 400)
		if packError != nil {
			return quizController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return quizController.responder.HTTP(400, strings.NewReader(errorResponse))
	}

	result, err := quizController.quizService.Create(*quiz)
	if err != nil {
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return quizController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return quizController.responder.HTTP(500, strings.NewReader(errorResponse))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return quizController.responder.HTTP(400, strings.NewReader(resErr.Error()))
	}

	return quizController.responder.HTTP(uint(result.Status), strings.NewReader(string(res)))

}
