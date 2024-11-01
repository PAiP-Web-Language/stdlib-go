package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/PAiP-Web-Language/stdlib-go/internal"
)

// CompareEqual compares two values and if they are not equal it will Fail test
// This function is using comparison operator and uses typed parameters
func CompareEqual[T comparable](t testing.TB, expected, actual T, args ...any) {
	t.Helper()
	if expected != actual {
		ErrorFormat(t, 1, fmt.Sprintf("expected %v, got %v", expected, actual), args...)
	}
}

// Equal compares two values and if they are not equal it will Fail test
// This function is using DeepEqual function from reflect package
func Equal(t testing.TB, expected, actual any, args ...any) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		ErrorFormat(t, 1, fmt.Sprintf("expected %v, got %v", expected, actual), args...)
	}
}

// DeepEqual compares two values and if they are not equal it will Fail test
// This function is using internal.ReflectDeeperEqual function for checking even stricter equality
func DeepEqual(t testing.TB, expected, actual any, args ...any) {
	t.Helper()
	if !internal.ReflectDeeperEqual(expected, actual) {
		ErrorFormat(t, 1, fmt.Sprintf("expected %v, got %v", expected, actual), args...)
	}
}

// CompareNotEqual compares two values and if they are equal it will Fail test
// This function is using comparison operator and uses typed parameters
func CompareNotEqual[T comparable](t testing.TB, expected, actual T, args ...any) {
	t.Helper()
	if expected == actual {
		ErrorFormat(t, 1, fmt.Sprintf("expected %v, got %v", expected, actual), args...)
	}
}

// NotEqual compares two values and if they are equal it will Fail test
// This function is using DeepEqual function from reflect package
func NotEqual(t testing.TB, expected, actual any, args ...any) {
	t.Helper()
	if reflect.DeepEqual(expected, actual) {
		ErrorFormat(t, 1, fmt.Sprintf("expected %v, got %v", expected, actual), args...)
	}
}

// NotDeepEqual compares two values and if they are equal it will Fail test
// This function is using internal.ReflectDeeperEqual function for checking even stricter equality
func NotDeepEqual(t testing.TB, expected, actual any, args ...any) {
	t.Helper()
	if internal.ReflectDeeperEqual(expected, actual) {
		ErrorFormat(t, 1, fmt.Sprintf("expected %v, got %v", expected, actual), args...)
	}
}

// NotNil compares value and if it is nil it will Fail test
func NotNil(t testing.TB, value any, args ...any) {
	t.Helper()
	if value == nil {
		ErrorFormat(t, 1, fmt.Sprintf("expected %v to be not nil", value), args...)
	}
}

// Nil compares value and if it is not nil it will Fail test
func Nil(t testing.TB, value any, args ...any) {
	t.Helper()
	if value != nil {
		ErrorFormat(t, 1, fmt.Sprintf("expected %v to be nil", value), args...)
	}
}
