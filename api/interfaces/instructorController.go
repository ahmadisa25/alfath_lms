package interfaces

import (
	"alfath_lms/instructor/domain/entity"
	"alfath_lms/instructor/domain/service"
	"alfath_lms/api/definitions"
	"context"
	"fmt"
	"strings"
	"github.com/go-playground/validator/v10"
	"strconv"
	"flamingo.me/flamingo/v3/framework/web"
)

type (
	InstructorController struct {
		responder         *web.Responder
		validate		  *validator.Validate
		instructorService service.InstructorServiceInterface
	}

	GetInstructorResponse struct {
		Status int
		Data entity.Instructor
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
	fmt.Println(req)
	if len(req.Params) == 0{
		return instructorController.responder.HTTP(400, strings.NewReader("Please provide a valid request body"))
	}

	instructor := &entity.Instructor{
		Name: req.Params["Name"],
		Email: req.Params["Email"],
		MobilePhone: req.Params["MobilePhone"],
	}
	fmt.Println(*instructor)
	err := instructorController.validate.Struct(instructor)
	fmt.Println(err)

	result, err	:= instructorController.instructorService.CreateInstructor(*instructor)
	if err != nil{
		return instructorController.responder.HTTP(500, strings.NewReader("Failed to create instructor. Please contract support or try again"))
	}

	return instructorController.responder.HTTP(uint(result.Status), strings.NewReader("Instructor created successfully"))

}	

func (instructorController *InstructorController) Get(ctx context.Context, req *web.Request) web.Result {
	if req.Params["id"] == "" {
		return instructorController.responder.Data(definitions.GenericAPIMessage{
			Status: 400,
			Message: "Please select an instructor!",
		})
	}

	intID, err := strconv.Atoi(req.Params["id"]);
	//PrintError(err)

	if intID <=0 {
		return instructorController.responder.Data(definitions.GenericAPIMessage{
			Status: 400,
			Message: "Please select an instructor!",
		})
	}

	instructor, err := instructorController.instructorService.GetInstructor(req.Params["id"])
	if err != nil {
		return instructorController.responder.Data(definitions.GenericAPIMessage{
			Status: 500,
			Message: "We cannot process your request. Please try again or contact support!",
		})
	}

	if instructor.ID <= 0{
		return instructorController.responder.Data(definitions.GenericAPIMessage{
			Status: 404,
			Message: "Instructor Not Found!",
		})
	}

	return instructorController.responder.Data(GetInstructorResponse{
		Status: 200,
		Data: instructor,
	})
}

func PrintError(err error) error {
	fmt.Println(err)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	return nil
}
