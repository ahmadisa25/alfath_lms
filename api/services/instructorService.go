package services

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/deps/pagination"
	"alfath_lms/api/models"
	"errors"

	"gorm.io/gorm"
)

type InstructorService struct {
	db        *gorm.DB
	paginator *pagination.Paginator
}

func (instructorSvc *InstructorService) Inject(db *gorm.DB, paginator *pagination.Paginator) {
	instructorSvc.db = db
	instructorSvc.paginator = paginator
}

func (instructorSvc *InstructorService) GetAllInstructors(req definitions.PaginationRequest) (definitions.PaginationResult, error) {
	paginationParams := definitions.PaginationParam{
		Sql:          "Select -select- from (Select * from ms_instructor) as foo -where-",
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
			"deleted_at":   "foo.deleted_at",
		},
		NullFilterFields: map[string]bool{
			"deleted_at": true,
		},
	}

	res := instructorSvc.paginator.Paginate(req, paginationParams)
	return res, nil
}

func (instructorSvc *InstructorService) CreateInstructor(instructor models.Instructor) (definitions.GenericCreationMessage, error) {

	var instructorTemp models.Instructor
	instructorSvc.db.Where("Email = ?", instructor.Email).First(&instructorTemp)
	if instructorTemp.Email != "" {
		return definitions.GenericCreationMessage{}, errors.New("Data with that email already exists!")
	}

	instructorSvc.db.Where("Mobile_Phone = ?", instructor.MobilePhone).First(&instructorTemp)
	if instructorTemp.MobilePhone != "" {
		return definitions.GenericCreationMessage{}, errors.New("Data with that mobile phone already exists!")
	}

	result := instructorSvc.db.Create(&instructor)
	if result.Error != nil {
		return definitions.GenericCreationMessage{}, result.Error
	}
	return definitions.GenericCreationMessage{
		Status:     201,
		InstanceID: instructor.ID,
	}, nil
}

func (instructorSvc *InstructorService) UpdateInstructor(id int, instructor models.Instructor, existingInstructor models.Instructor) (definitions.GenericAPIMessage, error) {
	var instructorTemp models.Instructor
	/*instructorSvc.db.Where("Email = ?", instructor.Email).First(&instructorTemp)
	if instructorTemp.Email != "" && existingInstructor.Email != instructor.Email {
		return definitions.GenericAPIMessage{
			Status:  400,
			Message: "Instructor with that email already exists!",
		}, nil
	}*/

	instructorSvc.db.Where("Mobile_Phone = ?", instructor.MobilePhone).First(&instructorTemp)
	if instructorTemp.MobilePhone != "" && existingInstructor.Email != instructor.Email {
		return definitions.GenericAPIMessage{
			Status:  400,
			Message: "Instructor with that mobile phone already exists!",
		}, nil
	}
	result := instructorSvc.db.Model(&instructorTemp).Where("id = ?", id).Updates(&instructor)

	if result.Error != nil {
		return definitions.GenericAPIMessage{}, result.Error
	}

	return definitions.GenericAPIMessage{
		Status:  200,
		Message: "Instructor is successfully updated",
	}, nil
}

func (instructorSvc *InstructorService) DeleteInstructor(id int) (definitions.GenericAPIMessage, error) {
	//delete here means soft delete
	result := instructorSvc.db.Where("id = ?", id).Delete(&models.Instructor{})
	if result.Error != nil {
		return definitions.GenericAPIMessage{}, result.Error
	}
	return definitions.GenericAPIMessage{
		Status:  200,
		Message: "Instructor has been deleted successfully",
	}, nil
}

func (instructorSvc *InstructorService) GetInstructor(id int) (models.Instructor, error) {
	var instructor models.Instructor

	result := &instructor
	instructorSvc.db.First(result, "id = ?", id)

	return *result, nil

}
