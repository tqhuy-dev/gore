package utilities

import "sync"

type GorePool[T any] struct {
	Pool sync.Pool
}

func (p *GorePool[T]) Get() T {
	return p.Pool.Get().(T)
}

func NewGorePool[T any](fn func() T) *GorePool[T] {
	return &GorePool[T]{
		Pool: sync.Pool{
			New: func() any {
				return fn()
			},
		},
	}
}

func (p *GorePool[T]) Put(x T) {
	p.Pool.Put(x)
}
