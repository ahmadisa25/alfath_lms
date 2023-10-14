package api

import (
	"alfath_lms/api/controllers"

	"flamingo.me/flamingo/v3/framework/web"
)

type Routes struct {
	instructorController controllers.InstructorController
	studentController    controllers.StudentController
}

func (routes *Routes) Inject(
	instructorController *controllers.InstructorController,
	studentController *controllers.StudentController,
) {
	routes.instructorController = *instructorController
	routes.studentController = *studentController
}

func (routes *Routes) Routes(registry *web.RouterRegistry) {
	registry.Route("/instructor/:id", "instructor")
	registry.HandleGet("instructor", routes.instructorController.Get)
	registry.HandlePut("instructor", routes.instructorController.Update)
	registry.HandleDelete("instructor", routes.instructorController.Delete)

	registry.Route("/instructor/", "instructor")
	registry.HandlePost("instructor", routes.instructorController.Create)

	registry.Route("/instructor-all/", "instructor-all")
	registry.HandleGet("instructor-all", routes.instructorController.GetAll)

	registry.Route("/student/:id", "student")
	registry.HandleGet("student", routes.studentController.Get)
	registry.HandlePut("student", routes.studentController.Update)
	registry.HandleDelete("student", routes.studentController.Delete)

	registry.Route("/student/", "student")
	registry.HandlePost("student", routes.studentController.Create)

	registry.Route("/student-all/", "student-all")
	registry.HandleGet("student-all", routes.studentController.GetAll)
}
