package test

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

// SetPrivateMember sets new a value in to the struct private field.
// This function is intended to be used for testing only
func SetPrivateMember(t testing.TB, obj interface{}, fieldName string, value interface{}) {
	reflectValue := reflect.ValueOf(obj)

	if reflectValue.Kind() != reflect.Ptr {
		TestUtilErrorFormat(t, 1, "object is not a pointer")
	}

	reflectValue = reflectValue.Elem()
	reflectType := reflect.TypeOf(obj).Elem()

	if reflectValue.Kind() != reflect.Struct {
		TestUtilErrorFormat(t, 1, "object is not a pointer to struct")
	}

	newReflectValueType := reflect.TypeOf(value)

	if fieldType, ok := reflectType.FieldByName(fieldName); ok {
		currentFieldType := fieldType.Type.String()

		if newReflectValueType.String() != currentFieldType {
			TestUtilErrorFormat(t, 1, fmt.Sprintf("incorrect type: must be %s", currentFieldType))
		}

		fieldPtr := unsafe.Pointer(reflectValue.UnsafeAddr() + fieldType.Offset)
		reflect.NewAt(newReflectValueType, fieldPtr).Elem().Set(reflect.ValueOf(value))

		return
	}

	TestUtilErrorFormat(t, 1, fmt.Sprintf("Field [%s] not found in struct", fieldName))
}
