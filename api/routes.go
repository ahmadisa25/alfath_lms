package api

import (
	"alfath_lms/api/controllers"
	"alfath_lms/api/middleware"
	"context"

	"flamingo.me/flamingo/v3/framework/web"
)

type Routes struct {
	questionController     controllers.QuestionController
	quizController         controllers.QuizController
	answerController       controllers.AnswerController
	materialController     controllers.MaterialController
	chapterController      controllers.ChapterController
	courseController       controllers.CourseController
	instructorController   controllers.InstructorController
	studentController      controllers.StudentController
	userController         controllers.UserController
	announcementController controllers.AnnouncementController
	authMdw                *middleware.AuthMiddleware
	responder              *web.Responder
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
	announcementController *controllers.AnnouncementController,
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
	routes.announcementController = *announcementController
	routes.authMdw = &middleware.AuthMiddleware{
		Responder: responder,
	}
	routes.responder = responder
}

func (routes *Routes) Routes(registry *web.RouterRegistry) {
	registry.Route("/instructor/:id", "instructor")
	registry.HandleGet("instructor", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.instructorController.Get, nil)
	})
	registry.HandlePut("instructor", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.instructorController.Update, []string{"administrator", "instructor"})
	})

	registry.HandleDelete("instructor", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.instructorController.Delete, []string{"administrator"})
	})

	registry.Route("/instructor/", "instructor")
	registry.HandlePost("instructor", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.instructorController.Create, []string{"administrator"})
	})

	registry.Route("/instructor-all/", "instructor-all")
	registry.HandleGet("instructor-all", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.instructorController.GetAll, nil)
	})

	registry.Route("/course-all/", "course-all")
	registry.HandleGet("course-all", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.courseController.GetAll, nil)
	})

	registry.Route("/course/", "course")
	registry.HandlePost("course", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.courseController.Create, []string{"administrator", "instructor"})
	})

	registry.Route("/course/:id", "course")
	registry.HandleGet("course", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.courseController.Get, nil)
	})
	registry.HandleDelete("course", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.courseController.Delete, []string{"administrator", "instructor"})
	})
	registry.HandlePut("course", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.courseController.Update, []string{"administrator", "instructor"})
	})

	registry.Route("/course-chapter/", "course-chapter")
	registry.HandlePost("course-chapter", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.chapterController.Create, []string{"administrator", "instructor"})
	})

	registry.Route("/course-chapter/:id", "course-chapter")
	registry.HandleGet("course-chapter", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.chapterController.Get, nil)
	})
	registry.HandleDelete("course-chapter", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.chapterController.Delete, []string{"administrator", "instructor"})
	})
	registry.HandlePut("course-chapter", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.chapterController.Update, []string{"administrator", "instructor"})
	})

	registry.Route("/chapter-material/", "chapter-material")
	registry.HandlePost("chapter-material", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.materialController.Create, []string{"administrator", "instructor"})
	})

	registry.Route("/chapter-material/:id", "chapter-material")
	registry.HandleGet("chapter-material", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.materialController.Get, nil)
	})
	registry.HandleDelete("chapter-material", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.materialController.Delete, []string{"administrator", "instructor"})
	})
	registry.HandlePut("chapter-material", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.materialController.Update, []string{"administrator", "instructor"})
	})

	registry.Route("/chapter-quiz/", "chapter-quiz")
	registry.HandlePost("chapter-quiz", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.quizController.Create, []string{"administrator", "instructor"})
	})

	registry.Route("/chapter-quiz/:id", "chapter-quiz")
	registry.HandleGet("chapter-quiz", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.quizController.Get, nil)
	})
	registry.HandleDelete("chapter-quiz", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.quizController.Delete, []string{"administrator", "instructor"})
	})
	registry.HandlePut("chapter-quiz", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.quizController.Update, []string{"administrator", "instructor"})
	})

	registry.Route("/quiz-question/", "quiz-question")
	registry.HandlePost("quiz-question", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.questionController.Create, []string{"administrator", "instructor"})
	})

	registry.Route("/quiz-question/:id", "quiz-question")
	registry.HandleGet("quiz-question", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.questionController.Get, nil)
	})
	registry.HandleDelete("quiz-question", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.questionController.Delete, []string{"administrator", "instructor"})
	})
	registry.HandlePut("quiz-question", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.questionController.Update, []string{"administrator", "instructor"})
	})

	registry.Route("/quiz-answer/", "quiz-answer")
	registry.HandlePost("quiz-answer", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.questionController.Create, []string{"student"})
	})

	registry.Route("/quiz-answer/:id", "quiz-answer")
	registry.HandleGet("quiz-answer", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.questionController.Get, nil)
	})
	registry.HandleDelete("quiz-answer", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.questionController.Delete, []string{"student"})
	})
	registry.HandlePut("quiz-answer", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.questionController.Update, []string{"student"})
	})

	registry.Route("/user/", "user")
	registry.HandlePost("user", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.userController.Create, []string{"administrator"})
	})

	registry.Route("/login/", "login")
	registry.HandlePost("login", routes.userController.Login)

	registry.Route("/refresh/", "refresh")
	registry.HandlePost("refresh", routes.userController.Refresh)

	registry.Route("/announcement-all/", "announcement-all")
	registry.HandleGet("announcement-all", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.announcementController.GetAll, nil)
	})

	registry.Route("/announcement/", "announcement")
	registry.HandlePost("announcement", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.announcementController.Create, []string{"administrator"})
	})

	registry.Route("/announcement/:id", "announcement")
	registry.HandlePut("announcement", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.announcementController.Update, []string{"administrator"})
	})
	registry.HandleDelete("announcement", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.announcementController.Delete, []string{"administrator"})
	})

	registry.Route("/student/:id", "student")
	registry.HandleGet("student", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.studentController.Get, nil)
	})
	registry.HandlePut("student", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.studentController.Update, nil)
	})
	registry.HandleDelete("student", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.studentController.Delete, []string{"administrator"})
	})
	registry.Route("/student/", "student")
	registry.HandlePost("student", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.studentController.Get, []string{"administrator", "student"})
	})

	registry.Route("/student-all/", "student-all")
	registry.HandleGet("student-all", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.studentController.GetAll, []string{"administrator"})
	})
}
