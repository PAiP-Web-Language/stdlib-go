package base

type Map[K comparable, V any] struct {
	Data map[K]V
}
