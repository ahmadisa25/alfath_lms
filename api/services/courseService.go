package services

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/deps/pagination"
	"alfath_lms/api/models"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type CourseService struct {
	db        *gorm.DB
	paginator *pagination.Paginator
}

func (courseSvc *CourseService) Inject(db *gorm.DB, paginator *pagination.Paginator) {
	courseSvc.db = db
	courseSvc.paginator = paginator
}

func (courseSvc *CourseService) GetAll(req definitions.PaginationRequest) (definitions.PaginationResult, error) {
	paginationParams := definitions.PaginationParam{
		Sql:          "Select -select- from (Select * from ms_course) as foo -where-",
		SelectFields: []string{"name", "email", "mobile_phone"},
		SearchFields: map[string]string{
			"name":         "foo.name",
			"email":        "foo.email",
			"mobile_phone": "foo.mobile_phone",
		},
		FilterFields: map[string]string{
			"name":         "foo.name",
			"email":        "foo.email",
			"mobile_phone": "foo.mobile_phone",
		},
	}

	res := courseSvc.paginator.Paginate(req, paginationParams)
	return res, nil
}

func (courseSvc *CourseService) Create(course models.Course, instructorList string) (definitions.GenericCreationMessage, error) {

	var instructors []*models.Instructor
	instructorIDs := []int{}
	for _, val := range strings.Split(instructorList, ",") {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			return definitions.GenericCreationMessage{}, err
		}
		instructorIDs = append(instructorIDs, intVal)
	}
	courseSvc.db.Table("ms_instructor").Where("id IN (?)", instructorIDs).Find(&instructors)
	if len(instructors) == 0 {
		return definitions.GenericCreationMessage{}, errors.New("Instructors don't exist")
	}

	course.Instructors = instructors

	result := courseSvc.db.Create(&course)
	if result.Error != nil {
		return definitions.GenericCreationMessage{}, result.Error
	}
	return definitions.GenericCreationMessage{
		Status:     201,
		InstanceID: course.ID,
	}, nil
}

func (courseSvc *CourseService) Update(id int, course models.Course, instructorList string) (definitions.GenericAPIMessage, error) {
	result := courseSvc.db.Model(&course).Where("id = ?", id).Updates(&course)
	fmt.Println(*result)
	if result.Error != nil {
		return definitions.GenericAPIMessage{}, result.Error
	}
	return definitions.GenericAPIMessage{
		Status:  200,
		Message: "Course is successfully updated",
	}, nil
}

func (courseSvc *CourseService) Delete(id int) (definitions.GenericAPIMessage, error) {
	//delete here means soft delete
	result := courseSvc.db.Where("id = ?", id).Delete(&models.Course{})
	if result.Error != nil {
		return definitions.GenericAPIMessage{}, result.Error
	}
	return definitions.GenericAPIMessage{
		Status:  200,
		Message: "Course has been deleted successfully",
	}, nil
}

func (courseSvc *CourseService) Get(id int) (models.Course, error) {
	var course models.Course

	result := &course
	courseSvc.db.Preload("Instructors").First(result, "id = ?", id)

	return *result, nil

}
