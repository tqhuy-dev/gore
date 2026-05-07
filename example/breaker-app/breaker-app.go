package breaker_app

import (
	"errors"
	"fmt"
	"github.com/sony/gobreaker"
	"math/rand"
	"time"
)

func unstableAPI() (string, error) {
	if rand.Intn(10) < 7 {
		return "", errors.New("simulated error")
	}
	return "Success!", nil
}

func RunBreaker() {
	settings := gobreaker.Settings{
		MaxRequests: 2,
		Name:        "MyCB",
		Timeout:     5 * time.Second, // after this time, circuit tries half-open
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			// Break if error rate > 60%
			return counts.ConsecutiveFailures >= 3
		},
		OnStateChange: func(name string, from, to gobreaker.State) {
			fmt.Printf("Circuit breaker state changed: %s -> %s\n", from.String(), to.String())
		},
	}

	cb := gobreaker.NewCircuitBreaker(settings)

	for i := 0; i < 10; i++ {
		res, err := callWithBreaker(cb, "", mainFunction, fallbackFunction)
		fmt.Printf("Result: %v, Err: %v\n", res, err)
		time.Sleep(1 * time.Second)
	}
}

func callWithBreaker[Req any, Res any](cb *gobreaker.CircuitBreaker, param Req, fn func(Req) (Res, error), fallback func(Req) (Res, error)) (res Res, err error) {
	result, err := cb.Execute(func() (interface{}, error) {
		return fn(param)
	})

	if err != nil {
		if errors.Is(err, gobreaker.ErrOpenState) || errors.Is(err, gobreaker.ErrTooManyRequests) {
			// Fallback khi breaker đang Open
			return fallback(param)
		}
		// Error thực tế trong function
		return
	}
	if val, ok := result.(Res); ok {
		return val, nil
	}
	err = fmt.Errorf("invalid result type: %T", result)
	return
}

func mainFunction(param string) (string, error) {
	if rand.Intn(10) < 7 {
		fmt.Println("Run Fail")
		return "", errors.New("simulated error")
	}
	fmt.Println("Run Success")
	return "Success!", nil
}

func fallbackFunction(param string) (string, error) {
	fmt.Println("✅ Executing fallback logic")
	return "default fallback result", nil
}
