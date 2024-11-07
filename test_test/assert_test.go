package test_test

import (
	"testing"

	"github.com/PAiP-Web-Language/stdlib-go/test"
)

func TestAssert(t *testing.T) {
	test.Assert(t, true, 1, 2, 3)
}

func TestAssert2(_t *testing.T) {
	t := test.WrapTB(_t)
	t.ExpectedFailure = true
	defer t.Defer()
	test.Assert(t, false, 1, 2, 3)
}
