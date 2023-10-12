package api

import (
	"alfath_lms/api/deps/db"
	"alfath_lms/api/deps/pagination"
	"alfath_lms/api/deps/validator"
	instructor_service "alfath_lms/api/instructor/domain/service"
	instructor_infrastructure "alfath_lms/api/instructor/infrastructure"
	student_service "alfath_lms/api/student/domain/service"
	student_infrastructure "alfath_lms/api/student/infrastructure"

	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/web"
)

type Module struct{}

func (module *Module) Configure(injector *dingo.Injector) {
	web.BindRoutes(injector, new(Routes))
	/*if os.Getenv("fake") == "true" {
		injector.Bind(new(service.InstructorServiceInterface)).To(infrastructure.FakeOrderService{})
	} else {*/
	injector.Bind(new(instructor_service.InstructorServiceInterface)).To(instructor_infrastructure.InstructorService{})
	injector.Bind(new(student_service.StudentServiceInterface)).To(student_infrastructure.StudentService{})
	//}
}

func (module *Module) Depends() []dingo.Module {
	return []dingo.Module{
		new(pagination.Module),
		new(db.Module),
		new(validator.Module),
	}
}
