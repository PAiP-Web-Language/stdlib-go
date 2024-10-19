package base

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/PAiP-Web-Language/stdlib-go/errors"
)

// Option represents data that can be Some(v) or None
type Option[T any] struct {
	Data  T
	Some bool
}

// Some Create a Option from a value
func Some[T any](value T) Option[T] {
	return Option[T]{Data: value, Some: true}
}

// OptionValueFromPointer Create a Option from a pointer
func SomeFromPointer[T any](value *T) Option[T] {
	if value == nil {
		return Option[T]{Some: false}
	}
	return Some(*value)
}

// None Create a Option that is NULL with type T
func None[T any]() Option[T] {
	return Option[T]{}
}

// SomeDefaultValue Get default zero value of type T
func SomeDefaultValue[T any]() Option[T] {
	return Some(Null[T]().ValueOrZero())
}

// ValueOrZero Get Value, or default zero value if it is NULL
func (n Option[T]) ValueOrZero() T {
	if !n.Some {
		var ref T
		return ref
	}
	return n.Data
}

// IsDefault Check if value is default value
func (n Option[T]) IsDefault() bool {
	if !n.Some {
		return true
	}
	var ref T
	return any(ref) == any(n.Data)
}

// Equal Check if this Option is equal to another Option
func (n Option[T]) Equal(other Option[T]) bool {
	switch any(n.Data).(type) {
	case time.Time:
		nValue := any(n.Data).(time.Time)
		otherValue := any(other.Data).(time.Time)
		return n.Some == other.Some && (!n.Some || nValue.Equal(otherValue))
	}
	return n.ExactEqual(other)
}

// ExactEqual Check if this Option is exact equal to another Option, never using intern Equal method to check equality
func (n Option[T]) ExactEqual(other Option[T]) bool {
	return n.Some == other.Some && (!n.Some || any(n.Data) == any(other.Data))
}

// String Convert value to string
func (n Option[T]) String() string {
	return fmt.Sprintf("%s", any(n.Data))
}

// Get Go String
func (n Option[T]) GoString() string {
	var ref T
	return fmt.Sprintf(
		"Option[%T]{Data:%#v,Valid:%#v}",
		ref,
		n.Data,
		n.Some,
	)
}

// Get underlying value
func (n Option[T]) Value() (driver.Value, error) {
	if !n.Some {
		return nil, nil
	}
	return n.Data, nil
}

// Run function fn when value is not NULL
func (n Option[T]) AndThen(fn func(T) any) any {
	if !n.Some {
		return nil
	}
	return fn(n.Data)
}

// Get underlying value or panic if error
func (n Option[T]) Unwrap() T {
	if !n.Some {
		panic(errors.UnwrapError{})
	}
	return n.Data
}

// Get underlying value or panic if error
func (n Option[T]) UnwrapOrDefault() T {
	return n.ValueOrZero()
}

// Get underlying error or panic if correct value
func (n Option[T]) UnwrapErr() error {
	if n.Some {
		panic(errors.UnwrapError{})
	}
	return errors.OptionError{}
}

// Get underlying value or return default value if error is found
func (n Option[T]) UnwrapOr(defaultVal T) T {
	if !n.Some {
		return defaultVal
	}
	return n.Data
}

// Unwrap but if error occurs run specified function f instead of panic
// Return result of underlying value (having err if it's found)
func (n Option[T]) UnwrapOrErr(f func(Option[T])) T {
	if !n.Some {
		f(n)
	}
	return n.ValueOrZero()
}

// Unwraps as result type for if error checks and run specified function f
// instead of panic
// Return result of underlying value (having err if it's found)
func (n Option[T]) UnwrapAsResultOrErr(f func(Option[T])) Result[T] {
	if !n.Some {
		f(n)
		return MakeErrorResult[T](errors.UnwrapError{})
	}
	return MakeOkResult[T](n.ValueOrZero())
}

// Unwrap value and error separately (Result -> Go normal returns)
func (n Option[T]) UnwrapWithErr() (T, error) {
	var err error = nil
	if !n.Some {
		err = errors.OptionError{}
	}
	return n.Data, err
}

// Expect correct value if error is found panic with specified message
func (n Option[T]) Expect(err any) {
	if !n.Some {
		panic(err)
	}
}

// Expect error value if error is not found panic with specified message
func (n Option[T]) ExpectErr(err any) {
	if n.Some {
		panic(err)
	}
}

// Check if Object has error
func (n Option[T]) IsError() bool {
	return !n.Some
}

// Get Error if Object has error or nil if not
func (n Option[T]) GetError() error {
	if n.Some {
		return nil
	}
	return errors.OptionError{}
}

// Get value or zero
func (n Option[T]) GetT() T {
    return n.ValueOrZero()
}

// Get value or zero
func (n Option[T]) GetValueT() T {
    return n.ValueOrZero()
}

// Get value or zero as any
func (n Option[T]) Get() any {
    return n.ValueOrZero()
}

// Get value or zero as any
func (n Option[T]) GetValue() any {
    return n.ValueOrZero()
}

// Check if value is None
func (n Option[T]) IsNone() bool {
    return !n.Some
}

// Check if value is Some
func (n Option[T]) IsSome() bool {
	return n.Some
}

// Cast to Pointer
func (n *Option[T]) ToPointer() *Option[T] {
	return n
}

// Cast to Value
func (n Option[T]) ToValue() Option[T] {
	return n
}

// Cast to Any Option
func (n Option[T]) ToOption() Option[any] {
	return Option[any]{Data: n.Data, Some: n.Some}
}

// Cast to Nullable
func (n Option[T]) ToNullable() Nullable[any] {
	return NullableValue[any](n.ValueOrZero())
}

// Cast to Nullable
func (n Option[T]) ToNullableT() Nullable[T] {
	return NullableValue[T](n.ValueOrZero())
}

// Clone this object
func (n Option[T]) Clone() Option[T] {
	return Option[T]{Data: n.Data, Some: n.Some}
}

// Inspect value of this object
func (n Option[T]) Inspect(f func(Option[T])) Option[T] {
	f(n)
	return n
}

// Inspect value of this objects reference
func (n *Option[T]) InspectRef(f func(*Option[T])) *Option[T] {
	f(n)
	return n
}

// Set value of this option to None
func (n Option[T]) SetNone() {
	var ref T
	n.Some = false
	n.Data = ref
}

// Convert to Result with specified error as Err() on None()
func (n Option[T]) OkOr(err error) Result[T] {
	if n.IsNone() {
		return MakeErrorResult[T](err)
	}
	return MakeOkResult[T](n.ValueOrZero())
}

// Convert to Result using specified function's result as Err() on None()
func (n Option[T]) OkOrElse(f func() error) Result[T] {
	if n.IsNone() {
		return MakeErrorResult[T](f())
	}
	return MakeOkResult[T](n.ValueOrZero())
}

// And return Some(other) if both Option are Some()
// Or return None() if one of them is None()
func (n Option[T]) And(other Option[any]) Option[any] {
	if n.IsNone() {
		return None[any]()
	}
	if other.IsNone() {
		return None[any]()
	}
	return Some[any](other.ValueOrZero())
}

// And return Some(other) if both Option are Some()
// Or return None() if one of them is None()
func (n Option[T]) AndT(other Option[T]) Option[T] {
	if n.IsNone() {
		return None[T]()
	}
	if other.IsNone() {
		return None[T]()
	}
	return Some[T](other.ValueOrZero())
}

// Or return Some(self) if self is Some
// Otherwise return None() or Some(other) depending on state of other
func (n Option[T]) Or(other Option[any]) Option[any] {
	if n.IsNone() {
		return other.Clone()
	}
	return Some[any](n.ValueOrZero())
}

// Or return Some(self) if self is Some
// Otherwise return None() or Some(other) depending on state of other
func (n Option[T]) OrT(other Option[T]) Option[T] {
	if n.IsNone() {
		return other.Clone()
	}
	return Some[T](n.ValueOrZero())
}

// Xor return Some() if exactly one of Option is Some
// Otherwise return None()
func (n Option[T]) Xor(other Option[any]) Option[any] {
	if n.IsSome() && other.IsNone() {
		return Some[any](n.ValueOrZero())
	}
	if n.IsNone() && other.IsSome() {
		return Some[any](other.ValueOrZero())
	}
	return None[any]()
}

// Xor return Some() if exactly one of Option is Some
// Otherwise return None()
func (n Option[T]) XorT(other Option[T]) Option[T] {
	if n.IsSome() && other.IsNone() {
		return Some[T](n.ValueOrZero())
	}
	if n.IsNone() && other.IsSome() {
		return Some[T](other.ValueOrZero())
	}
	return None[T]()
}
