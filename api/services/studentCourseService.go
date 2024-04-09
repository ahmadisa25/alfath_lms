package services

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/deps/pagination"
	"alfath_lms/api/models"

	"gorm.io/gorm"
)

type StudentCourseService struct {
	db        *gorm.DB
	paginator *pagination.Paginator
}

func (studCourseSvc *StudentCourseService) Inject(db *gorm.DB, paginator *pagination.Paginator) {
	studCourseSvc.db = db
	studCourseSvc.paginator = paginator
}

func (studCourseSvc *StudentCourseService) Create(stdCourse models.StudentCourse) (definitions.GenericCreationMessage, error) {
	result := studCourseSvc.db.Create(&stdCourse)
	if result.Error != nil {
		return definitions.GenericCreationMessage{}, result.Error
	}
	return definitions.GenericCreationMessage{
		Status:     201,
		InstanceID: stdCourse.ID,
	}, nil
}

func (studCourseSvc *StudentCourseService) Delete(id int) (definitions.GenericAPIMessage, error) {
	//delete here means soft delete
	result := studCourseSvc.db.Where("id = ?", id).Delete(&models.StudentCourse{})
	if result.Error != nil {
		return definitions.GenericAPIMessage{}, result.Error
	}
	return definitions.GenericAPIMessage{
		Status:  200,
		Message: "Enrollment has been deleted successfully",
	}, nil
}
