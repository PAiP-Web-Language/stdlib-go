package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/PAiP-Web-Language/stdlib-go/internal"
)

// Panic is a function that checks if function panics if it does not panic it will Fail test
func Panic(t testing.TB, f func(), args ...any) {
	t.Helper()
	if !internal.DoFPanics(f) {
		ErrorFormat(t, 1, "function did not panic", args...)
	}
}

// NotPanic is a function that checks if function panics if it does panic it will Fail test
func NotPanic(t testing.TB, f func(), args ...any) {
	t.Helper()
    if internal.DoFPanics(f) {
        ErrorFormat(t, 1, "function panicked", args...)
    }
}

// PanicMatches is a function that checks if function panics with specific error
// if it does not panic it will Fail test or if it panics with different value
func PanicMatches(t testing.TB, f func(), err any, args ...any) {
	t.Helper()
	check, e := internal.DoFPanicsWithErr(f)
	if !check {
		ErrorFormat(t, 1, "function did not panic", args...)
	}
	if !reflect.DeepEqual(err, e) {
		ErrorFormat(t, 1, fmt.Sprintf("function panicked with %v, expected %v", e, err), args...)
	}
}

// NotPanicMatches is a function that checks if function panics with any error other than specified
// if it does panic it will Fail test or if it panics with same value
func PanicNotMatches(t testing.TB, f func(), err any, args ...any) {
	t.Helper()
	check, e := internal.DoFPanicsWithErr(f)
	if !check {
		ErrorFormat(t, 1, "function did not panic", args...)
	}
	if reflect.DeepEqual(err, e) {
		ErrorFormat(t, 1, fmt.Sprintf("function panicked with %v, expected to not panic with %v", e, err), args...)
	}
}
