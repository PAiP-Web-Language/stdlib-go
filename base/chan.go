package base

type Channel[T any] struct {
	Data chan T
}

type ReadChannel[T any] struct {
	Data <-chan T
}

type WriteChannel[T any] struct {
	Data chan<- T
}
