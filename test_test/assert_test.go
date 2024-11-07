package test_test

import (
	"testing"

	"github.com/PAiP-Web-Language/stdlib-go/test"
)

func TestAssert(t *testing.T) {
	test.Assert(t, true, "Assertion should pass on true")
}

func TestAssert2(_t *testing.T) {
	t := test.WrapTB(_t)
	t.ExpectedFailure = true
	defer t.Defer()
	test.LogIfFailed(t, "Assertion should fail on false")
	test.Assert(t, false, "Assertion should fail on false (Expected fail)")
}
