package infrastructure

import (
	"alfath_lms/instructor/domain/entity"
	//"encoding/json"
	"fmt"
	//"io/ioutil"
	//"net/http"
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
	/*response, err := http.Get(fmt.Sprintf("instructor-service/instructors/%d", id))
	PrintError(err)

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)

	PrintError(err)

	if err := json.Unmarshal(responseBody, &instructor); err != nil {
		PrintError(err)
	}*/

	result := &instructor
	instructorSvc.db.First(result, "id = ?", 1)

	fmt.Printf("read ID: %d, Code: %s",
    result.ID, result.Name)

	return *result, nil

}

func PrintError(err error) error {
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	return nil
}
