package entity

import (
	"reflect"
)

type Movie struct {
	ID     string
	Name   string
	Link   string
	Rating int
}

func (m *Movie) FromMap(result map[string]interface{}) {
	// Use reflection to iterate over struct fields and map keys
	elem := reflect.ValueOf(m).Elem()
	typeOfMovie := elem.Type()

	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		key := typeOfMovie.Field(i).Tag.Get("sql")
		switch field.Kind() {
		case reflect.String:
			if strValue, ok := result[key].(string); ok {
				field.SetString(strValue)
			}
		case reflect.Int:
			if intValue, ok := result[key].(int); ok {
				field.SetInt(int64(intValue))
			}
		default:
		}
	}
}
