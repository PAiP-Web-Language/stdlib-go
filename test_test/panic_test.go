package test_test

import (
	"testing"
	"github.com/PAiP-Web-Language/stdlib-go/test"
)

func TestPanic1(_t *testing.T) {
    t := test.WrapTB(_t)
    defer t.Defer()
    test.LogIfFailed(t, "Panic should pass on panic")
    test.Panic(t, func() { panic("Panic") }, "Panic should pass on panic")
}

func TestPanic2(_t *testing.T) {
    t := test.WrapTB(_t)
	t.ExpectedFailure = true
    defer t.Defer()
    test.LogIfFailed(t, "Panic should fail on non panic")
    test.Panic(t, func() {}, "Panic should fail on non panic")
}

func TestNotPanic1(_t *testing.T) {
    t := test.WrapTB(_t)
    defer t.Defer()
    test.LogIfFailed(t, "NotPanic should pass on non panic")
    test.NotPanic(t, func() {}, "NotPanic should pass on non panic")
}

func TestNotPanic2(_t *testing.T) {
    t := test.WrapTB(_t)
	t.ExpectedFailure = true
    defer t.Defer()
    test.LogIfFailed(t, "NotPanic should fail on panic")
    test.NotPanic(t, func() { panic("Panic") }, "NotPanic should fail on panic")
}

func TestPanicMatches1(_t *testing.T) {
    t := test.WrapTB(_t)
    defer t.Defer()
    test.LogIfFailed(t, "PanicMatches should pass on panic with same error")
    test.PanicMatches(t, func() { panic("Panic") }, "Panic", "Panic should pass on panic")
}

func TestPanicMatches2(_t *testing.T) {
    t := test.WrapTB(_t)
	t.ExpectedFailure = true
    defer t.Defer()
    test.LogIfFailed(t, "PanicMatches should fail on panic with different error")
    test.PanicMatches(t, func() { panic("Panic1") }, "Panic", "NotPanic", "Panic should fail on panic")
}
