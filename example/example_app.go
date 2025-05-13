package main

import (
	"github.com/s-platform/gore/dsa/tree"
)

var counter int32

func main() {
	tree.ExampleSerialize()
}

func Exam1() {
	arr := make([]int, 0)
	for i := 0; i < 10000000; i++ {
		arr = append(arr, i)
	}
}

func Exam2() {
	arr := make([]int, 0, 10000000)
	for i := 0; i < 10000000; i++ {
		arr = append(arr, i)
	}
}
