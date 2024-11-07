package test

import (
	"runtime"
	"testing"

	"github.com/PAiP-Web-Language/stdlib-go/dev"
	"github.com/PAiP-Web-Language/stdlib-go/test/internal"
)

// ErrorFormat is a function that will fail test with specified message
// It will add some information about location of error
func ErrorFormat(t testing.TB, skip int, msg string, args ...any) {
	t.Helper()
	_, file, line, _ := runtime.Caller(skip + 1)
	if TestConfig_AssertFailFast {
		t.Fatalf("%s:%d %s %v", file, line, msg, args)
	} else {
		t.Errorf("%s:%d %s %v", file, line, msg, args)
	}
}

// FatalFormat is a function that will fail test with specified message
// It will add some information about location of error
func FatalFormat(t testing.TB, skip int, msg string, args ...any) {
	t.Helper()
	_, file, line, _ := runtime.Caller(skip + 1)
	t.Fatalf("%s:%d %s %v", file, line, msg, args)
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

// Wrap testing.TB into type supporting more advanced testing features
// Type returned implements testing.TB interface and can be used as this interface
func WrapTB(t testing.TB) *internal.TB {
    t.Helper()
	return &internal.TB{TB: t}
}

// Wrap testing.T into type supporting more advanced testing features
// Type returned implements testing.T interface and can be used as this interface
func WrapT(t *testing.T) *internal.T {
    t.Helper()
	return &internal.T{T: t}
}

// Wrap testing.B into type supporting more advanced testing features
// Type returned implements testing.B interface and can be used as this interface
func WrapB(t *testing.B) *internal.B {
    t.Helper()
	return &internal.B{B: t}
}

// Wrap testing.F into type supporting more advanced testing features
// Type returned implements testing.F interface and can be used as this interface
func WrapF(t *testing.F) *internal.F {
    t.Helper()
	return &internal.F{F: t}
}
