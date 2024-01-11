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
	"go.mongodb.org/mongo-driver/bson"
)

type (
	StudentController struct {
		responder       *web.Responder
		customValidator *validator.CustomValidator
		studentService  interfaces.StudentServiceInterface
		userService     interfaces.UserServiceInterface
	}

	GetStudentResponse struct {
		Status int
		Data   models.Student
	}
)

func (studentController *StudentController) Inject(
	responder *web.Responder,
	customValidator *validator.CustomValidator,
	studentService interfaces.StudentServiceInterface,
	userService interfaces.UserServiceInterface,
) {
	studentController.responder = responder
	studentController.customValidator = customValidator
	studentController.studentService = studentService
	studentController.userService = userService
}

func (studentController *StudentController) Create(ctx context.Context, req *web.Request) web.Result {
	formError := req.Request().ParseForm()
	if formError != nil {
		return funcs.CorsedResponse(studentController.responder.HTTP(400, strings.NewReader(formError.Error())))
	}

	form := req.Request().Form

	student := &models.Student{
		Name:        funcs.ValidateStringFormKeys("Name", form, "string").(string),
		Email:       funcs.ValidateStringFormKeys("Email", form, "string").(string),
		MobilePhone: funcs.ValidateStringFormKeys("MobilePhone", form, "string").(string),
		CreatedAt:   time.Now(),
	}

	//fmt.Printf("validator: %+v\n", studentController.validator.validate)
	validateError := studentController.customValidator.Validate.Struct(student)
	if validateError != nil {
		errorResponse := funcs.ErrorPackagingForMaps(studentController.customValidator.TranslateError(validateError))
		errorResponse, packError := funcs.ErrorPackaging(errorResponse, 400)
		if packError != nil {
			return funcs.CorsedResponse(studentController.responder.HTTP(500, strings.NewReader(packError.Error())))
		}
		return funcs.CorsedResponse(studentController.responder.HTTP(400, strings.NewReader(errorResponse)))
	}

	result, err := studentController.studentService.CreateStudent(*student)
	if err != nil {
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return funcs.CorsedResponse(studentController.responder.HTTP(500, strings.NewReader(packError.Error())))
		}
		return funcs.CorsedResponse(studentController.responder.HTTP(500, strings.NewReader(errorResponse)))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return funcs.CorsedResponse(studentController.responder.HTTP(400, strings.NewReader(resErr.Error())))
	}

	return funcs.CorsedResponse(studentController.responder.HTTP(uint(result.Status), strings.NewReader(string(res))))

}

func (studentController *StudentController) Get(ctx context.Context, req *web.Request) web.Result {
	if req.Params["id"] == "" {
		return funcs.CorsedDataResponse(studentController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an student!",
		}))
	}

	intID, err := strconv.Atoi(req.Params["id"])
	//PrintError(err)

	if intID <= 0 {
		return funcs.CorsedDataResponse(studentController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an student!",
		}))
	}

	student, err := studentController.studentService.GetStudent(intID)
	if err != nil {
		return funcs.CorsedDataResponse(studentController.responder.Data(definitions.GenericAPIMessage{
			Status:  500,
			Message: "We cannot process your request. Please try again or contact support!",
		}))
	}

	if student.ID <= 0 {
		return funcs.CorsedDataResponse(studentController.responder.Data(definitions.GenericAPIMessage{
			Status:  404,
			Message: "student Not Found!",
		}))
	}

	return funcs.CorsedDataResponse(studentController.responder.Data(GetStudentResponse{
		Status: 200,
		Data:   student,
	}))
}

func (studentController *StudentController) Delete(ctx context.Context, req *web.Request) web.Result {
	if req.Params["id"] == "" {
		return funcs.CorsedDataResponse(studentController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an student!",
		}))
	}

	intID, err := strconv.Atoi(req.Params["id"])
	//PrintError(err)

	if intID <= 0 {
		return funcs.CorsedDataResponse(studentController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an student!",
		}))
	}

	student, err := studentController.studentService.GetStudent(intID)
	if err != nil {
		return funcs.CorsedDataResponse(studentController.responder.Data(definitions.GenericAPIMessage{
			Status:  500,
			Message: "We cannot process your request. Please try again or contact support!",
		}))
	}

	if student.ID <= 0 {
		return funcs.CorsedDataResponse(studentController.responder.Data(definitions.GenericAPIMessage{
			Status:  404,
			Message: "student Not Found!",
		}))
	}

	result, err := studentController.studentService.DeleteStudent(intID)
	if err != nil {
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return funcs.CorsedResponse(studentController.responder.HTTP(500, strings.NewReader(packError.Error())))
		}
		return funcs.CorsedResponse(studentController.responder.HTTP(500, strings.NewReader(errorResponse)))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return funcs.CorsedResponse(studentController.responder.HTTP(400, strings.NewReader(resErr.Error())))
	}

	_, deleteUserErr := studentController.userService.Delete(student.Email)

	if deleteUserErr != nil {
		return funcs.CorsedResponse(studentController.responder.HTTP(400, strings.NewReader(deleteUserErr.Error())))
	}

	return funcs.CorsedResponse(studentController.responder.HTTP(uint(result.Status), strings.NewReader(string(res))))
}

func (studentController *StudentController) Update(ctx context.Context, req *web.Request) web.Result {
	if req.Params["id"] == "" {
		return funcs.CorsedDataResponse(studentController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an student!",
		}))
	}

	intID, err := strconv.Atoi(req.Params["id"])
	if err != nil {
		return funcs.CorsedDataResponse(studentController.responder.Data(definitions.GenericAPIMessage{
			Status:  500,
			Message: "We cannot process your request. Please try again or contact support!",
		}))
	}

	if intID <= 0 {
		return funcs.CorsedDataResponse(studentController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an student!",
		}))
	}

	student, err := studentController.studentService.GetStudent(intID)
	if err != nil {
		return funcs.CorsedDataResponse(studentController.responder.Data(definitions.GenericAPIMessage{
			Status:  500,
			Message: "We cannot process your request. Please try again or contact support!",
		}))
	}

	if student.ID <= 0 {
		return funcs.CorsedDataResponse(studentController.responder.Data(definitions.GenericAPIMessage{
			Status:  404,
			Message: "student Not Found!",
		}))
	}

	formError := req.Request().ParseForm()
	if formError != nil {
		return funcs.CorsedResponse(studentController.responder.HTTP(400, strings.NewReader(formError.Error())))
	}

	form := req.Request().Form

	studentData := &models.Student{
		Name:        funcs.ValidateOrOverwriteStringFormKeys("Name", form, "string", student).(string),
		Email:       funcs.ValidateOrOverwriteStringFormKeys("Email", form, "string", student).(string),
		MobilePhone: funcs.ValidateOrOverwriteStringFormKeys("MobilePhone", form, "string", student).(string),
	}

	//fmt.Printf("validator: %+v\n", studentController.validator.validate)
	validateError := studentController.customValidator.Validate.Struct(studentData)
	if validateError != nil {
		errorResponse := funcs.ErrorPackagingForMaps(studentController.customValidator.TranslateError(validateError))
		errorResponse, packError := funcs.ErrorPackaging(errorResponse, 400)
		if packError != nil {
			return funcs.CorsedResponse(studentController.responder.HTTP(500, strings.NewReader(packError.Error())))
		}
		return funcs.CorsedResponse(studentController.responder.HTTP(400, strings.NewReader(errorResponse)))
	}

	result, err := studentController.studentService.UpdateStudent(intID, *studentData, student)
	if err != nil {
		fmt.Println(err)
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return funcs.CorsedResponse(studentController.responder.HTTP(500, strings.NewReader(packError.Error())))
		}
		return funcs.CorsedResponse(studentController.responder.HTTP(500, strings.NewReader(errorResponse)))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return funcs.CorsedResponse(studentController.responder.HTTP(400, strings.NewReader(resErr.Error())))
	}

	studentController.userService.Update(studentData.Email, []bson.E{{"email", studentData.Email}, {"name", studentData.Name}, {"mobilephone", studentData.MobilePhone}})

	return funcs.CorsedResponse(studentController.responder.HTTP(uint(result.Status), strings.NewReader(string(res))))
}

func (studentController *StudentController) GetAll(ctx context.Context, req *web.Request) web.Result {
	query := req.QueryAll()
	paginationReq := definitions.PaginationRequest{
		SelectedColumns: funcs.ValidateStringFormKeys("select", query, "string").(string),
		Search:          funcs.ValidateStringFormKeys("search", query, "string").(string),
		Page:            funcs.ValidateStringFormKeys("page", query, "string").(string),
		PerPage:         funcs.ValidateStringFormKeys("perpage", query, "string").(string),
		OrderBy:         funcs.ValidateStringFormKeys("order", query, "string").(string),
		Filter:          funcs.ValidateStringFormKeys("filter", query, "string").(string),
	}

	result, err := studentController.studentService.GetAllStudents(paginationReq)
	if err != nil {
		fmt.Println(err)
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return funcs.CorsedResponse(studentController.responder.HTTP(500, strings.NewReader(packError.Error())))
		}
		return funcs.CorsedResponse(studentController.responder.HTTP(500, strings.NewReader(errorResponse)))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return funcs.CorsedResponse(studentController.responder.HTTP(400, strings.NewReader(resErr.Error())))
	}

	return funcs.CorsedResponse(studentController.responder.HTTP(uint(result.Status), strings.NewReader(string(res))))

}

/*func PrintError(err error) error {
	fmt.Println(err)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	return nil
}*/
