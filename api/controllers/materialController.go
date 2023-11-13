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
	MaterialController struct {
		responder       *web.Responder
		customValidator *validator.CustomValidator
		materialService interfaces.MaterialServiceInterface
	}

	GetMaterialResponse struct {
		Status int
		Data   models.ChapterMaterial
	}
)

func (materialController *MaterialController) Inject(
	responder *web.Responder,
	customValidator *validator.CustomValidator,
	materialService interfaces.MaterialServiceInterface,
) {
	materialController.responder = responder
	materialController.customValidator = customValidator
	materialController.materialService = materialService
}

func (materialController *MaterialController) Create(ctx context.Context, req *web.Request) web.Result {
	fmt.Println("test")
	formError := req.Request().ParseForm()
	if formError != nil {
		return materialController.responder.HTTP(400, strings.NewReader(formError.Error()))
	}

	form := req.Request().Form

	material := &models.ChapterMaterial{
		Name:            funcs.ValidateStringFormKeys("Name", form, "string").(string),
		Description:     funcs.ValidateStringFormKeys("Description", form, "string").(string),
		FileUrl:         funcs.ValidateStringFormKeys("FileUrl", form, "string").(string),
		CourseChapterID: funcs.ValidateStringFormKeys("CourseID", form, "int").(int),
		CreatedAt:       time.Now(),
	}

	//fmt.Printf("validator: %+v\n", materialController.validator.validate)
	validateError := materialController.customValidator.Validate.Struct(material)
	if validateError != nil {
		errorResponse := funcs.ErrorPackagingForMaps(materialController.customValidator.TranslateError(validateError))
		fmt.Println(errorResponse)
		errorResponse, packError := funcs.ErrorPackaging(errorResponse, 400)
		if packError != nil {
			return materialController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return materialController.responder.HTTP(400, strings.NewReader(errorResponse))
	}

	result, err := materialController.materialService.Create(*material)
	if err != nil {
		fmt.Println(err)
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return materialController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return materialController.responder.HTTP(500, strings.NewReader(errorResponse))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return materialController.responder.HTTP(400, strings.NewReader(resErr.Error()))
	}

	return materialController.responder.HTTP(uint(result.Status), strings.NewReader(string(res)))

}

func (materialController *MaterialController) Get(ctx context.Context, req *web.Request) web.Result {
	if req.Params["id"] == "" {
		return materialController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an material!",
		})
	}

	intID, err := strconv.Atoi(req.Params["id"])
	//PrintError(err)

	if intID <= 0 {
		return materialController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select a material!",
		})
	}

	material, err := materialController.materialService.Get(intID)
	if err != nil {
		return materialController.responder.Data(definitions.GenericAPIMessage{
			Status:  500,
			Message: "We cannot process your request. Please try again or contact support!",
		})
	}

	if material.ID <= 0 {
		return materialController.responder.Data(definitions.GenericAPIMessage{
			Status:  404,
			Message: "material Not Found!",
		})
	}

	return materialController.responder.Data(GetMaterialResponse{
		Status: 200,
		Data:   material,
	})
}

func (materialController *MaterialController) Delete(ctx context.Context, req *web.Request) web.Result {
	if req.Params["id"] == "" {
		return materialController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an material!",
		})
	}

	intID, err := strconv.Atoi(req.Params["id"])
	//PrintError(err)

	if intID <= 0 {
		return materialController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an material!",
		})
	}

	material, err := materialController.materialService.Get(intID)
	if err != nil {
		return materialController.responder.Data(definitions.GenericAPIMessage{
			Status:  500,
			Message: "We cannot process your request. Please try again or contact support!",
		})
	}

	if material.ID <= 0 {
		return materialController.responder.Data(definitions.GenericAPIMessage{
			Status:  404,
			Message: "material Not Found!",
		})
	}

	result, err := materialController.materialService.Delete(intID)
	if err != nil {
		fmt.Println(err)
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return materialController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return materialController.responder.HTTP(500, strings.NewReader(errorResponse))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return materialController.responder.HTTP(400, strings.NewReader(resErr.Error()))
	}

	return materialController.responder.HTTP(uint(result.Status), strings.NewReader(string(res)))
}

func (materialController *MaterialController) Update(ctx context.Context, req *web.Request) web.Result {
	if req.Params["id"] == "" {
		return materialController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an material!",
		})
	}

	intID, err := strconv.Atoi(req.Params["id"])
	if err != nil {
		return materialController.responder.HTTP(500, strings.NewReader(err.Error()))
	}
	//PrintError(err)

	if intID <= 0 {
		return materialController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an material!",
		})
	}

	material, err := materialController.materialService.Get(intID)
	if err != nil {
		return materialController.responder.Data(definitions.GenericAPIMessage{
			Status:  500,
			Message: "We cannot process your request. Please try again or contact support!",
		})
	}

	if material.ID <= 0 {
		return materialController.responder.Data(definitions.GenericAPIMessage{
			Status:  404,
			Message: "material Not Found!",
		})
	}
	formError := req.Request().ParseForm()
	if formError != nil {
		return materialController.responder.HTTP(400, strings.NewReader(formError.Error()))
	}

	form := req.Request().Form

	materialData := &models.ChapterMaterial{
		Name:            funcs.ValidateStringFormKeys("Name", form, "string").(string),
		Description:     funcs.ValidateStringFormKeys("Description", form, "string").(string),
		FileUrl:         funcs.ValidateStringFormKeys("FileUrl", form, "string").(string),
		CourseChapterID: funcs.ValidateStringFormKeys("CourseChapterID", form, "int").(int),
		CreatedAt:       time.Now(),
	}

	//fmt.Printf("validator: %+v\n", materialController.validator.validate)
	validateError := materialController.customValidator.Validate.Struct(materialData)
	if validateError != nil {
		errorResponse := funcs.ErrorPackagingForMaps(materialController.customValidator.TranslateError(validateError))
		errorResponse, packError := funcs.ErrorPackaging(errorResponse, 400)
		if packError != nil {
			return materialController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return materialController.responder.HTTP(400, strings.NewReader(errorResponse))
	}

	result, err := materialController.materialService.Update(intID, *materialData)
	if err != nil {
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return materialController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return materialController.responder.HTTP(500, strings.NewReader(errorResponse))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return materialController.responder.HTTP(400, strings.NewReader(resErr.Error()))
	}

	return materialController.responder.HTTP(uint(result.Status), strings.NewReader(string(res)))
}
