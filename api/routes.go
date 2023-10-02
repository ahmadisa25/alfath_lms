package api

import (
	"alfath_lms/api/interfaces"
	"flamingo.me/flamingo/v3/framework/web"
)

type Routes struct {
	instructorController interfaces.InstructorController
}

func (routes *Routes) Inject(instructorController *interfaces.InstructorController) {
	routes.instructorController = *instructorController
}

func (routes *Routes) Routes(registry *web.RouterRegistry) {
	registry.Route("/instructor/:id", "instructor")
	registry.HandleGet("instructor", routes.instructorController.Get)
	registry.HandlePut("instructor", routes.instructorController.Update)

	registry.Route("/instructor/", "instructor")
	registry.HandlePost("instructor", routes.instructorController.Create)
}