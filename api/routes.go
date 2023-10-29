package api

import (
	"alfath_lms/api/controllers"

	"flamingo.me/flamingo/v3/framework/web"
)

type Routes struct {
	chapterController    controllers.ChapterController
	courseController     controllers.CourseController
	instructorController controllers.InstructorController
	studentController    controllers.StudentController
}

func (routes *Routes) Inject(
	chapterController *controllers.ChapterController,
	courseController *controllers.CourseController,
	instructorController *controllers.InstructorController,
	studentController *controllers.StudentController,
) {
	routes.courseController = *courseController
	routes.chapterController = *chapterController
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

	registry.Route("/student/:id", "student")
	registry.HandleGet("student", routes.studentController.Get)
	registry.HandlePut("student", routes.studentController.Update)
	registry.HandleDelete("student", routes.studentController.Delete)

	registry.Route("/student/", "student")
	registry.HandlePost("student", routes.studentController.Create)

	registry.Route("/student-all/", "student-all")
	registry.HandleGet("student-all", routes.studentController.GetAll)
}
