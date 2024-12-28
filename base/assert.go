package base

import "github.com/PAiP-Web-Language/stdlib-go/errors"

// Assert that condition is true if it's not panic with AssertionError
func Assert(cond bool) {
	if !cond {
		panic(errors.AssertionError{})
	}
}

// Assert that condition is true if it's not panic with AssertionError with custom message suffix
func AssertMsg(cond bool, msg string) {
	if !cond {
		panic(errors.AssertionError{CustomMessage: msg})
	}
}

// Assert that err is nil if it's not panic with AssertionError and err message as custom message suffix
func AssertErr(err error) {
	if err != nil {
		panic(errors.AssertionError{CustomMessage: err.Error()})
	}
}

// Assert that v is compatible with T
// This function is noop. It's basing on LSP and build system to inform you of failed assertion
func AssertCompatibleType[T any](v T) {
}
