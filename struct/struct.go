package gostruct

import (
	"fmt"
	"reflect"
)

// HasField struct has field
func HasField(s interface{}, fieldName string) (bool, error) {

	// reflect.Value
	sValue := reflect.ValueOf(s)

	// is pointer
	if sValue.Kind() == reflect.Ptr {
		sValue = reflect.ValueOf(sValue.Elem().Interface())
	}

	if sValue.Kind() != reflect.Struct {
		return false, fmt.Errorf("bad request parameters : s isnot struct")
	}
	return sValue.FieldByName(fieldName).IsValid(), nil
}
