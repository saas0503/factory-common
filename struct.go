package futils

import (
	"fmt"
	"reflect"
)

func StringInSlice(value string, slice []string) bool {
	for _, elem := range slice {
		if elem == value {
			return true
		}
	}
	return false
}

func MergeStructs(structs ...interface{}) reflect.Type {
	var structFields []reflect.StructField
	var structFieldNames []string

	for _, item := range structs {
		rt := reflect.TypeOf(item)
		for i := 0; i < rt.NumField(); i++ {
			field := rt.Field(i)
			if !StringInSlice(field.Name, structFieldNames) {
				structFields = append(structFields, field)
				structFieldNames = append(structFieldNames, field.Name)
			}
		}
	}

	return reflect.StructOf(structFields)
}

func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("no such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return fmt.Errorf("provide value type %s didn't match obj field type %s", val.Type(), structFieldType)
	}

	structFieldValue.Set(val)
	return nil
}
