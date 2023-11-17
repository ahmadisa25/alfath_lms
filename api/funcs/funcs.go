package funcs

import (
	"alfath_lms/api/definitions"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"reflect"
	"strconv"
)

func IsStringArrayItemExists(needle string, haystack []string) bool {

	for _, element := range haystack {
		if needle == element {
			return true
		}
	}

	return false
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

func GetStructField(value interface{}, fieldName string) interface{} {
	// used to validate if a struct by the name value has a field named fieldName
	ref := reflect.ValueOf(value)

	if ref.Kind() != reflect.Struct {
		return nil
	}

	field := ref.FieldByName(fieldName)
	if !field.IsValid() {
		return nil
	}

	return field.Interface()
}

func HashStringToSHA256(input string) string {
	if input == "" {
		return ""
	} else {
		hasher := sha256.New()
		hasher.Write([]byte(input))
		hashSum := hasher.Sum(nil)
		return hex.EncodeToString(hashSum)
	}
}

func ValidateStringFormKeys(mapKey string, form map[string][]string, dataType string) interface{} {
	// map[dataType]interface{} means that the map has key of dataTypes and value of any type (yes the interface{} there is a powerful syntax.)
	//used form Flamingo Form Requests (r.Request().Form) or Queries (r.QueryAll())

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
	if dataType == "int" {
		res, err := strconv.Atoi(key[0])
		if err != nil {
			return err
		}
		return res
	}
	return key[0]
}

func ValidateStringParamKeys(mapKey string, form map[string]string, dataType string) interface{} {
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

	return key[0]
}

func ValidateOrOverwriteStringFormKeys(mapKey string, form map[string][]string, dataType string, input interface{}) interface{} {
	//used form Flamingo Form Requests (r.Request().Form)

	if _, isNotStruct := input.(struct{}); isNotStruct {
		return nil
	} else {
		key, keyOk := form[mapKey]
		if !keyOk {
			//ganti supaya ngakses fieldnya si input
			return GetStructField(input, mapKey)
		}
		//test15
		return key[0]
	}

}
