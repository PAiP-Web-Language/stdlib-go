package base

import (
	"fmt"

	"github.com/PAiP-Web-Language/stdlib-go/errors"
)

// TypedErrorResult is Result wrapper type that additionally annotates error types
// This error type is not checked its just annotation
type TypedErrorResult[T any, E any] struct {
	*Result[T]
}

// TypedErrorResultE is Result wrapper type that additionally annotates error types
// This error type is not checked its just annotation
// This structure is same as TypedErrorResult but error type on this one is more strict to error interface
// Whereas TypedErrorResult is taking anything
type TypedErrorResultE[T any, E error] struct {
	*Result[T]
}

// Result is a type that represents a value that may be an error
type Result[T any] struct {
	Data  T
	Err error
}

// Make Result
func MakeResult[T any]() Result[T] {
	return Result[T]{}
}

// Make Ok Result
func Ok[T any](value T) Result[T] {
	r := MakeResult[T]()
	return r.Ok(value)
}

// Make Error Result
func Err[T any](err error) Result[T] {
	r := MakeResult[T]()
	return r.Error(err)
}

// Make Error Result with additional data
func ErrWithValue[T any](value T, err error) Result[T] {
	return Result[T]{Data: value, Err: err}
}

// Get Pointer of Result
func (r Result[T]) ToPointer() *Result[T] {
	return &r
}

// Get Result from Pointer
func (r *Result[T]) ToValue() Result[T] {
	return *r
}

// Ok Result
func (r *Result[T]) Ok(value T) Result[T] {
	r.Data = value
	return *r
}

// Error Result
func (r *Result[T]) Error(err error) Result[T] {
	r.Err = err
	return *r
}

// Run function fn when Result is Ok
func (r Result[T]) AndThen(fn func(T) any) any {
	if r.IsError() {
		return nil
	}
	return fn(r.Data)
}

// String Convert value to string
func (r Result[T]) String() string {
	if r.IsError() {
		return fmt.Sprintf("%s", r.Err)
	}
	return fmt.Sprintf("%s", any(r.Data))
}

// Get Go String
func (r Result[T]) GoString() string {
	var ref T
	return fmt.Sprintf(
		"Result[%T]{DataValue:%#v,ErrorValue:%#v}",
		ref,
		r.Data,
		r.Err,
	)
}

// Get Value
func (r Result[T]) GetValue() T {
	return r.Data
}

// Check if Result has error
func (r Result[T]) IsError() bool {
	if r.Err == nil {
		return false
	}
	return true
}

// If result has error get error otherwise return nil
func (r Result[T]) GetError() error {
	return r.Err
}

// Get underlying value or panic if error
func (r Result[T]) Unwrap() T {
	if r.IsError() {
		panic(errors.UnwrapError{})
	}
	return r.Data
}

// Get underlying error or panic if correct value
func (r Result[T]) UnwrapErr() error {
	if !r.IsError() {
		panic(errors.UnwrapError{})
	}
	return r.Err
}

// Get underlying value or return default value if error is found
func (r Result[T]) UnwrapOr(defaultVal T) T {
	if r.IsError() {
		return defaultVal
	}
	return r.Data
}

// Unwrap but if error occurs run specified function f instead of panic
// Return result of underlying value (having err if it's found)
func (r Result[T]) UnwrapOrErr(f func(Result[T])) T {
	if r.IsError() {
		f(r)
	}
	return r.Data
}

// Unwraps as result type for if error checks and run specified function f
// instead of panic
// Return result of underlying value (having err if it's found)
func (r Result[T]) UnwrapAsResultOrErr(f func(Result[T])) Result[T] {
	if r.IsError() {
		f(r)
	}
	return r
}

// Unwrap value and error separately (Result -> Go normal returns)
func (r Result[T]) UnwrapWithErr() (T, error) {
	return r.Data, r.Err
}

// Expect correct value if error is found panic with specified message
func (r Result[T]) Expect(err any) {
	if r.IsError() {
		panic(err)
	}
}

// Expect error value if error is not found panic with specified message
func (r Result[T]) ExpectErr(err any) {
	if !r.IsError() {
		panic(err)
	}
}

// Get value or null
func (r Result[T]) Get() T {
	if r.IsError() {
		return Null[T]().ValueOrZero()
	}
	return r.Data
}

// Get value or nil as any
func (r Result[T]) GetAny() any {
	if r.IsError() {
		return nil
	}
    return r.Data
}

// Check if value is Ok
func (r Result[T]) IsOk() bool {
    return !r.IsError()
}

// Check if value is Err
func (r Result[T]) IsErr() bool {
    return r.IsError()
}

// Inspect value of this object
func (n Result[T]) Inspect(f func(Result[T])) Result[T] {
	if n.IsOk() {
		f(n)
	}
	return n
}

// Inspect value of this objects reference
func (n *Result[T]) InspectRef(f func(*Result[T])) *Result[T] {
	if n.IsOk() {
		f(n)
	}
	return n
}

// Inspect value of this object
func (n Result[T]) InspectErr(f func(Result[T])) Result[T] {
	if n.IsErr() {
		f(n)
	}
	return n
}

// Inspect value of this objects reference
func (n *Result[T]) InspectErrRef(f func(*Result[T])) *Result[T] {
	if n.IsErr() {
		f(n)
	}
	return n
}

// Cast to Result
func (n Result[T]) ToResult() Result[any] {
	return Result[any]{Data: n.Data, Err: n.Err}
}

// Cast to Option
func (n Result[T]) ToOption() Option[any] {
	if n.IsError() {
		return None[any]()
	}
	return Some[any](n.Data)
}

// Cast to Option
func (n Result[T]) ToOptionT() Option[T] {
	if n.IsError() {
		return None[T]()
	}
	return Some[T](n.Data)
}

// Cast to Nullable
func (n Result[T]) ToNullable() Nullable[any] {
	if n.IsError() {
		return Null[any]()
	}
	return NullableValue[any](n.Data)
}

// Cast to Nullable
func (n Result[T]) ToNullableT() Nullable[T] {
	if n.IsError() {
		return Null[T]()
	}
	return NullableValue[T](n.Data)
}

// Clone this object
func (n Result[T]) Clone() Result[T] {
	return Result[T]{Data: n.Data, Err: n.Err}
}
