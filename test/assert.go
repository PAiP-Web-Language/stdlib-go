package test

import (
	"testing"
)

// Assert that condition provided is true and if not it will forward args to Error function as message
// This function is meant to be used for testing
// This function is not intended to be used for asserting logic outside of tests
func Assert(t testing.TB, condition bool, args ...any) {
	t.Helper()
	if !condition {
		ErrorFormat(t, 1, "Assertion failed", args...)
	}
}