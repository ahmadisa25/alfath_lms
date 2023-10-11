package infrastructure

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/deps/pagination"
	"alfath_lms/api/student/domain/entity"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type StudentService struct {
	db        *gorm.DB
	paginator *pagination.Paginator
}

func (studentSvc *StudentService) Inject(db *gorm.DB, paginator *pagination.Paginator) {
	studentSvc.db = db
	studentSvc.paginator = paginator
}

func (studentSvc *StudentService) GetAllStudents(req definitions.PaginationRequest) (definitions.PaginationResult, error) {
	paginationParams := definitions.PaginationParam{
		Sql:          "Select -select- from (Select * from ms_student) as foo -where-",
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

	res := studentSvc.paginator.Paginate(req, paginationParams)
	return res, nil
}

func (studentSvc *StudentService) CreateStudent(student entity.Student) (definitions.GenericCreationMessage, error) {

	var studentTemp entity.Student
	studentSvc.db.Where("Email = ?", student.Email).First(&studentTemp)
	if studentTemp != (entity.Student{}) {
		return definitions.GenericCreationMessage{}, errors.New("Data with that email already exists!")
	}

	studentSvc.db.Where("Mobile_Phone = ?", student.MobilePhone).First(&studentTemp)
	if studentTemp != (entity.Student{}) {
		return definitions.GenericCreationMessage{}, errors.New("Data with that mobile phone already exists!")
	}

	result := studentSvc.db.Create(&student)
	if result.Error != nil {
		return definitions.GenericCreationMessage{}, result.Error
	}
	return definitions.GenericCreationMessage{
		Status:     201,
		InstanceID: student.ID,
	}, nil
}

func (studentSvc *StudentService) UpdateStudent(id int, student entity.Student) (definitions.GenericAPIMessage, error) {
	var studentTemp entity.Student
	result := studentSvc.db.Model(&studentTemp).Where("id = ?", id).Updates(&student)
	fmt.Println(*result)
	if result.Error != nil {
		return definitions.GenericAPIMessage{}, result.Error
	}
	return definitions.GenericAPIMessage{
		Status:  200,
		Message: "Student is successfully updated",
	}, nil
}

func (studentSvc *StudentService) DeleteStudent(id int) (definitions.GenericAPIMessage, error) {
	//delete here means soft delete
	result := studentSvc.db.Where("id = ?", id).Delete(&entity.Student{})
	if result.Error != nil {
		return definitions.GenericAPIMessage{}, result.Error
	}
	return definitions.GenericAPIMessage{
		Status:  200,
		Message: "Student has been deleted successfully",
	}, nil
}

func (studentSvc *StudentService) GetStudent(id int) (entity.Student, error) {
	var student entity.Student

	result := &student
	studentSvc.db.First(result, "id = ?", id)

	return *result, nil

}

func PrintError(err error) error {
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	return nil
}
