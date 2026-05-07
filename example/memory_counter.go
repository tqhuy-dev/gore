package example

import (
	"fmt"
	"sync"
)

var InMemoryCounter sync.Map
var flush chan interface{}

func do(key string, increment int, wg *sync.WaitGroup) {
	value, isLoaded := InMemoryCounter.LoadOrStore(key, increment)
	if isLoaded {
		newValue := value.(int) + increment
		if !InMemoryCounter.CompareAndSwap(key, value, newValue) {
			do(key, increment, wg)
		} else {
			wg.Done()
		}
	} else {
		wg.Done()
	}
}

func flushKey() {
	InMemoryCounter.Range(func(key, value interface{}) bool {
		currentValue, _ := InMemoryCounter.LoadAndDelete(key)
		fmt.Println(currentValue)
		return true
	})
}
