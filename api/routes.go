package api

import (
	"alfath_lms/api/controllers"
	"alfath_lms/api/middleware"
	"context"

	"flamingo.me/flamingo/v3/framework/web"
)

type Routes struct {
	questionController   controllers.QuestionController
	quizController       controllers.QuizController
	answerController     controllers.AnswerController
	materialController   controllers.MaterialController
	chapterController    controllers.ChapterController
	courseController     controllers.CourseController
	instructorController controllers.InstructorController
	studentController    controllers.StudentController
	userController       controllers.UserController
	authMdw              *middleware.AuthMiddleware
	responder            *web.Responder
}

func (routes *Routes) Inject(
	questionController *controllers.QuestionController,
	answerController *controllers.AnswerController,
	quizController *controllers.QuizController,
	materialController *controllers.MaterialController,
	chapterController *controllers.ChapterController,
	courseController *controllers.CourseController,
	instructorController *controllers.InstructorController,
	studentController *controllers.StudentController,
	userController *controllers.UserController,
	responder *web.Responder,
) {
	routes.answerController = *answerController
	routes.quizController = *quizController
	routes.materialController = *materialController
	routes.courseController = *courseController
	routes.chapterController = *chapterController
	routes.instructorController = *instructorController
	routes.studentController = *studentController
	routes.questionController = *questionController
	routes.userController = *userController
	routes.authMdw = &middleware.AuthMiddleware{
		Responder: responder,
	}
	routes.responder = responder
}

func (routes *Routes) Routes(registry *web.RouterRegistry) {
	registry.Route("/instructor/:id", "instructor")
	registry.HandleGet("instructor", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.instructorController.Get, "all")
	})
	registry.HandlePut("instructor", routes.instructorController.Update)
	registry.HandleDelete("instructor", routes.instructorController.Delete)

	registry.Route("/instructor/", "instructor")
	registry.HandlePost("instructor", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.instructorController.Create, "administrator")
	})

	registry.Route("/instructor-all/", "instructor-all")
	registry.HandleGet("instructor-all", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.instructorController.GetAll, "all")
	})

	registry.Route("/course-all/", "course-all")
	registry.HandleGet("course-all", routes.courseController.GetAll)

	registry.Route("/course/", "course")
	registry.HandlePost("course", routes.courseController.Create)

	registry.Route("/course/:id", "course")
	registry.HandleGet("course", routes.courseController.Get)
	registry.HandleDelete("course", routes.courseController.Delete)
	registry.HandlePut("course", routes.courseController.Update)

	registry.Route("/course-chapter/", "course-chapter")
	registry.HandlePost("course-chapter", routes.chapterController.Create)

	registry.Route("/course-chapter/:id", "course-chapter")
	registry.HandleGet("course-chapter", routes.chapterController.Get)
	registry.HandleDelete("course-chapter", routes.chapterController.Delete)
	registry.HandlePut("course-chapter", routes.chapterController.Update)

	registry.Route("/chapter-material/", "chapter-material")
	registry.HandlePost("chapter-material", routes.materialController.Create)

	registry.Route("/chapter-material/:id", "chapter-material")
	registry.HandleGet("chapter-material", routes.materialController.Get)
	registry.HandleDelete("chapter-material", routes.materialController.Delete)
	registry.HandlePut("chapter-material", routes.materialController.Update)

	registry.Route("/chapter-quiz/", "chapter-quiz")
	registry.HandlePost("chapter-quiz", routes.quizController.Create)

	registry.Route("/chapter-quiz/:id", "chapter-quiz")
	registry.HandleGet("chapter-quiz", routes.quizController.Get)
	registry.HandleDelete("chapter-quiz", routes.quizController.Delete)
	registry.HandlePut("chapter-quiz", routes.quizController.Update)

	registry.Route("/quiz-question/", "quiz-question")
	registry.HandlePost("quiz-question", routes.questionController.Create)

	registry.Route("/quiz-question/:id", "quiz-question")
	registry.HandleGet("quiz-question", routes.questionController.Get)
	registry.HandleDelete("quiz-question", routes.questionController.Delete)
	registry.HandlePut("quiz-question", routes.questionController.Update)

	registry.Route("/quiz-answer/", "quiz-answer")
	registry.HandlePost("quiz-answer", routes.answerController.Create)

	registry.Route("/quiz-answer/:id", "quiz-answer")
	registry.HandleGet("quiz-answer", routes.answerController.Get)
	registry.HandleDelete("quiz-answer", routes.answerController.Delete)
	registry.HandlePut("quiz-answer", routes.answerController.Update)

	registry.Route("/user/", "user")
	registry.HandlePost("user", routes.userController.Create)

	registry.Route("/login/", "login")
	registry.HandlePost("login", routes.userController.Login)

	registry.Route("/student/:id", "student")
	registry.HandleGet("student", routes.studentController.Get)
	registry.HandlePut("student", routes.studentController.Update)
	registry.HandleDelete("student", routes.studentController.Delete)

	registry.Route("/student/", "student")
	registry.HandlePost("student", routes.studentController.Create)

	registry.Route("/student-all/", "student-all")
	registry.HandleGet("student-all", routes.studentController.GetAll)
}
