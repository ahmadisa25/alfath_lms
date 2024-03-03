package controllers

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/deps/validator"
	"alfath_lms/api/funcs"
	"alfath_lms/api/interfaces"
	"alfath_lms/api/models"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"flamingo.me/flamingo/v3/framework/web"
)

type (
	AnswerController struct {
		responder       *web.Responder
		customValidator *validator.CustomValidator
		answerService   interfaces.AnswerServiceInterface
	}

	GetAnswerResponse struct {
		Status int
		Data   models.QuizAnswer
	}
)

func (answerController *AnswerController) Inject(
	responder *web.Responder,
	customValidator *validator.CustomValidator,
	answerService interfaces.AnswerServiceInterface,
) {
	answerController.responder = responder
	answerController.customValidator = customValidator
	answerController.answerService = answerService
}

func (answerController *AnswerController) Get(ctx context.Context, req *web.Request) web.Result {
	if req.Params["id"] == "" {
		return answerController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an answer!",
		})
	}

	intID, err := strconv.Atoi(req.Params["id"])
	//PrintError(err)

	if intID <= 0 {
		return answerController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an answer!",
		})
	}

	answer, err := answerController.answerService.Get(intID)
	if err != nil {
		return answerController.responder.Data(definitions.GenericAPIMessage{
			Status:  500,
			Message: "We cannot process your request. Please try again or contact support!",
		})
	}

	if answer.ID <= 0 {
		return answerController.responder.Data(definitions.GenericAPIMessage{
			Status:  404,
			Message: "answer Not Found!",
		})
	}

	return answerController.responder.Data(GetAnswerResponse{
		Status: 200,
		Data:   answer,
	})
}

func (answerController *AnswerController) Update(ctx context.Context, req *web.Request) web.Result {
	if req.Params["id"] == "" {
		return answerController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an answer!",
		})
	}

	intID, err := strconv.Atoi(req.Params["id"])
	if err != nil {
		return answerController.responder.HTTP(500, strings.NewReader(err.Error()))
	}
	//PrintError(err)

	if intID <= 0 {
		return answerController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an answer!",
		})
	}

	answer, err := answerController.answerService.Get(intID)
	if err != nil {
		return answerController.responder.Data(definitions.GenericAPIMessage{
			Status:  500,
			Message: "We cannot process your request. Please try again or contact support!",
		})
	}

	if answer.ID <= 0 {
		return answerController.responder.Data(definitions.GenericAPIMessage{
			Status:  404,
			Message: "answer Not Found!",
		})
	}

	formError := req.Request().ParseForm()
	if formError != nil {
		return answerController.responder.HTTP(400, strings.NewReader(formError.Error()))
	}

	form := req.Request().Form

	answerData := &models.QuizAnswer{
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

func (answerController *AnswerController) GetAll(ctx context.Context, req *web.Request) web.Result {
	query := req.QueryAll()
	paginationReq := definitions.PaginationRequest{
		SelectedColumns: funcs.ValidateStringFormKeys("select", query, "string").(string),
		Search:          funcs.ValidateStringFormKeys("search", query, "string").(string),
		Page:            funcs.ValidateStringFormKeys("page", query, "string").(string),
		PerPage:         funcs.ValidateStringFormKeys("perpage", query, "string").(string),
		OrderBy:         funcs.ValidateStringFormKeys("order", query, "string").(string),
		Filter:          funcs.ValidateStringFormKeys("filter", query, "string").(string),
	}

	result, err := answerController.answerService.GetAll(paginationReq)
	if err != nil {
		fmt.Println(err)
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return funcs.CorsedResponse(answerController.responder.HTTP(500, strings.NewReader(packError.Error())))
		}
		return funcs.CorsedResponse(answerController.responder.HTTP(500, strings.NewReader(errorResponse)))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return funcs.CorsedResponse(answerController.responder.HTTP(400, strings.NewReader(resErr.Error())))
	}

	return funcs.CorsedResponse(answerController.responder.HTTP(uint(result.Status), strings.NewReader(string(res))))

}

func (answerController *AnswerController) GetAllDistinct(ctx context.Context, req *web.Request) web.Result {
	query := req.QueryAll()
	paginationReq := definitions.PaginationRequest{
		SelectedColumns: funcs.ValidateStringFormKeys("select", query, "string").(string),
		Search:          funcs.ValidateStringFormKeys("search", query, "string").(string),
		Page:            funcs.ValidateStringFormKeys("page", query, "string").(string),
		PerPage:         funcs.ValidateStringFormKeys("perpage", query, "string").(string),
		OrderBy:         funcs.ValidateStringFormKeys("order", query, "string").(string),
		Filter:          funcs.ValidateStringFormKeys("filter", query, "string").(string),
	}

	result, err := answerController.answerService.GetAll(paginationReq)
	if err != nil {
		fmt.Println(err)
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return funcs.CorsedResponse(answerController.responder.HTTP(500, strings.NewReader(packError.Error())))
		}
		return funcs.CorsedResponse(answerController.responder.HTTP(500, strings.NewReader(errorResponse)))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return funcs.CorsedResponse(answerController.responder.HTTP(400, strings.NewReader(resErr.Error())))
	}

	return funcs.CorsedResponse(answerController.responder.HTTP(uint(result.Status), strings.NewReader(string(res))))

}

func (answerController *AnswerController) Delete(ctx context.Context, req *web.Request) web.Result {
	if req.Params["id"] == "" {
		return answerController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an answer!",
		})
	}

	intID, err := strconv.Atoi(req.Params["id"])
	//PrintError(err)

	if intID <= 0 {
		return answerController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an answer!",
		})
	}

	answer, err := answerController.answerService.Get(intID)
	if err != nil {
		return answerController.responder.Data(definitions.GenericAPIMessage{
			Status:  500,
			Message: "We cannot process your request. Please try again or contact support!",
		})
	}

	if answer.ID <= 0 {
		return answerController.responder.Data(definitions.GenericAPIMessage{
			Status:  404,
			Message: "answer Not Found!",
		})
	}

	result, err := answerController.answerService.Delete(intID)
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

func (answerController *AnswerController) Create(ctx context.Context, req *web.Request) web.Result {
	formError := req.Request().ParseForm()
	if formError != nil {
		return answerController.responder.HTTP(400, strings.NewReader(formError.Error()))
	}

	form := req.Request().Form

	answer := &models.QuizAnswer{
		Answer:         funcs.ValidateStringFormKeys("Answer", form, "string").(string),
		QuizQuestionID: funcs.ValidateStringFormKeys("QuizQuestionID", form, "int").(int),
		StudentID:      funcs.ValidateStringFormKeys("StudentID", form, "int").(int),
		CreatedAt:      time.Now(),
	}

	//fmt.Printf("validator: %+v\n", instructorController.validator.validate)
	validateError := answerController.customValidator.Validate.Struct(answer)
	if validateError != nil {
		errorResponse := funcs.ErrorPackagingForMaps(answerController.customValidator.TranslateError(validateError))
		errorResponse, packError := funcs.ErrorPackaging(errorResponse, 400)
		if packError != nil {
			return answerController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return answerController.responder.HTTP(400, strings.NewReader(errorResponse))
	}

	result, err := answerController.answerService.Create(*answer)
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
