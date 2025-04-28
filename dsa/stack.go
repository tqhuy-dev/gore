package dsa

import "fmt"

type Stack[T any] struct {
	data []T
	top  int
}

func InitStack[T any]() Stack[T] {
	return Stack[T]{
		data: make([]T, 0),
	}
}

func (s *Stack[T]) Push(element T) {
	s.top++
	s.data = append(s.data, element)
}

func (s *Stack[T]) Pop() T {
	s.top--
	last := s.data[s.top]
	s.data = s.data[:s.top]
	return last
}

func (s *Stack[T]) Scan() bool {
	return s.top > 0
}

type ExpData struct {
	Value int
}

func ExampleStack() {
	stack := InitStack[ExpData]()
	stack.Push(ExpData{Value: 1})
	stack.Push(ExpData{Value: 2})
	stack.Push(ExpData{Value: 3})
	for stack.Scan() {
		fmt.Println(stack.Pop().Value)
	}
}
