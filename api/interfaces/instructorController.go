type (
	InstructorController struct {
		responder         *web.Responder
		instructorService service.InstructorServiceInterface
	}

	GetInstructorResponse struct {
		Instructor entity.Instructor
	}
)