package test

import (
	"testing"

	"github.com/PAiP-Web-Language/stdlib-go/test/internal"
)

// RunSubTestOrBenchmark is a function that will run subtest or subbenchmark with specified name
// It will return true if subtest was run correctly according to (testing.T/testing.B).Run or false if failed
func RunSubTestOrBenchmark(t internal.TBI, name string, f func(t testing.TB)) bool {
	t.Helper()
	return t.RunTB(name, f)
}

// RunSubTest is a function that will run subtest with specified name
// It will return true if subtest was run correctly according to testing.T.Run or false if failed
func RunSubTest(t internal.TI, name string, f func(t *testing.T)) bool {
	t.Helper()
	return t.Run(name, f)
}

// RunSubBenchmark is a function that will run subbenchmark with specified name
// It will return true if subbenchmark was run correctly according to testing.B.Run or false if failed
func RunSubBenchmark(t internal.BI, name string, f func(t *testing.B)) bool {
	t.Helper()
	return t.Run(name, f)
}

// RunRequiredSubTestOrBenchmark is a function that will run subtest or subbenchmark with specified name
// It will fully fail (testing.TB.FailNow) main test if subtest failed according to (testing.T/testing.B).Run
func RunRequiredSubTestOrBenchmark(t internal.TBI, name string, f func(t testing.TB)) {
	t.Helper()
	tr := t.RunTB(name, f)
	if !tr {
		t.Fatalf("Subtest %s failed but was required to continue", name)
	}
}

// RunRequiredSubTest is a function that will run subtest with specified name
// It will fully fail (testing.T.FailNow) main test if subtest failed according to testing.T.Run
func RunRequiredSubTest(t internal.TI, name string, f func(t *testing.T)) {
	t.Helper()
	tr := t.Run(name, f)
	if !tr {
		t.Fatalf("Subtest %s failed but was required to continue", name)
	}
}

// RunRequiredSubBenchmark is a function that will run subbenchmark with specified name
// It will fully fail (testing.B.FailNow) main test if subtest failed according to testing.B.Run
func RunRequiredSubBenchmark(t internal.BI, name string, f func(t *testing.B)) {
	t.Helper()
	tr := t.Run(name, f)
	if !tr {
		t.Fatalf("Subtest %s failed but was required to continue", name)
	}
}
