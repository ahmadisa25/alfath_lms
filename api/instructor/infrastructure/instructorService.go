package infrastructure

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/instructor/domain/entity"
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

func (instructorSvc *InstructorService) CreateInstructor(instructor entity.Instructor) (definitions.GenericCreationMessage, error) {

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

func (instructorSvc *InstructorService) UpdateInstructor(id int, instructor entity.Instructor) (definitions.GenericAPIMessage, error) {
	var instructorTemp entity.Instructor
	result := instructorSvc.db.Model(&instructorTemp).Where("id = ?", id).Updates(&instructor)
	fmt.Println(*result)
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
	result := instructorSvc.db.Where("id = ?", id).Delete(&entity.Instructor{})
	if result.Error != nil {
		return definitions.GenericAPIMessage{}, result.Error
	}
	return definitions.GenericAPIMessage{
		Status:  200,
		Message: "Instructor has been deleted successfully",
	}, nil
}

func (instructorSvc *InstructorService) GetInstructor(id int) (entity.Instructor, error) {
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
