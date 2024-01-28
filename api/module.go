package api

import (
	"alfath_lms/api/deps/amqp"
	"alfath_lms/api/deps/db"
	"alfath_lms/api/deps/mongo"
	"alfath_lms/api/deps/pagination"
	"alfath_lms/api/deps/redis"
	"alfath_lms/api/deps/validator"
	"alfath_lms/api/interfaces"
	"alfath_lms/api/services"

	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/web"
)

type Module struct{}

func (module *Module) Configure(injector *dingo.Injector) {
	web.BindRoutes(injector, new(Routes))
	/*if os.Getenv("fake") == "true" {
		injector.Bind(new(service.InstructorServiceInterface)).To(infrastructure.FakeOrderService{})
	} else {*/
	injector.Bind(new(interfaces.UserServiceInterface)).To(services.UserService{})
	injector.Bind(new(interfaces.CourseServiceInterface)).To(services.CourseService{})
	injector.Bind(new(interfaces.InstructorServiceInterface)).To(services.InstructorService{})
	injector.Bind(new(interfaces.StudentServiceInterface)).To(services.StudentService{})
	injector.Bind(new(interfaces.ChapterServiceInterface)).To(services.ChapterService{})
	injector.Bind(new(interfaces.MaterialServiceInterface)).To(services.MaterialService{})
	injector.Bind(new(interfaces.QuizServiceInterface)).To(services.QuizService{})
	injector.Bind(new(interfaces.AnswerServiceInterface)).To(services.AnswerService{})
	injector.Bind(new(interfaces.QuestionServiceInterface)).To(services.QuestionService{})
	injector.Bind(new(interfaces.AnnouncementServiceInterface)).To(services.AnnouncementService{})
	//}
}

func (module *Module) Depends() []dingo.Module {
	return []dingo.Module{
		new(pagination.Module),
		new(db.Module),
		new(validator.Module),
		new(mongo.Module),
		new(redis.Module),
		new(amqp.Module),
	}
}
