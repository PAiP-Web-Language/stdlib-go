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
func (o Option[T]) ValueOrZero() T {
	if !o.Some {
		var ref T
		return ref
	}
	return o.Data
}

// IsDefault Check if value is default value
func (o Option[T]) IsDefault() bool {
	if !o.Some {
		return true
	}
	var ref T
	return any(ref) == any(o.Data)
}

// Equal Check if this Option is equal to another Option
func (o Option[T]) Equal(other Option[T]) bool {
	switch any(o.Data).(type) {
	case time.Time:
		nValue := any(o.Data).(time.Time)
		otherValue := any(other.Data).(time.Time)
		return o.Some == other.Some && (!o.Some || nValue.Equal(otherValue))
	}
	return o.ExactEqual(other)
}

// ExactEqual Check if this Option is exact equal to another Option, never using intern Equal method to check equality
func (o Option[T]) ExactEqual(other Option[T]) bool {
	return o.Some == other.Some && (!o.Some || any(o.Data) == any(other.Data))
}

// String Convert value to string
func (o Option[T]) String() string {
	return fmt.Sprintf("%s", any(o.Data))
}

// Get Go String
func (o Option[T]) GoString() string {
	var ref T
	return fmt.Sprintf(
		"Option[%T]{Data:%#v,Valid:%#v}",
		ref,
		o.Data,
		o.Some,
	)
}

// Get underlying value
func (o Option[T]) Value() (driver.Value, error) {
	if !o.Some {
		return nil, nil
	}
	return o.Data, nil
}

// Run function fn when value is not NULL
func (o Option[T]) AndThen(fn func(T) any) any {
	if !o.Some {
		return nil
	}
	return fn(o.Data)
}

// Get underlying value or panic if error
func (o Option[T]) Unwrap() T {
	if !o.Some {
		panic(errors.UnwrapError{})
	}
	return o.Data
}

// Get underlying value or panic if error
func (o Option[T]) UnwrapOrDefault() T {
	return o.ValueOrZero()
}

// Get underlying error or panic if correct value
func (o Option[T]) UnwrapErr() error {
	if o.Some {
		panic(errors.UnwrapError{})
	}
	return errors.OptionError{}
}

// Get underlying value or return default value if error is found
func (o Option[T]) UnwrapOr(defaultVal T) T {
	if !o.Some {
		return defaultVal
	}
	return o.Data
}

// Unwrap but if error occurs run specified function f instead of panic
// Return result of underlying value (having err if it's found)
func (o Option[T]) UnwrapOrErr(f func(Option[T])) T {
	if !o.Some {
		f(o)
	}
	return o.ValueOrZero()
}

// Unwraps as result type for if error checks and run specified function f
// instead of panic
// Return result of underlying value (having err if it's found)
func (o Option[T]) UnwrapAsResultOrErr(f func(Option[T])) Result[T] {
	if !o.Some {
		f(o)
		return Err[T](errors.UnwrapError{})
	}
	return Ok[T](o.ValueOrZero())
}

// Unwrap value and error separately (Result -> Go normal returns)
func (o Option[T]) UnwrapWithErr() (T, error) {
	var err error = nil
	if !o.Some {
		err = errors.OptionError{}
	}
	return o.Data, err
}

// Expect correct value if error is found panic with specified message
func (o Option[T]) Expect(err any) {
	if !o.Some {
		panic(err)
	}
}

// Expect error value if error is not found panic with specified message
func (o Option[T]) ExpectErr(err any) {
	if o.Some {
		panic(err)
	}
}

// Check if Object has error
func (o Option[T]) IsError() bool {
	return !o.Some
}

// Get Error if Object has error or nil if not
func (o Option[T]) GetError() error {
	if o.Some {
		return nil
	}
	return errors.OptionError{}
}

// Get value or zero
func (o Option[T]) GetT() T {
    return o.ValueOrZero()
}

// Get value or zero
func (o Option[T]) GetValueT() T {
    return o.ValueOrZero()
}

// Get value or zero as any
func (o Option[T]) Get() any {
    return o.ValueOrZero()
}

// Get value or zero as any
func (o Option[T]) GetValue() any {
    return o.ValueOrZero()
}

// Check if value is None
func (o Option[T]) IsNone() bool {
    return !o.Some
}

// Check if value is Some
func (o Option[T]) IsSome() bool {
	return o.Some
}

// Cast to Pointer
func (o *Option[T]) ToPointer() *Option[T] {
	return o
}

// Cast to Value
func (o Option[T]) ToValue() Option[T] {
	return o
}

// Cast to Any Option
func (o Option[T]) ToOption() Option[any] {
	return Option[any]{Data: o.Data, Some: o.Some}
}

// Cast to Nullable
func (o Option[T]) ToNullable() Nullable[any] {
	return NullableValue[any](o.ValueOrZero())
}

// Cast to Nullable
func (o Option[T]) ToNullableT() Nullable[T] {
	return NullableValue[T](o.ValueOrZero())
}

// Cast to Observable
func (o Option[T]) ToObservable() Observable[Option[any]] {
	return MakeObservable(o.ToOption())
}

// Cast to Observable
func (o Option[T]) ToObservableT() Observable[Option[T]] {
	return MakeObservable(o)
}

// Clone this object
func (o Option[T]) Clone() Option[T] {
	return Option[T]{Data: o.Data, Some: o.Some}
}

// Inspect value of this object
func (o Option[T]) Inspect(f func(Option[T])) Option[T] {
	if o.IsSome() {
		f(o)
	}
	return o
}

// Inspect value of this objects reference
func (o *Option[T]) InspectRef(f func(*Option[T])) *Option[T] {
	if o.IsSome() {
		f(o)
	}
	return o
}

// Inspect value of this object
func (o Option[T]) InspectNone(f func(Option[T])) Option[T] {
	if o.IsNone() {
		f(o)
	}
	return o
}

// Inspect value of this objects reference
func (o *Option[T]) InspectNoneRef(f func(*Option[T])) *Option[T] {
	if o.IsNone() {
		f(o)
	}
	return o
}

// Set value of this option to None
func (o Option[T]) SetNone() {
	var ref T
	o.Some = false
	o.Data = ref
}

// Convert to Result with specified error as Err() on None()
func (o Option[T]) OkOr(err error) Result[T] {
	if o.IsNone() {
		return Err[T](err)
	}
	return Ok[T](o.ValueOrZero())
}

// Convert to Result using specified function's result as Err() on None()
func (o Option[T]) OkOrElse(f func() error) Result[T] {
	if o.IsNone() {
		return Err[T](f())
	}
	return Ok[T](o.ValueOrZero())
}

// And return Some(other) if both Option are Some()
// Or return None() if one of them is None()
func (o Option[T]) And(other Option[any]) Option[any] {
	if o.IsNone() {
		return None[any]()
	}
	if other.IsNone() {
		return None[any]()
	}
	return Some[any](other.ValueOrZero())
}

// And return Some(other) if both Option are Some()
// Or return None() if one of them is None()
func (o Option[T]) AndT(other Option[T]) Option[T] {
	if o.IsNone() {
		return None[T]()
	}
	if other.IsNone() {
		return None[T]()
	}
	return Some[T](other.ValueOrZero())
}

// Or return Some(self) if self is Some
// Otherwise return None() or Some(other) depending on state of other
func (o Option[T]) Or(other Option[any]) Option[any] {
	if o.IsNone() {
		return other.Clone()
	}
	return Some[any](o.ValueOrZero())
}

// Or return Some(self) if self is Some
// Otherwise return None() or Some(other) depending on state of other
func (o Option[T]) OrT(other Option[T]) Option[T] {
	if o.IsNone() {
		return other.Clone()
	}
	return Some[T](o.ValueOrZero())
}

// Xor return Some() if exactly one of Option is Some
// Otherwise return None()
func (o Option[T]) Xor(other Option[any]) Option[any] {
	if o.IsSome() && other.IsNone() {
		return Some[any](o.ValueOrZero())
	}
	if o.IsNone() && other.IsSome() {
		return Some[any](other.ValueOrZero())
	}
	return None[any]()
}

// Xor return Some() if exactly one of Option is Some
// Otherwise return None()
func (o Option[T]) XorT(other Option[T]) Option[T] {
	if o.IsSome() && other.IsNone() {
		return Some[T](o.ValueOrZero())
	}
	if o.IsNone() && other.IsSome() {
		return Some[T](other.ValueOrZero())
	}
	return None[T]()
}

// Cast to ThreadSafeObject
func (o Option[T]) ToThreadSafeObject() ThreadSafeObject[Option[any]] {
	return MakeThreadSafeObject(o.ToOption())
}

// Cast to ThreadSafeObject
func (o Option[T]) ToThreadSafeObjectT() ThreadSafeObject[Option[T]] {
	return MakeThreadSafeObject(o)
}
