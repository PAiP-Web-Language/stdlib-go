package test

import (
	"runtime"
	"testing"

	"github.com/PAiP-Web-Language/stdlib-go/dev"
)

// ErrorFormat is a function that will fail test with specified message
// It will add some information about location of error
func ErrorFormat(t testing.TB, skip int, msg string, args ...any) {
	t.Helper()
	_, file, line, _ := runtime.Caller(skip + 1)
	if TestConfig_FailFast {
		t.Fatalf("%s:%d %s %v", file, line, msg, args)
	} else {
		t.Errorf("%s:%d %s %v", file, line, msg, args)
	}
}

// TestUtilErrorFormat is a function that will fully fail test with specified message
// This function is intended for testing helpers that should force fail test (not only error)
// It's intended to bu used in code paths that happen if there is logical error in test itself
func TestUtilErrorFormat(t testing.TB, skip int, msg string, args ...any) {
	t.Helper()
	_, file, line, _ := runtime.Caller(skip + 1)
	t.Fatalf("%s:%d Test Error: %s %v", file, line, msg, args)

	dev.Unreachable("Test Logic Error")
}
