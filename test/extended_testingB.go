// This functions have comments copied from testing package
// These comments are from authors of code in go stdlib
// These comments can be modified a bit to match their functionality but most content is copied
// Code is my own and is intended to fix something that stdlib lacks


// Notice for comments taken from testing package (stdlib of go)

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the https://github.com/golang/go/blob/go1.23.2/LICENSE file.


package test

import (
	"testing"
)

// This wrapper is arround testing.B addding some additional functionality
type B struct {
	*testing.B
	ExpectedFailure bool
	failed bool
}

// Fail marks the function as having failed but continues execution.
func (t *B) Fail() {
	t.Helper()
	if t.ExpectedFailure {
		t.failed = true
	} else {
		t.B.Fail()
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
func (t *B) FailNow() {
	t.Helper()
	if t.ExpectedFailure {
		panic(Err)
	} else {
		t.B.FailNow()
	}
}

// Fatal is equivalent to Log followed by FailNow.
func (t *B) Fatal(args ...any) {
	t.Helper()
	if t.ExpectedFailure {
		t.B.Log(args...)
		t.FailNow()
	} else {
		t.B.Fatal(args...)
	}
}

// Fatalf is equivalent to Logf followed by FailNow.
func (t *B) Fatalf(format string, args ...any) {
	t.Helper()
	if t.ExpectedFailure {
		t.B.Logf(format, args...)
		t.FailNow()
	} else {
		t.B.Fatalf(format, args...)
	}
}

// Error is equivalent to Log followed by Fail.
func (t *B) Error(args ...any) {
    t.Helper()
    if t.ExpectedFailure {
        t.B.Log(args...)
        t.Fail()
    } else {
        t.B.Error(args...)
    }
}

// Errorf is equivalent to Logf followed by Fail.
func (t *B) Errorf(format string, args ...any) {
	t.Helper()
	if t.ExpectedFailure {
		t.B.Logf(format, args...)
		t.Fail()
	} else {
		t.B.Errorf(format, args...)
	}
}

// Failed reports whether the function has failed.
//
// This function takes into account whether test is expected to fail
func (t *B) Failed() bool {
	if !t.ExpectedFailure {
		return t.B.Failed() || t.failed
	}
	return t.B.Failed() || !t.failed
}

// Defer logic for handling test failure
// This function takes into account whether test is expected to fail
func (t *B) Defer() {
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
			t.B.FailNow()
		}
	} else {
		if t.failed {
			t.B.FailNow()
		}
	}
}

// Run benchmarks f as a subbenchmark with the given name. It reports
// whether there were any failures.
//
// A subbenchmark is like any other benchmark. A benchmark that calls Run at
// least once will not be measured itself and will be called once with N=1.
func (t *B) RunTB(name string, f func(t testing.TB)) bool {
	t.Helper()
	return t.B.Run(name, func(t *testing.B) {
		f(t)
	})
}
