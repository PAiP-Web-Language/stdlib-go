package test

import (
	"fmt"
	"testing"
)

// IfFailed is a function that will call function f if test failed
func IfFailed(t testing.TB, f func()) {
	t.Helper()
	t.Cleanup(func() {
		t.Helper()
		if t.Failed() {
			f()
		}
	})
}

// LogIfFailed is a function that will log message in specified format if test failed
func LogIfFailed(t testing.TB, format string, args ...any) {
	t.Helper()
	msg := fmt.Sprintf(format, args)
	t.Cleanup(func() {
		t.Helper()
		if t.Failed() {
			t.Log(msg)
		}
	})
}
