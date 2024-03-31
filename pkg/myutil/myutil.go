package myutil

import (
	"encoding/json"
	"reflect"
)

func MapToJson(param map[string]interface{}) string {
	dataType, err := json.Marshal(param)
	if err != nil {
		panic(err)
	}
	dataString := string(dataType)
	return dataString
}

func JsonToMap(str string) map[string]interface{} {
	var tempMap map[string]interface{}
	err := json.Unmarshal([]byte(str), &tempMap)
	if err != nil {
		panic(err)
	}
	return tempMap
}

func IsMember[T int | float64 | string, Slice1 []T](elem T, elems Slice1) bool {
	for _, v := range elems {
		if v == elem {
			return true
		}
	}
	return false
}

func StructToMap(obj interface{}) map[string]interface{} {
	objValue := reflect.ValueOf(obj)
	if objValue.Kind() != reflect.Struct {
		return nil
	}

	objType := objValue.Type()
	result := make(map[string]interface{})

	for i := 0; i < objValue.NumField(); i++ {
		fieldName := objType.Field(i).Name
		fieldValue := objValue.Field(i).Interface()
		result[fieldName] = fieldValue
	}

	return result
}
