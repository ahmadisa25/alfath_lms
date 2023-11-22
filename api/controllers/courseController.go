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
	CourseController struct {
		responder       *web.Responder
		customValidator *validator.CustomValidator
		courseService   interfaces.CourseServiceInterface
	}
)

func (courseController *CourseController) Inject(
	responder *web.Responder,
	customValidator *validator.CustomValidator,
	courseService interfaces.CourseServiceInterface,
) {
	courseController.responder = responder
	courseController.customValidator = customValidator
	courseController.courseService = courseService
}

func (courseController *CourseController) GetAll(ctx context.Context, req *web.Request) web.Result {
	query := req.QueryAll()
	paginationReq := definitions.PaginationRequest{
		SelectedColumns: funcs.ValidateStringFormKeys("select", query, "string").(string),
		Search:          funcs.ValidateStringFormKeys("search", query, "string").(string),
		Page:            funcs.ValidateStringFormKeys("page", query, "string").(string),
		PerPage:         funcs.ValidateStringFormKeys("perpage", query, "string").(string),
		OrderBy:         funcs.ValidateStringFormKeys("order", query, "string").(string),
		Filter:          funcs.ValidateStringFormKeys("filter", query, "string").(string),
	}

	result, err := courseController.courseService.GetAll(paginationReq)
	if err != nil {
		fmt.Println(err)
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return courseController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return courseController.responder.HTTP(500, strings.NewReader(errorResponse))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return courseController.responder.HTTP(400, strings.NewReader(resErr.Error()))
	}

	return courseController.responder.HTTP(uint(result.Status), strings.NewReader(string(res)))

}

func (courseController *CourseController) Get(ctx context.Context, req *web.Request) web.Result {
	if req.Params["id"] == "" {
		return courseController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an course!",
		})
	}

	intID, err := strconv.Atoi(req.Params["id"])
	//PrintError(err)

	if intID <= 0 {
		return courseController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an course!",
		})
	}

	course, err := courseController.courseService.Get(intID)
	if err != nil {
		return courseController.responder.Data(definitions.GenericAPIMessage{
			Status:  500,
			Message: "We cannot process your request. Please try again or contact support!",
		})
	}

	if course.ID <= 0 {
		return courseController.responder.Data(definitions.GenericAPIMessage{
			Status:  404,
			Message: "course Not Found!",
		})
	}

	return courseController.responder.Data(definitions.GenericGetMessage[models.Course]{
		Status: 200,
		Data:   course,
	})
}

func (courseController *CourseController) Update(ctx context.Context, req *web.Request) web.Result {
	if req.Params["id"] == "" {
		return courseController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an course!",
		})
	}

	intID, err := strconv.Atoi(req.Params["id"])
	if err != nil {
		return courseController.responder.HTTP(500, strings.NewReader(err.Error()))
	}
	//PrintError(err)

	if intID <= 0 {
		return courseController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an course!",
		})
	}

	course, err := courseController.courseService.Get(intID)
	if err != nil {
		return courseController.responder.Data(definitions.GenericAPIMessage{
			Status:  500,
			Message: "We cannot process your request. Please try again or contact support!",
		})
	}

	if course.ID <= 0 {
		return courseController.responder.Data(definitions.GenericAPIMessage{
			Status:  404,
			Message: "Course Not Found!",
		})
	}

	formError := req.Request().ParseForm()
	if formError != nil {
		return courseController.responder.HTTP(400, strings.NewReader(formError.Error()))
	}

	form := req.Request().Form

	instructorList := ""
	instructors, instructorsOk := form["Instructors"]
	if !instructorsOk {
		errorResponse, packError := funcs.ErrorPackaging("Please select instructors!", 500)
		if packError != nil {
			return courseController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return courseController.responder.HTTP(500, strings.NewReader(errorResponse))
	} else {
		instructorList = instructors[0]
	}

	courseData := &models.Course{
		Name:        funcs.ValidateStringFormKeys("Name", form, "string").(string),
		Description: funcs.ValidateStringFormKeys("Description", form, "string").(string),
		Duration:    funcs.ValidateStringFormKeys("Duration", form, "int").(int),
		Instructors: []*models.Instructor{},
	}

	//fmt.Printf("validator: %+v\n", courseController.validator.validate)
	validateError := courseController.customValidator.Validate.Struct(courseData)
	if validateError != nil {
		errorResponse := funcs.ErrorPackagingForMaps(courseController.customValidator.TranslateError(validateError))
		errorResponse, packError := funcs.ErrorPackaging(errorResponse, 400)
		if packError != nil {
			return courseController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return courseController.responder.HTTP(400, strings.NewReader(errorResponse))
	}

	result, err := courseController.courseService.Update(intID, *courseData, instructorList)
	if err != nil {
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return courseController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return courseController.responder.HTTP(500, strings.NewReader(errorResponse))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return courseController.responder.HTTP(400, strings.NewReader(resErr.Error()))
	}

	return courseController.responder.HTTP(uint(result.Status), strings.NewReader(string(res)))
}

func (courseController *CourseController) Delete(ctx context.Context, req *web.Request) web.Result {
	if req.Params["id"] == "" {
		return courseController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an course!",
		})
	}

	intID, err := strconv.Atoi(req.Params["id"])
	//PrintError(err)

	if intID <= 0 {
		return courseController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an course!",
		})
	}

	course, err := courseController.courseService.Get(intID)
	if err != nil {
		return courseController.responder.Data(definitions.GenericAPIMessage{
			Status:  500,
			Message: "We cannot process your request. Please try again or contact support!",
		})
	}

	if course.ID <= 0 {
		return courseController.responder.Data(definitions.GenericAPIMessage{
			Status:  404,
			Message: "Course Not Found!",
		})
	}

	result, err := courseController.courseService.Delete(intID)
	if err != nil {
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return courseController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return courseController.responder.HTTP(500, strings.NewReader(errorResponse))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return courseController.responder.HTTP(400, strings.NewReader(resErr.Error()))
	}

	return courseController.responder.HTTP(uint(result.Status), strings.NewReader(string(res)))
}

func (courseController *CourseController) Create(ctx context.Context, req *web.Request) web.Result {
	formError := req.Request().ParseForm()
	if formError != nil {
		return courseController.responder.HTTP(400, strings.NewReader(formError.Error()))
	}

	form := req.Request().Form

	instructorList := ""
	instructors, instructorsOk := form["Instructors"]
	if !instructorsOk {
		errorResponse, packError := funcs.ErrorPackaging("Please select instructors!", 500)
		if packError != nil {
			return courseController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return courseController.responder.HTTP(500, strings.NewReader(errorResponse))
	} else {
		instructorList = instructors[0]
	}

	course := &models.Course{
		Name:        funcs.ValidateStringFormKeys("Name", form, "string").(string),
		Description: funcs.ValidateStringFormKeys("Description", form, "string").(string),
		Duration:    funcs.ValidateStringFormKeys("Duration", form, "int").(int),
		CreatedAt:   time.Now(),
		Instructors: []*models.Instructor{},
	}

	//fmt.Printf("validator: %+v\n", instructorController.validator.validate)
	validateError := courseController.customValidator.Validate.Struct(course)
	if validateError != nil {
		errorResponse := funcs.ErrorPackagingForMaps(courseController.customValidator.TranslateError(validateError))
		errorResponse, packError := funcs.ErrorPackaging(errorResponse, 400)
		if packError != nil {
			return courseController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return courseController.responder.HTTP(400, strings.NewReader(errorResponse))
	}

	result, err := courseController.courseService.Create(*course, instructorList)
	if err != nil {
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return courseController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return courseController.responder.HTTP(500, strings.NewReader(errorResponse))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return courseController.responder.HTTP(400, strings.NewReader(resErr.Error()))
	}

	return courseController.responder.HTTP(uint(result.Status), strings.NewReader(string(res)))

}
