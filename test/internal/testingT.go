// This functions have comments copied from testing package
// These comments are from authors of code in go stdlib
// These comments can be modified a bit to match their functionality but most content is copied
// Code is my own and is intended to fix something that stdlib lacks


// Notice for comments taken from testing package (stdlib of go)

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the https://github.com/golang/go/blob/go1.23.2/LICENSE file.


package internal

import (
	"testing"
)

// This wrapper is arround testing.T addding some additional functionality
type T struct {
	*testing.T
	ExpectedFailure bool
	failed bool
}

// Fail marks the function as having failed but continues execution.
func (t *T) Fail() {
	t.Helper()
	if t.ExpectedFailure {
		t.failed = true
	} else {
		t.T.Fail()
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
		t.T.FailNow()
	}
}

// Fatal is equivalent to Log followed by FailNow.
func (t *T) Fatal(args ...any) {
	t.Helper()
	if t.ExpectedFailure {
		t.T.Log(args...)
		t.FailNow()
	} else {
		t.T.Fatal(args...)
	}
}

// Fatalf is equivalent to Logf followed by FailNow.
func (t *T) Fatalf(format string, args ...any) {
	t.Helper()
	if t.ExpectedFailure {
		t.T.Logf(format, args...)
		t.FailNow()
	} else {
		t.T.Fatalf(format, args...)
	}
}

// Error is equivalent to Log followed by Fail.
func (t *T) Error(args ...any) {
    t.Helper()
    if t.ExpectedFailure {
        t.T.Log(args...)
        t.Fail()
    } else {
        t.T.Error(args...)
    }
}

// Errorf is equivalent to Logf followed by Fail.
func (t *T) Errorf(format string, args ...any) {
	t.Helper()
	if t.ExpectedFailure {
		t.T.Logf(format, args...)
		t.Fail()
	} else {
		t.T.Errorf(format, args...)
	}
}

// Failed reports whether the function has failed.
//
// This function takes into account whether test is expected to fail
func (t *T) Failed() bool {
	if !t.ExpectedFailure {
		return t.T.Failed() || t.failed
	}
	return t.T.Failed() || !t.failed
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
			t.T.FailNow()
		}
	} else {
		if t.failed {
			t.T.FailNow()
		}
	}
}

// Run runs f as a subtest of t called name. It runs f in a separate goroutine
// and blocks until f returns or calls t.Parallel to become a parallel test.
// Run reports whether f succeeded (or at least did not fail before calling t.Parallel).
//
// Run may be called simultaneously from multiple goroutines, but all such calls
// must return before the outer test function for t returns.
func (t *T) RunTB(name string, f func(t testing.TB)) bool {
	t.Helper()
	return t.T.Run(name, func(t *testing.T) {
		f(t)
	})
}
