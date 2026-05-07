package utilities

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// PoolHandle định nghĩa hàm xử lý job
type PoolHandle[Req any, Res any] func(ctx context.Context, req Req) (Res, error)

// PoolOption chứa cấu hình của worker pool
type PoolOption struct {
	Name        string
	WorkerLimit int
	TotalTask   int
	Ctx         context.Context
}

// WorkerPool quản lý workers và job queue
type WorkerPool[Req any, Res any] struct {
	option  PoolOption
	Jobs    chan Req
	Handler PoolHandle[Req, Res]
	Results chan Res
	Errors  chan error
	wg      sync.WaitGroup
	ctx     context.Context
	cancel  context.CancelFunc
}

// NewWorkerPool: Khởi tạo worker pool
func NewWorkerPool[Req any, Res any](option PoolOption, handle PoolHandle[Req, Res], Ctx context.Context) *WorkerPool[Req, Res] {
	ctx, cancel := context.WithCancel(Ctx)
	return &WorkerPool[Req, Res]{
		option:  option,
		Jobs:    make(chan Req, option.TotalTask),
		Handler: handle,
		Results: make(chan Res, option.TotalTask),
		Errors:  make(chan error, option.TotalTask),
		ctx:     ctx,
		cancel:  cancel,
	}
}

// Start: Bắt đầu workers
func (wp *WorkerPool[Req, Res]) Start() {
	for w := 1; w <= wp.option.WorkerLimit; w++ {
		go func(workerId int) {
			wp.worker(workerId)
		}(w)
	}
}

// Push: Đẩy công việc vào queue
func (wp *WorkerPool[Req, Res]) Push(value Req) error {
	select {
	case <-wp.ctx.Done():
		return wp.ctx.Err()
	default:
		wp.wg.Add(1)
		wp.Jobs <- value
		return nil
	}
}

// Close: Đợi workers hoàn thành và đóng channels
func (wp *WorkerPool[Req, Res]) WaitAndClose() {
	wp.wg.Wait()
	close(wp.Jobs)
	close(wp.Results)
	close(wp.Errors)
	wp.cancel()
}

// Worker function: Nhận công việc từ jobs channel và xử lý
func (wp *WorkerPool[Req, Res]) worker(id int) {
	for {
		select {
		case <-wp.ctx.Done():
			return
		case req, ok := <-wp.Jobs:
			if !ok {
				return
			}

			res, err := wp.Handler(wp.ctx, req)

			wp.Results <- res
			wp.Errors <- err

			wp.wg.Done()
		}
	}
}

// Example sử dụng WorkerPool
func Example() {
	option := PoolOption{
		Name:        "example",
		WorkerLimit: 3,
		TotalTask:   10,
	}
	type tmp struct {
		int int
	}

	type tmpRes struct {
		int int
	}

	results := make([]int, 0)
	var resultWg sync.WaitGroup
	resultWg.Add(1)
	// Tạo worker pool với 3 workers
	workerPool := NewWorkerPool[*tmp, *tmpRes](option, func(ctx context.Context, s *tmp) (*tmpRes, error) {
		time.Sleep(2 * time.Second)
		if s.int%2 == 0 {
			return nil, errors.New("tmp error")
		}
		return &tmpRes{int: s.int * 2}, nil
	}, context.Background())

	workerPool.Start()

	// Thêm 10 công việc
	for i := 0; i < option.TotalTask; i++ {
		if err := workerPool.Push(&tmp{int: i}); err != nil {
			fmt.Println("Push error:", err)
		}
	}

	go func() {
		defer resultWg.Done()
		for r := range workerPool.Results {
			if r == nil {
				continue
			}
			results = append(results, r.int)
		}
	}()

	go func() {
		for e := range workerPool.Errors {
			if e != nil {
				fmt.Println("Error:", e)
			}
		}
	}()

	// Đóng worker pool
	workerPool.WaitAndClose()
	resultWg.Wait()
	fmt.Println("All jobs completed. Run next process")
	fmt.Println("Results:", results)
}
