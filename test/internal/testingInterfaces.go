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
	"time"

	"github.com/PAiP-Web-Language/stdlib-go/base"
)

// TBFI is the interface common to T, B, and F.
type TBFI interface {
	// Cleanup registers a function to be called when the test (or subtest) and all its
	// subtests complete. Cleanup functions will be called in last added,
	// first called order.
	Cleanup(func())
	// Error is equivalent to Log followed by Fail.
	Error(args ...any)
	// Errorf is equivalent to Logf followed by Fail.
	Errorf(format string, args ...any)
	// Fail marks the function as having failed but continues execution.
	Fail()
	// FailNow marks the function as having failed and stops its execution
	// by calling runtime.Goexit (which then runs all deferred calls in the
	// current goroutine).
	// Execution will continue at the next test or benchmark.
	// FailNow must be called from the goroutine running the
	// test or benchmark function, not from other goroutines
	// created during the test. Calling FailNow does not stop
	// those other goroutines.
	FailNow()
	// Failed reports whether the function has failed.
	Failed() bool
	// Fatal is equivalent to Log followed by FailNow.
	Fatal(args ...any)
	// Fatalf is equivalent to Logf followed by FailNow.
	Fatalf(format string, args ...any)
	// Helper marks the calling function as a test helper function.
	// When printing file and line information, that function will be skipped.
	// Helper may be called simultaneously from multiple goroutines.
	Helper()
	// Log formats its arguments using default formatting, analogous to Println,
	// and records the text in the error log. For tests, the text will be printed only if
	// the test fails or the -test.v flag is set. For benchmarks, the text is always
	// printed to avoid having performance depend on the value of the -test.v flag.
	Log(args ...any)
	// Logf formats its arguments according to the format, analogous to Printf, and
	// records the text in the error log. A final newline is added if not provided. For
	// tests, the text will be printed only if the test fails or the -test.v flag is
	// set. For benchmarks, the text is always printed to avoid having performance
	// depend on the value of the -test.v flag.
	Logf(format string, args ...any)
	// Name returns the name of the running (sub-) test or benchmark.
	//
	// The name will include the name of the test along with the names of
	// any nested sub-tests. If two sibling sub-tests have the same name,
	// Name will append a suffix to guarantee the returned name is unique.
	Name() string
	// Setenv calls os.Setenv(key, value) and uses Cleanup to
	// restore the environment variable to its original value
	// after the test.
	//
	// Because Setenv affects the whole process, it cannot be used
	// in parallel tests or tests with parallel ancestors.
	Setenv(key, value string)
	// Skip is equivalent to Log followed by SkipNow.
	Skip(args ...any)
	// Skipf is equivalent to Logf followed by SkipNow.
	Skipf(format string, args ...any)
	// SkipNow marks the test as having been skipped and stops its execution
	// by calling [runtime.Goexit].
	// If a test fails (see Error, Errorf, Fail) and is then skipped,
	// it is still considered to have failed.
	// Execution will continue at the next test or benchmark. See also FailNow.
	// SkipNow must be called from the goroutine running the test, not from
	// other goroutines created during the test. Calling SkipNow does not stop
	// those other goroutines.
	SkipNow()
	// Skipped reports whether the test was skipped.
	Skipped() bool
	// TempDir returns a temporary directory for the test to use.
	// The directory is automatically removed when the test and
	// all its subtests complete.
	// Each subsequent call to t.TempDir returns a unique directory;
	// if the directory creation fails, TempDir terminates the test by calling Fatal.
	TempDir() string
}

// TBI is the interface common to T, B
type TBI interface {
	// TBFI
	TBFI
	// TBI

	// Run runs f as subtest or subbenchmark named using name parameter.
	//
	// Specifics for tests and benchmarks
	// /// For tests (using testing.T):
	// Run runs f as a subtest of t called name. It runs f in a separate goroutine
	// and blocks until f returns or calls t.Parallel to become a parallel test.
	// Run reports whether f succeeded (or at least did not fail before calling t.Parallel).
	//
	// Run may be called simultaneously from multiple goroutines, but all such calls
	// must return before the outer test function for t returns.
	//
	// /// For benchmarks (using testing.B):
	// Run benchmarks f as a subbenchmark with the given name. It reports
	// whether there were any failures.
	//
	// A subbenchmark is like any other benchmark. A benchmark that calls Run at
	// least once will not be measured itself and will be called once with N=1.
	//
	// This function exist for those types that should implement this interface
	// But testing.T and testing.B have different signatures due to them passing dev.Self (meaning testing.T / testing.B)
	// So this function cant work in this interface
	// Run(name string, f func(t TBI)) bool

	// Run runs f as subtest or subbenchmark named using name parameter.
	//
	// Specifics for tests and benchmarks
	// /// For tests (using testing.T):
	// Run runs f as a subtest of t called name. It runs f in a separate goroutine
	// and blocks until f returns or calls t.Parallel to become a parallel test.
	// Run reports whether f succeeded (or at least did not fail before calling t.Parallel).
	//
	// Run may be called simultaneously from multiple goroutines, but all such calls
	// must return before the outer test function for t returns.
	//
	// /// For benchmarks (using testing.B):
	// Run benchmarks f as a subbenchmark with the given name. It reports
	// whether there were any failures.
	//
	// A subbenchmark is like any other benchmark. A benchmark that calls Run at
	// least once will not be measured itself and will be called once with N=1.
	//
	// This function have enforced type that both should implement but only in wrappers
	RunTB(name string, f func(t testing.TB)) bool
}

// T is a type passed to Test functions to manage test state and support formatted test logs.
//
// A test ends when its Test function returns or calls any of the methods
// FailNow, Fatal, Fatalf, SkipNow, Skip, or Skipf. Those methods, as well as
// the Parallel method, must be called only from the goroutine running the
// Test function.
//
// The other reporting methods, such as the variations of Log and Error,
// may be called simultaneously from multiple goroutines.
//
// This is interface for testing.T
type TI interface {
	TBFI
	// Specific TBI

	// Run runs f as a subtest of t called name. It runs f in a separate goroutine
	// and blocks until f returns or calls t.Parallel to become a parallel test.
	// Run reports whether f succeeded (or at least did not fail before calling t.Parallel).
	//
	// Run may be called simultaneously from multiple goroutines, but all such calls
	// must return before the outer test function for t returns.
	Run(name string, f func(t *testing.T)) bool

	// TI

	// Parallel signals that this test is to be run in parallel with (and only with)
	// other parallel tests. When a test is run multiple times due to use of
	// -test.count or -test.cpu, multiple instances of a single test never run in
	// parallel with each other.
	Parallel()
	// Deadline reports the time at which the test binary will have
	// exceeded the timeout specified by the -timeout flag.
	//
	// The ok result is false if the -timeout flag indicates “no timeout” (0).
	Deadline() (deadline time.Time, ok bool)
}

// B is a type passed to [Benchmark] functions to manage benchmark
// timing and to specify the number of iterations to run.
//
// A benchmark ends when its Benchmark function returns or calls any of the methods
// FailNow, Fatal, Fatalf, SkipNow, Skip, or Skipf. Those methods must be called
// only from the goroutine running the Benchmark function.
// The other reporting methods, such as the variations of Log and Error,
// may be called simultaneously from multiple goroutines.
//
// Like in tests, benchmark logs are accumulated during execution
// and dumped to standard output when done. Unlike in tests, benchmark logs
// are always printed, so as not to hide output whose existence may be
// affecting benchmark results.
//
// This is interface for testing.B
type BI interface {
	TBFI
	// Specific TBI

	// Run benchmarks f as a subbenchmark with the given name. It reports
	// whether there were any failures.
	//
	// A subbenchmark is like any other benchmark. A benchmark that calls Run at
	// least once will not be measured itself and will be called once with N=1.
	Run(name string, f func(t *testing.B)) bool
	// BI

	// N int

	// Elapsed returns the measured elapsed time of the benchmark.
	// The duration reported by Elapsed matches the one measured by
	// [BI.StartTimer], [BI.StopTimer], and [BI.ResetTimer].
	Elapsed() time.Duration
	// ReportAllocs enables malloc statistics for this benchmark.
	// It is equivalent to setting -test.benchmem, but it only affects the
	// benchmark function that calls ReportAllocs.
	ReportAllocs()
	// ReportMetric adds "n unit" to the reported benchmark results.
	// If the metric is per-iteration, the caller should divide by b.N,
	// and by convention units should end in "/op".
	// ReportMetric overrides any previously reported value for the same unit.
	// ReportMetric panics if unit is the empty string or if unit contains
	// any whitespace.
	// If unit is a unit normally reported by the benchmark framework itself
	// (such as "allocs/op"), ReportMetric will override that metric.
	// Setting "ns/op" to 0 will suppress that built-in metric.
	ReportMetric(n float64, unit string)
	// ResetTimer zeroes the elapsed benchmark time and memory allocation counters
	// and deletes user-reported metrics.
	// It does not affect whether the timer is running.
	ResetTimer()
	// RunParallel runs a benchmark in parallel.
	// It creates multiple goroutines and distributes b.N iterations among them.
	// The number of goroutines defaults to GOMAXPROCS. To increase parallelism for
	// non-CPU-bound benchmarks, call [BI.SetParallelism] before RunParallel.
	// RunParallel is usually used with the go test -cpu flag.
	//
	// The body function will be run in each goroutine. It should set up any
	// goroutine-local state and then iterate until pb.Next returns false.
	// It should not use the [BI.StartTimer], [BI.StopTimer], or [BI.ResetTimer] functions,
	// because they have global effect. It should also not call [BI.Run].
	//
	// RunParallel reports ns/op values as wall time for the benchmark as a whole,
	// not the sum of wall time or CPU time over each parallel goroutine.
	RunParallel(func(p *testing.PB))
	// SetBytes records the number of bytes processed in a single operation.
	// If this is called, the benchmark will report ns/op and MB/s.
	SetBytes(n int64)
	// SetParallelism sets the number of goroutines used by [BI.RunParallel] to p*GOMAXPROCS.
	// There is usually no need to call SetParallelism for CPU-bound benchmarks.
	// If p is less than 1, this call will have no effect.
	SetParallelism(p int)
	// StartTimer starts timing a test. This function is called automatically
	// before a benchmark starts, but it can also be used to resume timing after
	// a call to [BI.StopTimer].
	StartTimer()
	// StopTimer stops timing a test. This can be used to pause the timer
	// while performing complex initialization that you don't
	// want to measure.
	StopTimer()
}

// FuzzFuncArgumentType is a type alias for all argument types that fuzz tests support according to documentation
// This type is just type annotation for all types supported but works same as any
type FuzzFuncArgumentType = base.LUI17[
	[]byte,
	string,
	bool,
	byte,
	rune,
	float32,
	float64,
	int,
	int8,
	int16,
	int32,
	int64,
	uint,
	uint8,
	uint16,
	uint32,
	uint64,
]

// FuzzFunc is a type alias for function that fuzz tests support according to documentation
// This type is just type annotation for function type supported but works same as any so it doesn't enforce anything
type FuzzFunc = base.LUI1[func(t *testing.T, args ...FuzzFuncArgumentType)]

// F is a type passed to fuzz tests.
//
// Fuzz tests run generated inputs against a provided fuzz target, which can
// find and report potential bugs in the code being tested.
//
// A fuzz test runs the seed corpus by default, which includes entries provided
// by (*F).Add and entries in the testdata/fuzz/<FuzzTestName> directory. After
// any necessary setup and calls to (*F).Add, the fuzz test must then call
// (*F).Fuzz to provide the fuzz target. See the testing package documentation
// for an example, and see the [FI.Fuzz] and [FI.Add] method documentation for
// details.
//
// *F methods can only be called before (*F).Fuzz. Once the test is
// executing the fuzz target, only (*T) methods can be used. The only *F methods
// that are allowed in the (*F).Fuzz function are (*F).Failed and (*F).Name.
//
// This is interface for testing.F with correct annotations for types
type FIT interface {
	TBFI
	// FI

	// Add will add the arguments to the seed corpus for the fuzz test. This will be
	// a no-op if called after or within the fuzz target, and args must match the
	// arguments for the fuzz target.
	Add(args ...FuzzFuncArgumentType)
	// Fuzz runs the fuzz function, ff, for fuzz testing. If ff fails for a set of
	// arguments, those arguments will be added to the seed corpus.
	//
	// ff must be a function with no return value whose first argument is *T and
	// whose remaining arguments are the types to be fuzzed.
	// For example:
	//
	//	f.Fuzz(func(t *testing.T, b []byte, i int) { ... })
	//
	// The following types are allowed: []byte, string, bool, byte, rune, float32,
	// float64, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64.
	// More types may be supported in the future.
	//
	// ff must not call any *F methods, e.g. (*F).Log, (*F).Error, (*F).Skip. Use
	// the corresponding *T method instead. The only *F methods that are allowed in
	// the (*F).Fuzz function are (*F).Failed and (*F).Name.
	//
	// This function should be fast and deterministic, and its behavior should not
	// depend on shared state. No mutable input arguments, or pointers to them,
	// should be retained between executions of the fuzz function, as the memory
	// backing them may be mutated during a subsequent invocation. ff must not
	// modify the underlying data of the arguments provided by the fuzzing engine.
	//
	// When fuzzing, F.Fuzz does not return until a problem is found, time runs out
	// (set with -fuzztime), or the test process is interrupted by a signal. F.Fuzz
	// should be called exactly once, unless F.Skip or [F.Fail] is called beforehand.
	Fuzz(ff FuzzFunc)
}

// F is a type passed to fuzz tests.
//
// Fuzz tests run generated inputs against a provided fuzz target, which can
// find and report potential bugs in the code being tested.
//
// A fuzz test runs the seed corpus by default, which includes entries provided
// by (*F).Add and entries in the testdata/fuzz/<FuzzTestName> directory. After
// any necessary setup and calls to (*F).Add, the fuzz test must then call
// (*F).Fuzz to provide the fuzz target. See the testing package documentation
// for an example, and see the [FI.Fuzz] and [FI.Add] method documentation for
// details.
//
// *F methods can only be called before (*F).Fuzz. Once the test is
// executing the fuzz target, only (*T) methods can be used. The only *F methods
// that are allowed in the (*F).Fuzz function are (*F).Failed and (*F).Name.
//
// This is interface for testing.F
type FI interface {
	TBFI
	// FI

	// Add will add the arguments to the seed corpus for the fuzz test. This will be
	// a no-op if called after or within the fuzz target, and args must match the
	// arguments for the fuzz target.
	Add(args ...any)
	// Fuzz runs the fuzz function, ff, for fuzz testing. If ff fails for a set of
	// arguments, those arguments will be added to the seed corpus.
	//
	// ff must be a function with no return value whose first argument is *T and
	// whose remaining arguments are the types to be fuzzed.
	// For example:
	//
	//	f.Fuzz(func(t *testing.T, b []byte, i int) { ... })
	//
	// The following types are allowed: []byte, string, bool, byte, rune, float32,
	// float64, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64.
	// More types may be supported in the future.
	//
	// ff must not call any *F methods, e.g. (*F).Log, (*F).Error, (*F).Skip. Use
	// the corresponding *T method instead. The only *F methods that are allowed in
	// the (*F).Fuzz function are (*F).Failed and (*F).Name.
	//
	// This function should be fast and deterministic, and its behavior should not
	// depend on shared state. No mutable input arguments, or pointers to them,
	// should be retained between executions of the fuzz function, as the memory
	// backing them may be mutated during a subsequent invocation. ff must not
	// modify the underlying data of the arguments provided by the fuzzing engine.
	//
	// When fuzzing, F.Fuzz does not return until a problem is found, time runs out
	// (set with -fuzztime), or the test process is interrupted by a signal. F.Fuzz
	// should be called exactly once, unless F.Skip or [F.Fail] is called beforehand.
	Fuzz(ff any)
}
