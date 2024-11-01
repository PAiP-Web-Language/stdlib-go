package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/PAiP-Web-Language/stdlib-go/internal"
)

func Panic(t testing.TB, f func(), args ...any) {
	if !internal.DoFPanics(f) {
		ErrorFormat(t, 1, "function did not panic", args...)
	}
}

func NotPanic(t testing.TB, f func(), args ...any) {
    if internal.DoFPanics(f) {
        ErrorFormat(t, 1, "function panicked", args...)
    }
}

func PanicMatches(t testing.TB, f func(), err any, args ...any) {
	check, e := internal.DoFPanicsWithErr(f)
	if !check {
		ErrorFormat(t, 1, "function did not panic", args...)
	}
	if !reflect.DeepEqual(err, e) {
		ErrorFormat(t, 1, fmt.Sprintf("function panicked with %v, expected %v", e, err), args...)
	}
}

func PanicNotMatches(t testing.TB, f func(), err any, args ...any) {
	check, e := internal.DoFPanicsWithErr(f)
	if !check {
		ErrorFormat(t, 1, "function did not panic", args...)
	}
	if reflect.DeepEqual(err, e) {
		ErrorFormat(t, 1, fmt.Sprintf("function panicked with %v, expected to not panic with %v", e, err), args...)
	}
}
