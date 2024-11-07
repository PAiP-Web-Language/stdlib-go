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
)

// This wrapper is arround testing.F addding some additional functionality
type F struct {
	*testing.F
}

// Add will add the arguments to the seed corpus for the fuzz test. This will be
// a no-op if called after or within the fuzz target, and args must match the
// arguments for the fuzz target.
func (f *F) Add(args ...FuzzFuncArgumentType) {
	f.Helper()
	var aargs []any
	for _, arg := range args {
		aargs = append(aargs, arg)
	}
	f.F.Add(aargs...)
}

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
func (f *F) Fuzz(ff FuzzFunc) {
	f.Helper()
	f.F.Fuzz(any(ff))
}
