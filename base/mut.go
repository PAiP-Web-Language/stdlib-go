package base

import "github.com/PAiP-Web-Language/stdlib-go/dev"

type Mut[T any] struct {
	t *T
}

// Passthrough method for value
func (m *Mut[T]) T() T {
	return *m.t
}

// Passthrough method for value pointer
func (m *Mut[T]) TP() *T {
	return m.t
}

// Debug method is dumping value to stderr and returns itself (so you can chain it if needed)
// This method is intended to be used for debugging purposes
// This method uses value for pointer receiver version check [Mut.DebugP]
func (m Mut[T]) Debug() Mut[T] {
	dev.DebugHelper(1, m)
	return m
}

// Debug method is dumping value to stderr and returns itself (so you can chain it if needed)
// This method is intended to be used for debugging purposes
// This method uses pointer for value receiver version check [Mut.Debug]
func (m *Mut[T]) DebugP() *Mut[T] {
	dev.DebugHelper(1, m)
	return m
}
