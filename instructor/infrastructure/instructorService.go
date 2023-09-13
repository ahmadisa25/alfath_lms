package infrastructure

import (
	"alfath_lms/instructor/domain/entity"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type InstructorService struct{}

func (instructorSvc InstructorService) GetOrder(id int) (entity.Instructor, error) {
	var instructor entity.Instructor
	response, err := http.Get(fmt.Sprintf("instructor-service/instructors/%d", id))
	PrintError(err)

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)

	PrintError(err)

	if err := json.Unmarshal(responseBody, &instructor); err != nil {
		PrintError(err)
	}

	return instructor, nil

}

func PrintError(err error) error {
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	return nil
}
