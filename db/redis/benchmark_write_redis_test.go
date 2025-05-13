package redis

import (
	"context"
	"encoding/json"
	"github.com/tqhuy-dev/gore/utilities"
	"strconv"
	"testing"
	"time"
)

func BenchmarkSetRedisSingleValue(b *testing.B) {
	redisProvider, closeFunc := NewRedisSingleProvider(Config{
		Address: "localhost",
		Port:    6379,
	})
	ctx := context.Background()
	defer closeFunc()
	for i := 0; i < b.N; i++ {
		randValue := strconv.Itoa(utilities.RandomRange(1, 10000))
		_ = redisProvider.Set(ctx, "benchmark", randValue, 60*time.Second)
	}
}

func BenchmarkSetRedisStructureValue(b *testing.B) {
	redisProvider, closeFunc := NewRedisSingleProvider(Config{
		Address: "localhost",
		Port:    6379,
	})
	ctx := context.Background()
	type tmp struct {
		Value int    `json:"value"`
		Key   string `json:"key"`
	}
	jsonStr, err := json.Marshal(tmp{})
	if err != nil {
		b.Fatal(err)
	}
	defer closeFunc()
	for i := 0; i < b.N; i++ {
		_ = redisProvider.Set(ctx, "benchmark", string(jsonStr), 60*time.Second)
	}
}

func BenchmarkHSetRedisStructureValue(b *testing.B) {
	redisProvider, closeFunc := NewRedisSingleProvider(Config{
		Address: "localhost",
		Port:    6379,
	})
	ctx := context.Background()
	defer closeFunc()
	for i := 0; i < b.N; i++ {
		_ = redisProvider.HSet(ctx, "benchmark1", map[string]interface{}{
			"value": 2,
			"key":   "23234234",
		})
	}
}

func BenchmarkHSetRedisStructureValueSomeField(b *testing.B) {
	redisProvider, closeFunc := NewRedisSingleProvider(Config{
		Address: "localhost",
		Port:    6379,
	})
	ctx := context.Background()
	defer closeFunc()
	for i := 0; i < b.N; i++ {
		_ = redisProvider.HSet(ctx, "benchmark1", map[string]interface{}{
			"key": "23234234",
		})
	}
}

func BenchmarkHSetRedisSingleValue(b *testing.B) {
	redisProvider, closeFunc := NewRedisSingleProvider(Config{
		Address: "localhost",
		Port:    6379,
	})
	ctx := context.Background()
	defer closeFunc()
	for i := 0; i < b.N; i++ {
		randValue := strconv.Itoa(utilities.RandomRange(1, 10000))
		_ = redisProvider.HSet(ctx, "benchmark1", map[string]interface{}{
			"value": randValue,
		})
	}
}

//BenchmarkSetRedisSingleValue-8   	   10388	    108412 ns/op
//BenchmarkHSetRedisSingleValue-8   	   10179	    110771 ns/op
//BenchmarkSetRedisStructureValue-8   	    9914	    115058 ns/op
//BenchmarkHSetRedisStructureValue-8   	    9426	    110206 ns/op
//BenchmarkHSetRedisStructureValueSomeField-8   	    9735	    104195 ns/op
