package base

import (
	"github.com/PAiP-Web-Language/stdlib-go/dev"
	"github.com/PAiP-Web-Language/stdlib-go/todo"
)

// TODO: Types
type Iterator[T any] interface {}
type ThreadSafeObject[T any] interface {}
type Arc[T any] interface {}
type Rc[T any] interface {}
type Mut[T any] interface {}
// -1 = < | 0 = = | 1 = >
// (a <=> b) < 0 if a < b,
// (a <=> b) > 0 if a > b,
// (a <=> b) == 0 if a and b are equal/equivalent.
type ThreeWayComparable uint8

type TemplateType interface {
	// Array Like

	// Length
	Len() LUI10[int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64]
	LenResult() Result[LUI10[int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64]]
	// Contains
	Contains(any) bool
	// Insert
	Insert(any) dev.Self
	InsertInplace(any)
	// Append
	Append(any) dev.Self
	AppendInplace(any)
	// Prepend
	Prepend(any) dev.Self
	PrependInplace(any)

	// Cast
	ToResult() Result[any]
	ToNullable() Nullable[any]
	ToOption() Option[any]
	ToIter() Iterator[any]
	ToThreadSafeObject() ThreadSafeObject[any]
	ToArc() Arc[any]
	ToRc() Rc[any]
	ToMut() Mut[any]
	ToObservable() Observable[any]
	ToOnceable() Onceable[any]
	ToArray() []any
	ToSlice() []any
	ToMap() map[any]any
	ToPointer() *any
	ToValue() any
	ToString() string
	String() string
	GoString() string
	ToX() todo.TodoInterface

	// Compare
	Compare(any) ThreeWayComparable
	Equal(any) bool
	NotEqual(any) bool
	LessThan(any) bool
	LessThanOrEqual(any) bool
	GreaterThan(any) bool
	GreaterThanOrEqual(any) bool

	// Copy
	Copy() dev.Self
	Clone() dev.Self
	Move() dev.Self

	// Elements
	View() string
	Render() string

	// Get
	Get() any
	GetValue() any

	// Set
	Set(any)
	SetValue(any)

	// Math
	Add(any) any
	Sub(any) any
	Mul(any) any
	Div(any) any
	Mod(any) any
	AddInplace(any) any
	SubInplace(any) any
	MulInplace(any) any
	DivInplace(any) any
	ModInplace(any) any

	// From
	FromString(string) any
	FromStringInplace(string)
	FromIter(Iterator[any]) any
	FromIterInplace(Iterator[any])
	FromMut(Mut[any]) any
	FromMutInplace(Mut[any])
	FromX(todo.TodoInterface) any
	FromXInplace(todo.TodoInterface)

	// To Tuple
	ToTuple1() (any)
	ToTuple2() (any, any)
	ToTuple3() (any, any, any)
	ToTuple4() (any, any, any, any)
	ToTuple5() (any, any, any, any, any)
	ToTuple6() (any, any, any, any, any, any)
	ToTuple7() (any, any, any, any, any, any, any)
	ToTuple8() (any, any, any, any, any, any, any, any)
	ToTuple9() (any, any, any, any, any, any, any, any, any)
	ToTuple10() (any, any, any, any, any, any, any, any, any, any)
	ToTuple11() (any, any, any, any, any, any, any, any, any, any, any)
	ToTuple12() (any, any, any, any, any, any, any, any, any, any, any, any)
	ToTuple13() (any, any, any, any, any, any, any, any, any, any, any, any, any)
	ToTuple14() (any, any, any, any, any, any, any, any, any, any, any, any, any, any)
	ToTuple15() (any, any, any, any, any, any, any, any, any, any, any, any, any, any, any)
	ToTuple16() (any, any, any, any, any, any, any, any, any, any, any, any, any, any, any, any)
	ToTuple17() (any, any, any, any, any, any, any, any, any, any, any, any, any, any, any, any, any)
	ToTuple18() (any, any, any, any, any, any, any, any, any, any, any, any, any, any, any, any, any, any)
	ToTuple19() (any, any, any, any, any, any, any, any, any, any, any, any, any, any, any, any, any, any, any)
	ToTuple20() (any, any, any, any, any, any, any, any, any, any, any, any, any, any, any, any, any, any, any, any)

	// To Fixed Array
	ToFixedArray1() ([1]any)
	ToFixedArray2() ([2]any)
	ToFixedArray3() ([3]any)
	ToFixedArray4() ([4]any)
	ToFixedArray5() ([5]any)
	ToFixedArray6() ([6]any)
	ToFixedArray7() ([7]any)
	ToFixedArray8() ([8]any)
	ToFixedArray9() ([9]any)
	ToFixedArray10() ([10]any)
	ToFixedArray11() ([11]any)
	ToFixedArray12() ([12]any)
	ToFixedArray13() ([13]any)
	ToFixedArray14() ([14]any)
	ToFixedArray15() ([15]any)
	ToFixedArray16() ([16]any)
	ToFixedArray17() ([17]any)
	ToFixedArray18() ([18]any)
	ToFixedArray19() ([19]any)
	ToFixedArray20() ([20]any)
}

type TypedTemplateType[Self any, T any, T2 any, TC comparable] interface {
	// Array Like

	// Contains
	ContainsT(T) bool
	// Insert
	InsertT(T) Self
	InsertInplaceT(T)
	// Append
	AppendT(T) Self
	AppendInplaceT(T)
	// Prepend
	PrependT(T) Self
	PrependInplaceT(T)

	// Cast
	ToResultT() Result[T]
	ToNullableT() Nullable[T]
	ToOptionT() Option[T]
	ToIterT() Iterator[T]
	ToThreadSafeObjectT() ThreadSafeObject[T]
	ToArcT() Arc[T]
	ToRcT() Rc[T]
	ToMutT() Mut[T]
	ToObservableT() Observable[T]
	ToOnceableT() Onceable[T]
	ToArrayT() []T
	ToSliceT() []T
	ToMapT() map[TC]T
	ToPointerT() *T
	ToValueT() T
	ToPointer() *Self
	ToValue() Self
	ToX() todo.TodoInterface

	// Compare
	Compare(T) ThreeWayComparable
	Equal(T) bool
	NotEqual(T) bool
	LessThan(T) bool
	LessThanOrEqual(T) bool
	GreaterThan(T) bool
	GreaterThanOrEqual(T) bool

	// Copy
	Copy() dev.Self
	Clone() dev.Self
	Move() dev.Self

	// Elements
	View() string
	Render() string

	// Get
	Get() T
	GetValue() T

	// Set
	Set(T)
	SetValue(T)

	// Math
	Add(T) T
	Sub(T) T
	Mul(T) T
	Div(T) T
	Mod(T) T
	AddInplace(T) T
	SubInplace(T) T
	MulInplace(T) T
	DivInplace(T) T
	ModInplace(T) T

	// From
	FromString(string) T
	FromStringInplace(string)
	FromIter(Iterator[T]) T
	FromIterInplace(Iterator[T])
	FromMut(Mut[T]) T
	FromMutInplace(Mut[T])
	FromX(todo.TodoInterface) any
	FromXInplace(todo.TodoInterface)

	// To Tuple
	ToTuple1() (T)
	ToTuple2() (T, T)
	ToTuple3() (T, T, T)
	ToTuple4() (T, T, T, T)
	ToTuple5() (T, T, T, T, T)
	ToTuple6() (T, T, T, T, T, T)
	ToTuple7() (T, T, T, T, T, T, T)
	ToTuple8() (T, T, T, T, T, T, T, T)
	ToTuple9() (T, T, T, T, T, T, T, T, T)
	ToTuple10() (T, T, T, T, T, T, T, T, T, T)
	ToTuple11() (T, T, T, T, T, T, T, T, T, T, T)
	ToTuple12() (T, T, T, T, T, T, T, T, T, T, T, T)
	ToTuple13() (T, T, T, T, T, T, T, T, T, T, T, T, T)
	ToTuple14() (T, T, T, T, T, T, T, T, T, T, T, T, T, T)
	ToTuple15() (T, T, T, T, T, T, T, T, T, T, T, T, T, T, T)
	ToTuple16() (T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T)
	ToTuple17() (T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T)
	ToTuple18() (T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T)
	ToTuple19() (T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T)
	ToTuple20() (T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T)

	// To Fixed Array
	ToFixedArray1() ([1]T)
	ToFixedArray2() ([2]T)
	ToFixedArray3() ([3]T)
	ToFixedArray4() ([4]T)
	ToFixedArray5() ([5]T)
	ToFixedArray6() ([6]T)
	ToFixedArray7() ([7]T)
	ToFixedArray8() ([8]T)
	ToFixedArray9() ([9]T)
	ToFixedArray10() ([10]T)
	ToFixedArray11() ([11]T)
	ToFixedArray12() ([12]T)
	ToFixedArray13() ([13]T)
	ToFixedArray14() ([14]T)
	ToFixedArray15() ([15]T)
	ToFixedArray16() ([16]T)
	ToFixedArray17() ([17]T)
	ToFixedArray18() ([18]T)
	ToFixedArray19() ([19]T)
	ToFixedArray20() ([20]T)
}
