package api

import (
	"alfath_lms/api/deps/db"
	"alfath_lms/api/deps/validator"
	"alfath_lms/api/deps/pagination"
	"alfath_lms/api/instructor/domain/service"
	"alfath_lms/api/instructor/infrastructure"

	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/web"
)

type Module struct{}

func (module *Module) Configure(injector *dingo.Injector) {
	web.BindRoutes(injector, new(Routes))
	/*if os.Getenv("fake") == "true" {
		injector.Bind(new(service.InstructorServiceInterface)).To(infrastructure.FakeOrderService{})
	} else {*/
	injector.Bind(new(service.InstructorServiceInterface)).To(infrastructure.InstructorService{})
	//}
}

func (module *Module) Depends() []dingo.Module {
	return []dingo.Module{
		new(pagination.Module),
		new(db.Module),
		new(validator.Module),
	}
}
