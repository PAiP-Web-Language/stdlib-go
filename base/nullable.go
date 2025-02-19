package base

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/PAiP-Web-Language/stdlib-go/errors"
)

// Check for case of nil value interface
func isNilString[T any](n T) (result bool) {
	defer func() {
		if r := recover(); r != nil {
			result = true
		}
	}()
	return fmt.Sprintf("%v", n) == fmt.Sprintf("%v", nil)
}

// Nullable represents data that also can be NULL
type Nullable[T any] struct {
	Data  T
	Valid bool
}

// NullableValue Create a Nullable from a value
func NullableValue[T any](value T) Nullable[T] {
	if any(value) == nil {
		return Nullable[T]{Valid: false}
	}
	if isNilString(value) {
		return Nullable[T]{Valid: false}
	}
	switch any(value).(type) {
	case errors.NullError:
		return Nullable[T]{Valid: false}
	default:
		return Nullable[T]{Data: value, Valid: true}
	}
}

// NullableValueFromPointer Create a Nullable from a pointer
func NullableValueFromPointer[T any](value *T) Nullable[T] {
	if value == nil {
		return Nullable[T]{Valid: false}
	}
	if isNilString(*value) {
		return Nullable[T]{Valid: false}
	}
	return NullableValue(*value)
}

// Null Create a Nullable that is NULL with type T
func Null[T any]() Nullable[T] {
	return Nullable[T]{}
}

// NullableZeroValue Get default zero value of type T
func NullableZeroValue[T any]() T {
	return Null[T]().ValueOrZero()
}

// ValueOrZero Get Value, or default zero value if it is NULL
func (n Nullable[T]) ValueOrZero() T {
	if !n.Valid {
		var ref T
		return ref
	}
	return n.Data
}

// IsZero Check if value is zero
func (n Nullable[T]) IsZero() bool {
	if !n.Valid {
		return true
	}
	var ref T
	return any(ref) == any(n.Data)
}

// Equal Check if this Nullable is equal to another Nullable
func (n Nullable[T]) Equal(other Nullable[T]) bool {
	switch any(n.Data).(type) {
	case time.Time:
		nValue := any(n.Data).(time.Time)
		otherValue := any(other.Data).(time.Time)
		return n.Valid == other.Valid && (!n.Valid || nValue.Equal(otherValue))
	}
	return n.ExactEqual(other)
}

// ExactEqual Check if this Nullable is exact equal to another Nullable, never using intern Equal method to check equality
func (n Nullable[T]) ExactEqual(other Nullable[T]) bool {
	return n.Valid == other.Valid && (!n.Valid || any(n.Data) == any(other.Data))
}

// String Convert value to string
func (n Nullable[T]) String() string {
	return fmt.Sprintf("%s", any(n.Data))
}

// Get Go String
func (n Nullable[T]) GoString() string {
	var ref T
	return fmt.Sprintf(
		"Nullable[%T]{Data:%#v,Valid:%#v}",
		ref,
		n.Data,
		n.Valid,
	)
}

// Get underlying value
func (n Nullable[T]) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Data, nil
}

// Run function fn when value is not NULL
func (n Nullable[T]) AndThen(fn func(T) any) any {
	if !n.Valid {
		return nil
	}
	return fn(n.Data)
}

// Get underlying value or panic if error
func (n Nullable[T]) Unwrap() T {
	if !n.Valid {
		panic(errors.UnwrapError{})
	}
	return n.Data
}

// Get underlying value or panic if error
func (n Nullable[T]) UnwrapOrDefault() T {
	return n.ValueOrZero()
}

// Get underlying error or panic if correct value
func (n Nullable[T]) UnwrapErr() error {
	if n.Valid {
		panic(errors.UnwrapError{})
	}
	return errors.NullableError{}
}

// Get underlying value or return default value if error is found
func (n Nullable[T]) UnwrapOr(defaultVal T) T {
	if !n.Valid {
		return defaultVal
	}
	return n.Data
}

// Unwrap but if error occurs run specified function f instead of panic
// Return result of underlying value (having err if it's found)
func (n Nullable[T]) UnwrapOrErr(f func(Nullable[T])) T {
	if !n.Valid {
		f(n)
	}
	return n.ValueOrZero()
}

// Unwraps as result type for if error checks and run specified function f
// instead of panic
// Return result of underlying value (having err if it's found)
func (n Nullable[T]) UnwrapAsResultOrErr(f func(Nullable[T])) Result[T] {
	if !n.Valid {
		f(n)
		return Err[T](errors.UnwrapError{})
	}
	return Ok[T](n.ValueOrZero())
}

// Unwrap value and error separately (Result -> Go normal returns)
func (n Nullable[T]) UnwrapWithErr() (T, error) {
	var err error = nil
	if !n.Valid {
		err = errors.NullableError{}
	}
	return n.Data, err
}

// Expect correct value if error is found panic with specified message
func (n Nullable[T]) Expect(err any) {
	if !n.Valid {
		panic(err)
	}
}

// Expect error value if error is not found panic with specified message
func (n Nullable[T]) ExpectErr(err any) {
	if n.Valid {
		panic(err)
	}
}

// Check if Object has error
func (n Nullable[T]) IsError() bool {
	return !n.Valid
}

// Get Error if Object has error or nil if not
func (n Nullable[T]) GetError() error {
	if n.Valid {
		return nil
	}
	return errors.NullableError{}
}

// Get value or zero
func (n Nullable[T]) GetT() T {
    return n.ValueOrZero()
}

// Get value or zero
func (n Nullable[T]) GetValueT() T {
    return n.ValueOrZero()
}

// Get value or zero as any
func (n Nullable[T]) Get() any {
    return n.ValueOrZero()
}

// Get value or zero as any
func (n Nullable[T]) GetValue() any {
    return n.ValueOrZero()
}

// Check if value is NULL
func (n Nullable[T]) IsNull() bool {
    return !n.Valid
}

// Check if value is not NULL
func (n Nullable[T]) IsValue() bool {
	return n.Valid
}

// Cast to Pointer
func (n *Nullable[T]) ToPointer() *Nullable[T] {
	return n
}

// Cast to Value
func (n Nullable[T]) ToValue() Nullable[T] {
	return n
}

// Cast to Any Nullable
func (n Nullable[T]) ToNullable() Nullable[any] {
	return Nullable[any]{Data: n.Data, Valid: n.Valid}
}

// Cast to Option
func (n Nullable[T]) ToOption() Option[any] {
	if n.IsNull() {
		return None[any]()
	}
	return Some[any](n.ValueOrZero())
}

// Cast to Option
func (n Nullable[T]) ToOptionT() Option[T] {
	if n.IsNull() {
		return None[T]()
	}
	return Some[T](n.ValueOrZero())
}

// Cast to Observable
func (n Nullable[T]) ToObservable() Observable[Nullable[any]] {
	return MakeObservable(n.ToNullable())
}

// Cast to Observable
func (n Nullable[T]) ToObservableT() Observable[Nullable[T]] {
	return MakeObservable(n)
}

// Clone this object
func (n Nullable[T]) Clone() Nullable[T] {
	return Nullable[T]{Data: n.Data, Valid: n.Valid}
}

// Inspect value of this object
func (n Nullable[T]) Inspect(f func(Nullable[T])) Nullable[T] {
	f(n)
	return n
}

// Inspect value of this objects reference
func (n *Nullable[T]) InspectRef(f func(*Nullable[T])) *Nullable[T] {
	f(n)
	return n
}

// Set value of this nullable to NULL
func (n Nullable[T]) SetNull() {
	var ref T
	n.Valid = false
	n.Data = ref
}

// Convert to Result with specified error as Err() on Null()
func (n Nullable[T]) OkOr(err error) Result[T] {
	if n.IsNull() {
		return Err[T](err)
	}
	return Ok[T](n.ValueOrZero())
}

// Convert to Result using specified function's result as Err() on Null()
func (n Nullable[T]) OkOrElse(f func() error) Result[T] {
	if n.IsNull() {
		return Err[T](f())
	}
	return Ok[T](n.ValueOrZero())
}

// Cast to ThreadSafeObject
func (n Nullable[T]) ToThreadSafeObject() ThreadSafeObject[Nullable[any]] {
	return MakeThreadSafeObject(n.ToNullable())
}

// Cast to ThreadSafeObject
func (n Nullable[T]) ToThreadSafeObjectT() ThreadSafeObject[Nullable[T]] {
	return MakeThreadSafeObject(n)
}

