package infrastructure

import (
	"alfath_lms/instructor/domain/entity"
	"fmt"
	"gorm.io/gorm"
	"alfath_lms/api/definitions"
)

type InstructorService struct{
	db *gorm.DB
}

func (instructorSvc *InstructorService) Inject(db *gorm.DB){
	instructorSvc.db = db
}

func (instructorSvc InstructorService) 	CreateInstructor(instructor entity.Instructor) (definitions.GenericCreationMessage, error){
	result := instructorSvc.db.Create(&instructor)
	if result.Error != nil{
		return definitions.GenericCreationMessage{}, result.Error
	} 

	return definitions.GenericCreationMessage{
		Status: 201,
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
