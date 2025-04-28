package pool_app

import (
	"fmt"
	"github.com/s-platform/gore/utilities"
	"runtime"
	"sync"
)

var (
	_pool = NewAppPool()
	Get   = _pool.Get
)

type PoolDto struct {
	Dto  Dto
	Pool AppPool
}
type Dto struct {
	Min int
	Max int
}

func (dto *PoolDto) Reset() {
	dto.Dto.Min = 0
	dto.Dto.Max = 0
}

func (dto *PoolDto) Free() {
	dto.Pool.put(dto)
}

type AppPool struct {
	P *utilities.GorePool[*PoolDto]
}

func NewAppPool() AppPool {
	return AppPool{
		P: utilities.NewGorePool(func() *PoolDto {
			return &PoolDto{
				Dto: Dto{
					Min: 0,
					Max: 0,
				},
			}
		}),
	}
}

var m runtime.MemStats

func (p AppPool) Get() *PoolDto {
	value := p.P.Get()
	value.Reset()
	value.Pool = p
	return value
}

func (p AppPool) put(dto *PoolDto) {
	p.P.Put(dto)
}

func worker() {
	poolValue := Get()
	poolValue.Dto.Min = utilities.RandomRange(1, 100)
	poolValue.Dto.Max = utilities.RandomRange(1, 100)
	fmt.Printf("%p\n", poolValue)
	defer poolValue.Free()
}

func workerWithoutPool() {
	poolValue := &Dto{}
	poolValue.Min = utilities.RandomRange(1, 100)
	poolValue.Max = utilities.RandomRange(1, 100)
	fmt.Printf("%p\n", poolValue)
}

func ExamplePool() {
	for i := 0; i < 10; i++ {
		workerWithoutPool()
	}
}

func ExampleWithoutPool() {
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			workerWithoutPool()
		}()
	}
	wg.Wait()
}
