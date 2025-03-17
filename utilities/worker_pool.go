package utilities

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// PoolHandle định nghĩa hàm xử lý job
type PoolHandle func(string) error

// PoolOption chứa cấu hình của worker pool
type PoolOption struct {
	Name        string
	WorkerLimit int
	TotalTask   int
}

// ResultPool chứa kết quả xử lý
type ResultPool struct {
	Error error
}

// WorkerPool quản lý workers và job queue
type WorkerPool struct {
	option  PoolOption
	Jobs    chan string
	Handler PoolHandle
	result  chan ResultPool
	wg      sync.WaitGroup
}

// Worker function: Nhận công việc từ jobs channel và xử lý
func (wp *WorkerPool) worker(id int) {
	for j := range wp.Jobs {
		fmt.Println("Worker", id, "processing job:", j)
		err := wp.Handler(j)
		wp.result <- ResultPool{err}
		fmt.Println("Worker", id, "finished job:", j)
		wp.wg.Done() // Đánh dấu công việc đã hoàn thành
	}
}

// NewWorkerPool: Khởi tạo worker pool
func NewWorkerPool(option PoolOption, handle PoolHandle) *WorkerPool {
	return &WorkerPool{
		option:  option,
		Jobs:    make(chan string, option.TotalTask),
		Handler: handle,
		result:  make(chan ResultPool, option.TotalTask),
	}
}

// Start: Bắt đầu workers
func (wp *WorkerPool) Start() {
	for w := 1; w <= wp.option.WorkerLimit; w++ {
		go wp.worker(w)
	}
}

// Push: Đẩy công việc vào queue
func (wp *WorkerPool) Push(value string) {
	wp.wg.Add(1) // Tăng counter trước khi thêm job
	wp.Jobs <- value
}

// Close: Đợi workers hoàn thành và đóng channels
func (wp *WorkerPool) Close() {
	wp.wg.Wait() // Đợi tất cả công việc hoàn thành
	close(wp.Jobs)
	close(wp.result)
}

// Example sử dụng WorkerPool
func Example() {
	option := PoolOption{
		Name:        "example",
		WorkerLimit: 3,
		TotalTask:   10,
	}

	// Tạo worker pool với 3 workers
	workerPool := NewWorkerPool(option, func(s string) error {
		time.Sleep(2 * time.Second)
		fmt.Println("Processing job:", s)
		return errors.New("Processing error")
	})

	workerPool.Start()

	// Thêm 10 công việc
	for i := 0; i < option.TotalTask; i++ {
		workerPool.Push(fmt.Sprintf("Task %d", i))
	}

	// Đóng worker pool
	workerPool.Close()

	fmt.Println("All jobs completed. Run next process")
}
