package api

import (
	"alfath_lms/instructor/domain/service"
	"alfath_lms/instructor/infrastructure"

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
