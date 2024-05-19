package services

import (
	"alfath_lms/api/definitions"

	"gorm.io/gorm"
)

type DashboardService struct {
	db *gorm.DB
}

func (dashSvc *DashboardService) Inject(db *gorm.DB) {
	dashSvc.db = db
}

func (dashSvc *DashboardService) GetDashboardData() (definitions.SimpleDashboardData, error) {
	var result definitions.SimpleDashboardData
	var courses int64
	var students int64
	var submissions int64
	dashSvc.db.Raw("Select count(1) from ms_course").Scan(&courses)
	dashSvc.db.Raw("Select count(1) from ms_student").Scan(&students)
	dashSvc.db.Raw("select count(1) from (select distinct quiz_id, student_id from ms_student_quiz) xxx").Scan(&submissions)
	result.TotalCourses = courses
	result.TotalStudents = students
	result.TotalSubmissions = submissions

	return result, nil

}
