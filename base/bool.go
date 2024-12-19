package base

import "github.com/PAiP-Web-Language/stdlib-go/dev"

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
