package base

// Iterator is an iterator over sequences of individual values.
// When called as seq(yield), seq calls yield(v) for each value v in the sequence,
// stopping early if yield returns false.
type Iterator[V any] func(yield func(V) bool)

// IteratorKV2 is an iterator over sequences of pairs of values, most commonly key-value pairs.
// When called as seq(yield), seq calls yield(k, v) for each pair (k, v) in the sequence,
// stopping early if yield returns false.
type IteratorKV[K, V any] func(yield func(K, V) bool)
