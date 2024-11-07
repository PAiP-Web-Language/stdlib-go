// This functions have comments copied from testing package
// These comments are from authors of code in go stdlib
// These comments can be modified a bit to match their functionality but most content is copied
// Code is my own and is intended to fix something that stdlib lacks

package internal

import (
	"errors"
	"testing"
)

var Err = errors.New("Test error")

type T struct {
	testing.TB
	ExpectedFailure bool
	failed bool
}

// Fail marks the function as having failed but continues execution.
func (t *T) Fail() {
	t.Helper()
	if t.ExpectedFailure {
		t.failed = true
	} else {
		t.TB.Fail()
	}
}

// FailNow marks the function as having failed and stops its execution
// by calling runtime.Goexit (which then runs all deferred calls in the
// current goroutine).
// Execution will continue at the next test or benchmark.
// FailNow must be called from the goroutine running the
// test or benchmark function, not from other goroutines
// created during the test. Calling FailNow does not stop
// those other goroutines.
//
// If test is marked as Failure is expected then FailNow will panic with error that is meant to be catched by T.Defer()
func (t *T) FailNow() {
	t.Helper()
	if t.ExpectedFailure {
		panic(Err)
	} else {
		t.TB.FailNow()
	}
}

// Fatal is equivalent to Log followed by FailNow.
func (t *T) Fatal(args ...any) {
	t.Helper()
	if t.ExpectedFailure {
		t.TB.Log(args...)
		t.FailNow()
	} else {
		t.TB.Fatal(args...)
	}
}

// Fatalf is equivalent to Logf followed by FailNow.
func (t *T) Fatalf(format string, args ...any) {
	t.Helper()
	if t.ExpectedFailure {
		t.TB.Logf(format, args...)
		t.FailNow()
	} else {
		t.TB.Fatalf(format, args...)
	}
}

// Error is equivalent to Log followed by Fail.
func (t *T) Error(args ...any) {
    t.Helper()
    if t.ExpectedFailure {
        t.TB.Log(args...)
        t.Fail()
    } else {
        t.TB.Error(args...)
    }
}

// Errorf is equivalent to Logf followed by Fail.
func (t *T) Errorf(format string, args ...any) {
	t.Helper()
	if t.ExpectedFailure {
		t.TB.Logf(format, args...)
		t.Fail()
	} else {
		t.TB.Errorf(format, args...)
	}
}

// Failed reports whether the function has failed.
//
// This function takes into account whether test is expected to fail
func (t *T) Failed() bool {
	if !t.ExpectedFailure {
		return t.TB.Failed() || t.failed
	}
	return t.TB.Failed() || !t.failed
}

// Defer logic for handling test failure
// This function takes into account whether test is expected to fail
func (t *T) Defer() {
	t.Helper()
	if r := recover(); r != nil {
		if t.ExpectedFailure {
			if r == Err {
				t.failed = true
			}
		}
	}
	if t.ExpectedFailure {
		if !t.failed {
			t.TB.FailNow()
		}
	} else {
		if t.failed {
			t.TB.FailNow()
		}
	}
}
