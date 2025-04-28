package dsa

import "fmt"

type Queue[T any] struct {
	data []T
}

func InitQueue[T any]() Queue[T] {
	return Queue[T]{
		data: make([]T, 0),
	}
}

func (q *Queue[T]) Push(element T) {
	q.data = append(q.data, element)
}

func (q *Queue[T]) Pop() T {
	first := q.data[0]
	q.data = q.data[1:]
	return first
}

func (q *Queue[T]) Scan() bool {
	return len(q.data) > 0
}

func ExampleQueue() {
	type Tmp struct {
		Data int
	}

	queue := InitQueue[Tmp]()
	queue.Push(Tmp{Data: 1})
	queue.Push(Tmp{Data: 2})
	queue.Push(Tmp{Data: 3})
	for queue.Scan() {
		fmt.Println(queue.Pop().Data)
	}
}
