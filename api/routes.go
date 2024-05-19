package api

import (
	"alfath_lms/api/controllers"
	"alfath_lms/api/middleware"
	"context"

	"flamingo.me/flamingo/v3/framework/web"
)

type Routes struct {
	questionController      controllers.QuestionController
	quizController          controllers.QuizController
	answerController        controllers.AnswerController
	materialController      controllers.MaterialController
	chapterController       controllers.ChapterController
	courseController        controllers.CourseController
	instructorController    controllers.InstructorController
	studentController       controllers.StudentController
	userController          controllers.UserController
	announcementController  controllers.AnnouncementController
	studentQuizController   controllers.StudentQuizController
	studentCourseController controllers.StudentCourseController
	dashboardController     controllers.DashboardController
	optionsHandler          controllers.OptionsHandler
	uploadHandler           controllers.UploadHandler
	authMdw                 *middleware.AuthMiddleware
	responder               *web.Responder
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
	optionsHandler *controllers.OptionsHandler,
	uploadHandler *controllers.UploadHandler,
	announcementController *controllers.AnnouncementController,
	studentQuizController *controllers.StudentQuizController,
	studentCourseController *controllers.StudentCourseController,
	dashboardController *controllers.DashboardController,
	responder *web.Responder,
) {
	routes.answerController = *answerController
	routes.optionsHandler = *optionsHandler
	routes.uploadHandler = *uploadHandler
	routes.quizController = *quizController
	routes.materialController = *materialController
	routes.courseController = *courseController
	routes.chapterController = *chapterController
	routes.instructorController = *instructorController
	routes.studentController = *studentController
	routes.questionController = *questionController
	routes.userController = *userController
	routes.announcementController = *announcementController
	routes.studentQuizController = *studentQuizController
	routes.studentCourseController = *studentCourseController
	routes.dashboardController = *dashboardController
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

	registry.HandleOptions("instructor", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.optionsHandler.Setup, nil)
	})

	registry.Route("/instructor-all/", "instructor-all")
	registry.HandleGet("instructor-all", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.instructorController.GetAll, nil)
	})

	registry.HandleOptions("instructor-all", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.optionsHandler.Setup, nil)
	})

	registry.Route("/course-all/", "course-all")
	registry.HandleGet("course-all", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.courseController.GetAll, nil)
	})

	registry.HandleOptions("course-all", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.optionsHandler.Setup, nil)
	})

	registry.HandleOptions("course", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.optionsHandler.Setup, nil)
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

	registry.Route("/course-chapter/", "course-chapter")
	registry.HandleOptions("course-chapter", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.optionsHandler.Setup, []string{})
	})

	registry.Route("/course-chapter/:id", "course-chapter")
	registry.HandleGet("course-chapter", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.chapterController.Get, nil)
	})

	registry.Route("/course-chapter/:id", "course-chapter")
	registry.HandleOptions("course-chapter", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.optionsHandler.Setup, []string{})
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

	registry.Route("/chapter-material/", "course-chapter")
	registry.HandleOptions("course-chapter", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.optionsHandler.Setup, []string{})
	})

	registry.Route("/chapter-material/:id", "chapter-material")
	registry.HandleOptions("chapter-material", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.optionsHandler.Setup, []string{})
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

	registry.Route("/student-quiz/:student_id", "student-quiz")
	registry.HandleGet("student-quiz", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.studentQuizController.Get, []string{"administrator", "instructor"})
	})

	registry.Route("/student-quiz/:id", "student-quiz")
	registry.HandleOptions("student-quiz", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.optionsHandler.Setup, []string{})
	})

	registry.Route("/student-quiz/", "student-quiz")
	registry.HandlePost("student-quiz", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.studentQuizController.Create, []string{"administrator", "instructor"})
	})

	registry.Route("/student-quiz/", "student-quiz")
	registry.HandleOptions("student-quiz", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.optionsHandler.Setup, []string{})
	})

	registry.Route("/student-course/", "student-course")
	registry.HandlePost("student-course", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.studentCourseController.Create, []string{"administrator", "student"})
	})

	registry.Route("/student-course/", "student-course")
	registry.HandleOptions("student-course", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.optionsHandler.Setup, []string{})
	})

	registry.Route("/student-course/:id", "student-course")
	registry.HandleDelete("student-course", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.studentCourseController.Delete, []string{"administrator", "student"})
	})

	registry.Route("/student-course/:id", "student-course")
	registry.HandleOptions("student-course", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.optionsHandler.Setup, []string{})
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

	registry.Route("/chapter-quiz/", "chapter-quiz")
	registry.HandleOptions("chapter-quiz", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.optionsHandler.Setup, []string{})
	})

	registry.Route("/chapter-quiz/:id", "chapter-quiz")
	registry.HandleOptions("chapter-quiz", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.optionsHandler.Setup, []string{})
	})

	registry.Route("/quiz-question/", "quiz-question")
	registry.HandlePost("quiz-question", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.questionController.Create, []string{"administrator", "instructor"})
	})

	registry.Route("/quiz-question/:id", "quiz-question")
	registry.HandleOptions("quiz-question", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.optionsHandler.Setup, []string{})
	})
	registry.Route("/quiz-question/", "quiz-question")
	registry.HandleOptions("quiz-question", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.optionsHandler.Setup, []string{})
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

	registry.Route("/quiz-answer-all-distinct/", "quiz-answer-all-distinct")
	registry.HandleGet("quiz-answer-all-distinct", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.answerController.GetAllDistinct, []string{"administrator", "instructor"})
	})

	registry.Route("/quiz-answer-all-distinct/", "quiz-answer-all-distinct")
	registry.HandleOptions("quiz-answer-all-distinct", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.optionsHandler.Setup, nil)
	})

	registry.Route("/quiz-answer-all/", "quiz-answer-all")
	registry.HandleGet("quiz-answer-all", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.answerController.GetAll, []string{"administrator", "instructor"})
	})

	registry.Route("/quiz-answer-all/", "quiz-answer-all")
	registry.HandleOptions("quiz-answer-all", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.optionsHandler.Setup, nil)
	})

	registry.Route("/quiz-answer/", "quiz-answer")
	registry.HandlePost("quiz-answer", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.answerController.Create, []string{"administrator", "student"})
	})

	registry.Route("/quiz-answer/", "quiz-answer")
	registry.HandleOptions("quiz-answer", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.optionsHandler.Setup, []string{})
	})

	registry.Route("/quiz-answer/:id", "quiz-answer")
	registry.HandleGet("quiz-answer", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.answerController.Get, nil)
	})
	registry.HandleDelete("quiz-answer", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.answerController.Delete, []string{"student"})
	})
	registry.HandlePut("quiz-answer", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.answerController.Update, []string{"student"})
	})

	registry.Route("/register/", "register")
	registry.HandlePost("register", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.userController.Create, []string{"administrator"})
	})

	registry.Route("/login-admin/", "login-admin")
	registry.HandlePost("login-admin", routes.userController.LoginAdmin)

	registry.Route("/login-admin/", "login-admin")
	registry.HandleOptions("login-admin", routes.optionsHandler.Setup)

	registry.Route("/login/", "login")
	registry.HandlePost("login", routes.userController.Login)

	registry.Route("/login/", "login")
	registry.HandleOptions("login", routes.optionsHandler.Setup)

	registry.Route("/refresh/", "refresh")
	registry.HandlePost("refresh", routes.userController.Refresh)

	registry.Route("/refresh/", "refresh")
	registry.HandleOptions("refresh", routes.optionsHandler.Setup)

	registry.Route("/announcement-all/", "announcement-all")
	registry.HandleGet("announcement-all", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.announcementController.GetAll, nil)
	})

	registry.Route("/uploads/:file_name", "upload_dir")
	registry.HandleGet("upload_dir", func(ctx context.Context, req *web.Request) web.Result {
		return routes.uploadHandler.Setup(ctx, req)
	})

	registry.HandleOptions("upload_dir", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.optionsHandler.Setup, nil)
	})

	registry.HandleOptions("announcement-all", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.optionsHandler.Setup, nil)
	})

	registry.Route("/announcement/", "announcement")
	registry.HandlePost("announcement", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.announcementController.Create, []string{"administrator"})
	})

	registry.HandleGet("announcement", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.announcementController.Get, nil)
	})

	registry.HandleOptions("announcement", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.optionsHandler.Setup, nil)
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
		return routes.authMdw.AuthCheck(ctx, req, routes.studentController.Create, []string{"administrator", "student"})
	})

	registry.HandleOptions("student", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.optionsHandler.Setup, nil)
	})

	registry.Route("/student-all/", "student-all")
	registry.HandleGet("student-all", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.studentController.GetAll, []string{"administrator"})
	})

	registry.HandleOptions("student-all", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.optionsHandler.Setup, nil)
	})

	registry.Route("/dashboard", "dashboard")
	registry.HandleOptions("dashboard", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.optionsHandler.Setup, []string{})
	})
	registry.HandleGet("dashboard", func(ctx context.Context, req *web.Request) web.Result {
		return routes.authMdw.AuthCheck(ctx, req, routes.dashboardController.GetDashboardData, nil)
	})
}
