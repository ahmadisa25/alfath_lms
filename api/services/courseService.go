package services

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/deps/pagination"
	"alfath_lms/api/models"
	"errors"
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
		SelectFields: []string{"name", "description", "duration"},
		SearchFields: map[string]string{
			"name": "foo.name",
		},
		FilterFields: map[string]string{
			"name":        "foo.name",
			"description": "foo.description",
			"deleted_at":  "foo.deleted_at",
		},
		NullFilterFields: map[string]bool{
			"deleted_at": true,
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
	var instructors []*models.Instructor
	instructorIDs := []int{}
	for _, val := range strings.Split(instructorList, ",") {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			return definitions.GenericAPIMessage{}, err
		}
		instructorIDs = append(instructorIDs, intVal)
	}
	courseSvc.db.Table("ms_instructor").Where("id IN (?)", instructorIDs).Find(&instructors)
	if len(instructors) == 0 {
		return definitions.GenericAPIMessage{}, errors.New("Instructors don't exist")
	}

	//fmt.Println(course.Instructors)
	course.Instructors = instructors
	course.ID = id

	//result := courseSvc.db.Updates(&course)
	if len(instructors) > 0 {
		instructorDelete := courseSvc.db.Table("ms_course_instructor").Where("course_id = ?", id).Unscoped().Delete(&models.Course{})
		if instructorDelete.Error != nil {
			return definitions.GenericAPIMessage{}, instructorDelete.Error
		}
	}

	result := courseSvc.db.Updates(&course)
	if result.Error != nil {
		return definitions.GenericAPIMessage{}, result.Error
	}
	return definitions.GenericAPIMessage{
		Status:  200,
		Message: "Course has been successfully updated",
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
