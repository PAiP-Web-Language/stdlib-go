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

func (t *T) Fail() {
	t.Helper()
	if t.ExpectedFailure {
		t.failed = true
	} else {
		t.TB.Fail()
	}
}

func (t *T) FailNow() {
	t.Helper()
	if t.ExpectedFailure {
		panic(Err)
	} else {
		t.TB.FailNow()
	}
}

func (t *T) Fatal(args ...any) {
	t.Helper()
	if t.ExpectedFailure {
		t.TB.Log(args...)
		t.FailNow()
	} else {
		t.TB.Fatal(args...)
	}
}

func (t *T) Fatalf(format string, args ...any) {
	t.Helper()
	if t.ExpectedFailure {
		t.TB.Logf(format, args...)
		t.FailNow()
	} else {
		t.TB.Fatalf(format, args...)
	}
}

func (t *T) Error(args ...any) {
    t.Helper()
    if t.ExpectedFailure {
        t.TB.Log(args...)
        t.Fail()
    } else {
        t.TB.Error(args...)
    }
}

func (t *T) Errorf(format string, args ...any) {
	t.Helper()
	if t.ExpectedFailure {
		t.TB.Logf(format, args...)
		t.Fail()
	} else {
		t.TB.Errorf(format, args...)
	}
}

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
