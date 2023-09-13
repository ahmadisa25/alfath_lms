package interfaces

import (
	"context"
	"alfath_lms/instructor/domain/entity"
	"alfath_lms/instructor/domain/service"
	"flamingo.me/flamingo/v3/framework/web"
)

type (
	InstructorController struct {
		responder         *web.Responder
		instructorService service.InstructorServiceInterface
	}

	GetInstructorResponse struct {
		Instructor entity.Instructor
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
	instructorID, err := req.Query1(1)
	PrintError(err)

	instructor, err := instructorController.instructorService.GetInstructor(instructorID)
	PrintError(err)

	return instructorController.responder.data(GetInstructorResponse{
		Instructor:instructor
	})
}


func PrintError(err error) error {
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	return nil
}