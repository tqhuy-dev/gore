package utilities

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync"
	"time"
)

type ThreadFunc func(ctx context.Context) (interface{}, error)

type MultiThreadRunner struct {
	threadCount int
	wg          *sync.WaitGroup
	threads     []ThreadFunc
}

type MultiThreadWithChannelResult struct {
	wg *sync.WaitGroup
	ch chan interface{}
}

func NewMultiThreadRunner() *MultiThreadRunner {
	return &MultiThreadRunner{wg: &sync.WaitGroup{}}
}

func (r *MultiThreadRunner) Add(f ThreadFunc) {
	r.threads = append(r.threads, f)
}

func (r *MultiThreadRunner) Run(ctx context.Context) ([]interface{}, []error) {
	r.threadCount = len(r.threads)
	r.wg.Add(r.threadCount)
	values := make([]interface{}, r.threadCount)
	errs := make([]error, r.threadCount)
	for index, thread := range r.threads {
		go func(wg *sync.WaitGroup, thread ThreadFunc, index int) {
			defer wg.Done()
			values[index], errs[index] = thread(ctx)
		}(r.wg, thread, index)
	}
	r.wg.Wait()
	return values, errs
}

func RunWaitGroupWithChannelResult() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(2)
	fmt.Println("Start Run")
	go func(wg *sync.WaitGroup, ch chan<- int) {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		ch <- 1
	}(&wg, ch)

	go func(wg *sync.WaitGroup, ch chan<- int) {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		ch <- 2
	}(&wg, ch)
	go func() {
		wg.Wait()
		close(ch)
	}()
	for n := range ch {
		fmt.Println(n)
	}
	fmt.Println("Run Done")
}

func RunWaitGroupWithMultiChannelAndSelect() {
	var wg sync.WaitGroup
	ch := make(chan int)
	errChan := make(chan error)
	wg.Add(2)
	fmt.Println("Start Run")
	go func(wg *sync.WaitGroup, ch chan<- int, errChan chan<- error) {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		errChan <- errors.New("error 1")
	}(&wg, ch, errChan)

	go func(wg *sync.WaitGroup, ch chan<- int, errChan chan<- error) {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		ch <- 2
	}(&wg, ch, errChan)
	go func() {
		wg.Wait()
		close(ch)
		close(errChan)
	}()
	for {
		select {
		case n, ok := <-ch:
			if !ok {
				ch = nil
			} else {
				fmt.Println(n)
			}
		case e, ok := <-errChan:
			if !ok {
				errChan = nil
			} else {
				fmt.Println(e.Error())
			}
		}
		if ch == nil && errChan == nil {
			break
		}
	}
	fmt.Println("Run Done")
}

func RunWithErrGroup() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	errG, ctx := errgroup.WithContext(ctx)
	var rs1, rs2 int
	fmt.Println("start run")
	errG.Go(func() error {
		time.Sleep(5 * time.Second)
		fmt.Println("Run 1")
		rs1 = 1
		return nil
	})
	errG.Go(func() error {
		time.Sleep(1 * time.Second)
		fmt.Println("Run 2 error")
		return errors.New("error 2")
	})
	if err := errG.Wait(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Run Success")
	}
	fmt.Printf("%d - %d\n", rs1, rs2)
}
func RunWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // Hủy context khi không còn cần thiết

	// Khai báo kết quả
	results := make([]string, 2)

	// Goroutine cho tác vụ đầu tiên (5 giây)
	go func() {
		time.Sleep(5 * time.Second) // Giả lập tác vụ mất 5 giây
		results[0] = "Kết quả tác vụ 1"
	}()

	// Goroutine cho tác vụ thứ hai (1 giây, trả về lỗi)
	go func() {
		time.Sleep(1 * time.Second) // Giả lập tác vụ mất 1 giây
		results[1] = "Kết quả tác vụ 2"
	}()

	// Kiểm tra kết quả sau khi có tín hiệu từ context hoặc goroutines hoàn thành
	select {
	case <-ctx.Done():
		fmt.Println("Thời gian chờ đã hết:", ctx.Err())
	case <-time.After(1 * time.Second):
		if len(results[0]) == 0 {
			results[0] = "Error"
		}
	}

	// Kiểm tra kết quả tác vụ 1
	fmt.Println("Run Done")
	fmt.Println(results)
}
