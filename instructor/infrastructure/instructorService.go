package infrastructure

import (
	"alfath_lms/api/definitions"
	"alfath_lms/instructor/domain/entity"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type InstructorService struct {
	db *gorm.DB
}

func (instructorSvc *InstructorService) Inject(db *gorm.DB) {
	instructorSvc.db = db
}

func (instructorSvc InstructorService) CreateInstructor(instructor entity.Instructor) (definitions.GenericCreationMessage, error) {

	var instructorTemp entity.Instructor
	instructorSvc.db.Where("Email = ?", instructor.Email).First(&instructorTemp)
	if instructorTemp != (entity.Instructor{}) {
		return definitions.GenericCreationMessage{}, errors.New("Data with that email already exists!")
	}

	instructorSvc.db.Where("Mobile_Phone = ?", instructor.MobilePhone).First(&instructorTemp)
	if instructorTemp != (entity.Instructor{}) {
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

func (instructorSvc InstructorService) GetInstructor(id string) (entity.Instructor, error) {
	var instructor entity.Instructor

	result := &instructor
	instructorSvc.db.First(result, "id = ?", id)

	return *result, nil

}

func PrintError(err error) error {
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	return nil
}
