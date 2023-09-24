package infrastructure

import (
	"alfath_lms/instructor/domain/entity"
	"fmt"
	"gorm.io/gorm"
)

type InstructorService struct{
	db *gorm.DB
}

func (instructorSvc *InstructorService) Inject(db *gorm.DB){
	instructorSvc.db = db
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
