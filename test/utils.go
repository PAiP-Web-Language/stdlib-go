package test

import (
	"runtime"
	"testing"
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
