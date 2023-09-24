package interfaces

import (
	"alfath_lms/instructor/domain/entity"
	"alfath_lms/instructor/domain/service"
	"alfath_lms/api/definitions"
	"context"
	"fmt"
	"strconv"
	"flamingo.me/flamingo/v3/framework/web"
)

type (
	InstructorController struct {
		responder         *web.Responder
		instructorService service.InstructorServiceInterface
	}

	GetInstructorResponse struct {
		Status int
		Data entity.Instructor
	}
)

func (instructorController *InstructorController) Inject(
	responder *web.Responder,
	instructorService service.InstructorServiceInterface,
) {
	instructorController.responder = responder
	instructorController.instructorService = instructorService
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
