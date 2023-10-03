package funcs

import (
	"alfath_lms/api/definitions"
	"encoding/json"
)

func ErrorPackagingForMaps(errs []error) string {
	errorStrings := ""

	for i, err := range errs {
		errorStrings += err.Error()
		if i < len(errs)-1 {
			errorStrings += ","
		}
	}

	return errorStrings

}

func ErrorPackaging(err string, status int) (string, error) {
	res, resErr := json.Marshal(definitions.GenericAPIMessage{
		Status:  status,
		Message: err,
	})

	if resErr != nil {
		return "", resErr
	}

	return string(res), nil
}

func ValidateStringFormKeys(mapKey string, form map[string][]string, dataType string) interface{} {
	// map[dataType]interface{} means that the map has key of dataTypes and value of any type (yes the interface{} there is a powerful syntax.)
	//used form Flamingo Form Requests (r.Request().Form)
	key, keyOk := form[mapKey]
	if !keyOk {
		if dataType == "string" {
			return ""
		} else if dataType == "int" || dataType == "uint" {
			return 0
		} else if dataType == "bool" {
			return false
		}

		return nil
	}
	//test15
	return key[0]
}

func ValidateOrOverwriteStringFormKeys(mapKey string, form map[string][]string, dataType string, input interface{}) interface{} {
	// map[dataType]interface{} means that the map has key of dataTypes and value of any type (yes the interface{} there is a powerful syntax.)
	//used form Flamingo Form Requests (r.Request().Form)

	if _, isNotStruct := input.(struct{}); isNotStruct{
		return nil
	} else {
		key, keyOk := form[mapKey]
		if !keyOk {
			//ganti supaya ngakses fieldnya si input
			return ""
		}
		//test15
		return key[0]
	}
	
	
}
