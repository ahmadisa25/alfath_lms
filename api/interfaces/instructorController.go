package interfaces

import (
	"alfath_lms/api/definitions"
	"alfath_lms/instructor/domain/entity"
	"alfath_lms/instructor/domain/service"
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"flamingo.me/flamingo/v3/framework/web"
	"github.com/go-playground/validator/v10"
)

type (
	InstructorController struct {
		responder         *web.Responder
		validate          *validator.Validate
		instructorService service.InstructorServiceInterface
	}

	GetInstructorResponse struct {
		Status int
		Data   entity.Instructor
	}
)

func (instructorController *InstructorController) Inject(
	responder *web.Responder,
	validate *validator.Validate,
	instructorService service.InstructorServiceInterface,
) {
	instructorController.responder = responder
	instructorController.validate = validate
	instructorController.instructorService = instructorService
}

func (instructorController *InstructorController) Create(ctx context.Context, req *web.Request) web.Result {
	formError := req.Request().ParseForm()
	if formError != nil {
		return instructorController.responder.HTTP(400, strings.NewReader(formError.Error()))
	}

	form := req.Request().Form

	name, nameOk := form["Name"]
	if !nameOk {
		return instructorController.responder.HTTP(400, strings.NewReader("You must put in instructor name"))
	}

	email, emailOk := form["Email"]
	if !emailOk {
		return instructorController.responder.HTTP(400, strings.NewReader("You must put in instructor email"))
	}

	mobilePhone, mobilePhoneOk := form["MobilePhone"]
	if !mobilePhoneOk {
		return instructorController.responder.HTTP(400, strings.NewReader("You must put in instructor mobile phone"))
	}

	instructor := &entity.Instructor{
		Name:        name[0],
		Email:       email[0],
		MobilePhone: mobilePhone[0],
		CreatedAt:   time.Now(),
	}
	fmt.Println(*instructor)
	validateError := instructorController.validate.Struct(instructor)
	if validateError != nil {
		return instructorController.responder.HTTP(400, strings.NewReader(validateError.Error()))
	}

	result, err := instructorController.instructorService.CreateInstructor(*instructor)
	if err != nil {
		return instructorController.responder.HTTP(500, strings.NewReader("Failed to create instructor. Please contract support or try again"))
	}

	return instructorController.responder.HTTP(uint(result.Status), strings.NewReader("Instructor created successfully"))

}

func (instructorController *InstructorController) Get(ctx context.Context, req *web.Request) web.Result {
	if req.Params["id"] == "" {
		return instructorController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an instructor!",
		})
	}

	intID, err := strconv.Atoi(req.Params["id"])
	//PrintError(err)

	if intID <= 0 {
		return instructorController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an instructor!",
		})
	}

	instructor, err := instructorController.instructorService.GetInstructor(req.Params["id"])
	if err != nil {
		return instructorController.responder.Data(definitions.GenericAPIMessage{
			Status:  500,
			Message: "We cannot process your request. Please try again or contact support!",
		})
	}

	if instructor.ID <= 0 {
		return instructorController.responder.Data(definitions.GenericAPIMessage{
			Status:  404,
			Message: "Instructor Not Found!",
		})
	}

	return instructorController.responder.Data(GetInstructorResponse{
		Status: 200,
		Data:   instructor,
	})
}

func PrintError(err error) error {
	fmt.Println(err)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	return nil
}
