package test_test

import (
	"testing"

	"github.com/PAiP-Web-Language/stdlib-go/test"
)

type TestSetPrivateMemberStruct struct {
	a int
}

func (t *TestSetPrivateMemberStruct) GetA() int {
	return t.a
}

func TestSetPrivateMember1(_t *testing.T) {
	t := test.WrapTB(_t)
	defer t.Defer()
	test.LogIfFailed(t, "SetPrivateMember should pass on correct value")
	t1 := &TestSetPrivateMemberStruct{}
	test.Assert(t, t1.GetA() == 0, "SetPrivateMember should pass on correct value")
	test.SetPrivateMember(t, t1, "a", 1)
	test.Assert(t, t1.GetA() == 1, "SetPrivateMember should pass on correct value")
}

func TestSetPrivateMember2(_t *testing.T) {
	t := test.WrapTB(_t)
	t.ExpectedFailure = true
	defer t.Defer()
	test.LogIfFailed(t, "SetPrivateMember should fail on incorrect field")
	t1 := &TestSetPrivateMemberStruct{}
	test.SetPrivateMember(t, t1, "A", 2)
}
