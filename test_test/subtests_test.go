package test_test

import (
	"testing"

	"github.com/PAiP-Web-Language/stdlib-go/test"
)

func TestRunSubTest1(_t *testing.T) {
	t := test.WrapT(_t)
	defer t.Defer()
	test.RunSubTestOrBenchmark(t, "TestRunSubTest1", func(t testing.TB) {
		test.Assert(t, true, "TestRunSubTest1 should pass")
	})
}

func TestRunSubTest2(_t *testing.T) {
	t := test.WrapT(_t)
	t.ExpectedFailure = true
	defer t.Defer()
	test.LogIfFailed(t, "Main test should fail if subtest fails")
	var stf bool = false
	test.RunRequiredSubTestOrBenchmark(t, "TestRunSubTest2", func(t2 testing.TB) {
		st := test.WrapTB(t2)
		st.ExpectedFailure = true
		defer st.Defer()
		test.Assert(st, false, "TestRunSubTest2 should fail")

		stf = !st.Failed()
	})
	if stf {
		t.FailNow()
	}
	t.Log("Main test should fail if subtest fails and not continue")
}
