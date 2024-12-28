package errors

import "fmt"

type Error = error
type Exception = error

// NullError is a type that is used to note that a value of error is null
// This is meant for situations where you need empty error value
type NullError struct {
}

func (NullError) Error() string {
	return "NullError"
}

// FunctionalError is a error that is used to note that error happened based on input or other assertions
type FunctionalError struct {
}

func (FunctionalError) Error() string {
	return "FunctionalError"
}

// AssertionError is a error that is used to note that assertion failed
type AssertionError struct {
	CustomMessage string
}

func (ae AssertionError) Error() string {
	if ae.CustomMessage != "" {
		return fmt.Sprintf("AssertionError: %v", ae.CustomMessage)
	}
	return "AssertionError"
}
