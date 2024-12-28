package base

import (
	"strings"

	"github.com/PAiP-Web-Language/stdlib-go/dev"
)

// Constants

const (
	// True
	True  = Bool(true)
	// False
	False = Bool(false)
	// Representation of true as string
	trueString = "true"
	// Representation of false as string
	falseString = "false"
)

// Boolean type that is used to represent true and false values
type Bool bool

// Interface

type ToBoolInterface interface {
	ToBool() Bool
}

// Operators

// Or operation
func (b Bool) Or(b2 Bool) Bool {
	return b || b2
}

// And operation
func (b Bool) And(b2 Bool) Bool {
	return b && b2
}

// Xor operation
func (b Bool) Xor(b2 Bool) Bool {
	return b != b2
}

// Xnor operation
func (b Bool) Xnor(b2 Bool) Bool {
	return b == b2
}

// Not operation
func (b Bool) Not() Bool {
	return !b
}

// Debug method is dumping value to stderr and returns itself (so you can chain it if needed)
// This method is intended to be used for debugging purposes
// This method uses value for pointer receiver version check [Bool.DebugP]
func (b Bool) Debug() Bool {
	dev.DebugHelper(1, b)
	return b
}

// Debug method is dumping value to stderr and returns itself (so you can chain it if needed)
// This method is intended to be used for debugging purposes
// This method uses pointer for value receiver version check [Bool.Debug]
func (b *Bool) DebugP() *Bool {
	dev.DebugHelper(1, b)
	return b
}

// Additional methods

// ThenSome returns Some(val) if b is true, None() otherwise
// This returns Option[any]
// If you need typed version use [Bool_ThenSome]
func (b Bool) ThenSome(val any) Option[any] {
	if b {
		return Some[any](val)
	}
	return None[any]()
}

// Then returns Some(f()) if b is true, None otherwise
// This returns Option[any]
// If you need typed version use [Bool_Then]
func (b Bool) Then(f func() any) Option[any] {
	if b {
		return Some[any](f())
	}
	return None[any]()
}

// ThenSome returns Some(val) if b is true, None otherwise
// This returns Option[T]
// If you need untyped version you can use [Bool.ThenSome]
func Bool_ThenSome[T any](b Bool, val T) Option[T] {
	dev.MetaGenericMethod[Bool]()
	if b {
		return Some[T](val)
	}
	return None[T]()
}

// ThenSome returns Some(f()) if b is true, None otherwise
// This returns Option[T]
// If you need untyped version you can use [Bool.Then]
func Bool_Then[T any](b Bool, f func() T) Option[T] {
	dev.MetaGenericMethod[Bool]()
	if b {
		return Some[T](f())
	}
	return None[T]()
}

// From and To Bool

// FromBool creates new Bool struct from specifed boolean variable
func (b Bool) FromBool(b2 bool) Bool {
	return Bool(b2)
}

// FromString creates new Bool struct from specifed string
// If string is invalid value it will assume that it is false
func (b Bool) FromString(s string) Bool {
	sl := strings.ToLower(s)
	if sl == trueString {
		return True
	}
	return False
}

// FromStringInplace changes value of this Bool struct to specified string
// If string is invalid value it will assume that it is false
func (b *Bool) FromStringInplace(s string) {
	sl := strings.ToLower(s)
	if sl == trueString {
		*b = True
	}
	*b = False
}

// FromStringToOption creates new Option[Bool] struct from specifed string
// If string is valid value it will return Some[Bool] with corresponding Bool value
// If string is not valid value it will return None[Bool]
func (b Bool) FromStringToOption(s string) Option[Bool] {
	sl := strings.ToLower(s)
	if sl == trueString {
		return Some[Bool](True)
	}
	if sl == falseString {
		return Some[Bool](False)
	}
	return None[Bool]()
}

// ToBool converts Bool struct to go boolean
func (b Bool) ToBool() bool {
	return bool(b)
}

// String representation

func (b Bool) ToString() string {
	if b {
		return trueString
	}
	return falseString
}

func (b Bool) String() string {
	if b {
		return trueString
	}
	return falseString
}

func (b Bool) GoString() string {
	if b {
		return trueString
	}
	return falseString
}

func (b Bool) ToStr() Str {
	if b {
		return trueString
	}
	return falseString
}

// TemplateType

func (b Bool) Get() Bool {
	return b
}

func (b Bool) GetValue() Bool {
	return b
}

func (b Bool) GetAny() any {
	return b
}

func (b Bool) GetValueAny() any {
	return b
}

func (b *Bool) Set(b2 bool) {
	*b = Bool(b2)
}

func (b *Bool) SetValue(b2 bool) {
	*b = Bool(b2)
}

func (b *Bool) SetAny(b2 any) {
	switch b2v := b2.(type) {
	case Bool:
		*b = b2v
	case bool:
		*b = Bool(b2v)
	case Str:
		b.FromStringInplace(string(b2v))
	case string:
		b.FromStringInplace(b2v)
	case I8:
		*b = Bool(b2v != 0)
	case I16:
		*b = Bool(b2v != 0)
	case I32:
		*b = Bool(b2v != 0)
	case I64:
		*b = Bool(b2v != 0)
	case U8:
		*b = Bool(b2v != 0)
	case U16:
		*b = Bool(b2v != 0)
	case U32:
		*b = Bool(b2v != 0)
	case U64:
		*b = Bool(b2v != 0)
	case Usize:
		*b = Bool(b2v != 0)
	case F32:
		*b = Bool(b2v != 0)
	case F64:
		*b = Bool(b2v != 0)
	case Complex64:
		*b = Bool(b2v != 0)
	case Complex128:
		*b = Bool(b2v != 0)
	case int:
		*b = Bool(b2v != 0)
	case int8:
		*b = Bool(b2v != 0)
	case int16:
		*b = Bool(b2v != 0)
	case int32:
		*b = Bool(b2v != 0)
	case int64:
		*b = Bool(b2v != 0)
	case uint:
		*b = Bool(b2v != 0)
	case uint8:
		*b = Bool(b2v != 0)
	case uint16:
		*b = Bool(b2v != 0)
	case uint32:
		*b = Bool(b2v != 0)
	case uint64:
		*b = Bool(b2v != 0)
	case uintptr:
		*b = Bool(b2v != 0)
	case float32:
		*b = Bool(b2v != 0)
	case float64:
		*b = Bool(b2v != 0)
	case complex64:
		*b = Bool(b2v != 0)
	case complex128:
		*b = Bool(b2v != 0)
	case ToBoolInterface:
		*b = b2v.ToBool()
	default:
		panic("Invalid Type")
	}
}

func (b *Bool) SetValueAny(b2 bool) {
	b.SetAny(b2)
}

func (b Bool) Copy() Bool {
	return Bool(b)
}

func (b Bool) Clone() Bool {
	return Bool(b)
}

// Compare

// -1 = < | 0 = = | 1 = >
// (a <=> b) < 0 if a < b,
// (a <=> b) > 0 if a > b,
// (a <=> b) == 0 if a and b are equal/equivalent.
func (b Bool) Compare(b2 Bool) ThreeWayComparable {
	if b == b2 {
		return ThreeWayComparable__Equal
	}
	if b {
		return ThreeWayComparable__GreaterThan
	}
	return ThreeWayComparable__LessThan
}

func (b Bool) Equal(b2 Bool) bool {
	return b == b2
}

func (b Bool) NotEqual(b2 Bool) bool {
	return b != b2
}

func (b Bool) LessThan(b2 Bool) bool {
	if b == b2 {
		return false
	}
	if b {
		return true
	}
	return false
}

func (b Bool) LessThanOrEqual(b2 Bool) bool {
	if b == b2 {
		return true
	}
	return !bool(b)
}

func (b Bool) GreaterThan(b2 Bool) bool {
	if b == b2 {
		return false
	}
	if !b {
		return true
	}
	return false
}

func (b Bool) GreaterThanOrEqual(b2 Bool) bool {
	if b == b2 {
		return true
	}
	return bool(b)
}

// Casts

func (b Bool) ToI8() I8 {
	if b {
		return 1
	}
	return 0
}

func (b Bool) ToI16() I16 {
	if b {
		return 1
	}
	return 0
}

func (b Bool) ToI32() I32 {
	if b {
		return 1
	}
	return 0
}

func (b Bool) ToI64() I64 {
	if b {
		return 1
	}
	return 0
}

func (b Bool) ToU8() U8 {
	if b {
		return 1
	}
	return 0
}

func (b Bool) ToU16() U16 {
	if b {
		return 1
	}
	return 0
}

func (b Bool) ToU32() U32 {
	if b {
		return 1
	}
	return 0
}

func (b Bool) ToU64() U64 {
	if b {
		return 1
	}
	return 0
}

func (b Bool) ToUsize() Usize {
	if b {
		return 1
	}
	return 0
}

func (b Bool) ToF32() F32 {
	if b {
		return 1
	}
	return 0
}

func (b Bool) ToF64() F64 {
	if b {
		return 1
	}
	return 0
}

func (b Bool) ToComplex64() Complex64 {
	if b {
		return 1
	}
	return 0
}

func (b Bool) ToComplex128() Complex128 {
	if b {
		return 1
	}
	return 0
}

func (b *Bool) ToPointer() *Bool {
	return b
}

func (b Bool) ToValue() Bool {
	return b
}
