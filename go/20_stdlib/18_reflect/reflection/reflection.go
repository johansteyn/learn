package reflection

import (
	"fmt"
	"reflect"
)

func GetFieldByName[T any](s *T, fieldName string) any {
	if s == nil || fieldName == "" {
		panic(fmt.Sprintf("nil or empty fieldName for type %T", s))
	}
	val := reflect.ValueOf(s).Elem()
	if val.Kind() != reflect.Struct || !val.IsValid() {
		panic(fmt.Sprintf("not a struct or invalid value for type %T", s))
	}
	field := val.FieldByName(fieldName)
	if !field.IsValid() {
		panic(fmt.Sprintf("field %s is not valid in type %T", fieldName, s))
	}
	return field.Interface().(any)
}

func SetFieldByName[T any](s *T, fieldName string, value any) {
	if s == nil || fieldName == "" {
		panic(fmt.Sprintf("nil or empty fieldName for type %T", s))
	}
	val := reflect.ValueOf(s).Elem()
	if val.Kind() != reflect.Struct || !val.IsValid() {
		panic(fmt.Sprintf("not a struct or invalid value for type %T", s))
	}
	field := val.FieldByName(fieldName)
	if !field.IsValid() {
		panic(fmt.Sprintf("field %s is not valid in type %T", fieldName, s))
	}
	field.Set(reflect.ValueOf(value))
}
