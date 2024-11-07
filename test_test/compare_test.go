package test_test

import (
	"testing"

	"github.com/PAiP-Web-Language/stdlib-go/test"
)

func TestCompareEqual1(_t *testing.T) {
	t := test.WrapTB(_t)
	defer t.Defer()
	test.LogIfFailed(t, "CompareEqual should pass on equal values")
	test.CompareEqual(t, 1, 1, "CompareEqual should pass on equal values")
}

func TestCompareEqual2(_t *testing.T) {
	t := test.WrapTB(_t)
	t.ExpectedFailure = true
	defer t.Defer()
	test.LogIfFailed(t, "CompareEqual should fail on unequal values")
	test.CompareEqual(t, 1, 2, "CompareEqual should fail on unequal values (Expected fail)")
}

func TestEqual1(_t *testing.T) {
	t := test.WrapTB(_t)
	defer t.Defer()
	test.LogIfFailed(t, "Equal should pass on equal values")
	test.Equal(t, 1, 1, "Equal should pass on equal values")
}

func TestEqual2(_t *testing.T) {
	t := test.WrapTB(_t)
	t.ExpectedFailure = true
	defer t.Defer()
	test.LogIfFailed(t, "Equal should fail on unequal values")
	test.Equal(t, 1, 2, "Equal should fail on unequal values (Expected fail)")
}

func TestDeepEqual1(_t *testing.T) {
	t := test.WrapTB(_t)
	defer t.Defer()
	test.LogIfFailed(t, "DeepEqual should pass on equal values")
	test.DeepEqual(t, 1, 1, "DeepEqual should pass on equal values")
}

func TestDeepEqual2(_t *testing.T) {
	t := test.WrapTB(_t)
	t.ExpectedFailure = true
	defer t.Defer()
	test.LogIfFailed(t, "DeepEqual should fail on unequal values")
	test.DeepEqual(t, 1, 2, "DeepEqual should fail on unequal values (Expected fail)")
}

func TestCompareNotEqual1(_t *testing.T) {
	t := test.WrapTB(_t)
	defer t.Defer()
	test.LogIfFailed(t, "CompareNotEqual should pass on unequal values")
	test.CompareNotEqual(t, 1, 2, "CompareNotEqual should pass on unequal values")
}

func TestCompareNotEqual2(_t *testing.T) {
	t := test.WrapTB(_t)
	t.ExpectedFailure = true
	defer t.Defer()
	test.LogIfFailed(t, "CompareNotEqual should fail on equal values")
	test.CompareNotEqual(t, 1, 1, "CompareNotEqual should fail on equal values (Expected fail)")
}

func TestNotEqual1(_t *testing.T) {
	t := test.WrapTB(_t)
	defer t.Defer()
	test.LogIfFailed(t, "NotEqual should pass on unequal values")
	test.NotEqual(t, 1, 2, "NotEqual should pass on unequal values")
}

func TestNotEqual2(_t *testing.T) {
	t := test.WrapTB(_t)
	t.ExpectedFailure = true
	defer t.Defer()
	test.LogIfFailed(t, "NotEqual should fail on equal values")
	test.NotEqual(t, 1, 1, "NotEqual should fail on equal values (Expected fail)")
}

func TestNotDeepEqual1(_t *testing.T) {
	t := test.WrapTB(_t)
	defer t.Defer()
	test.LogIfFailed(t, "NotDeepEqual should pass on unequal values")
	test.NotDeepEqual(t, 1, 2, "NotDeepEqual should pass on unequal values")
}

func TestNotDeepEqual2(_t *testing.T) {
	t := test.WrapTB(_t)
	t.ExpectedFailure = true
	defer t.Defer()
	test.LogIfFailed(t, "NotDeepEqual should fail on equal values")
	test.NotDeepEqual(t, 1, 1, "NotDeepEqual should fail on equal values (Expected fail)")
}

func TestNil1(_t *testing.T) {
	t := test.WrapTB(_t)
	defer t.Defer()
	test.LogIfFailed(t, "Nil should pass on nil value")
	test.Nil(t, nil, "Nil should pass on nil value")
}

func TestNil2(_t *testing.T) {
	t := test.WrapTB(_t)
	t.ExpectedFailure = true
	defer t.Defer()
	test.LogIfFailed(t, "Nil should fail on non nil value")
	test.Nil(t, 1, "Nil should fail on non nil value (Expected fail)")
}

func TestNotNil1(_t *testing.T) {
	t := test.WrapTB(_t)
	defer t.Defer()
	test.LogIfFailed(t, "NotNil should pass on non nil value")
	test.NotNil(t, 1, "NotNil should pass on non nil value")
}

func TestNotNil2(_t *testing.T) {
	t := test.WrapTB(_t)
	t.ExpectedFailure = true
	defer t.Defer()
	test.LogIfFailed(t, "NotNil should fail on nil value")
	test.NotNil(t, nil, "NotNil should fail on nil value (Expected fail)")
}
