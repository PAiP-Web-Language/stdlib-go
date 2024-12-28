package base

// -1 = < | 0 = = | 1 = >
// (a <=> b) < 0 if a < b,
// (a <=> b) > 0 if a > b,
// (a <=> b) == 0 if a and b are equal/equivalent.
type ThreeWayComparable int8

const (
    ThreeWayComparable__LessThan ThreeWayComparable = -1
    ThreeWayComparable__Equal ThreeWayComparable = 0
    ThreeWayComparable__GreaterThan ThreeWayComparable = 1
)
