package redis

import (
	"context"
	"testing"
)

func BenchmarkGetRedisSingleValue(b *testing.B) {
	redisProvider, closeFunc := NewRedisSingleProvider(Config{
		Address: "localhost",
		Port:    6379,
	})
	ctx := context.Background()
	defer closeFunc()
	for i := 0; i < b.N; i++ {
		_, _ = redisProvider.Get(ctx, "benchmark")
	}
}

func BenchmarkHMGetRedisSingleValue(b *testing.B) {
	redisProvider, closeFunc := NewRedisSingleProvider(Config{
		Address: "localhost",
		Port:    6379,
	})
	ctx := context.Background()
	defer closeFunc()
	for i := 0; i < b.N; i++ {
		_, _ = redisProvider.HMGet(ctx, "benchmark1", []string{"value"})
	}
}

//BenchmarkGetRedisSingleValue-8   	   10971	    109272 ns/op
//BenchmarkHMGetRedisSingleValue-8   	   11110	    105693 ns/op
