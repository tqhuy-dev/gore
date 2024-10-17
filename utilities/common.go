package utilities

func Empty[T any]() T {
	var zero T
	return zero
}

type Clonable[T any] interface {
	Clone() T
}

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}
