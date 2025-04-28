package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type redisSingleProvider struct {
	redisClient redis.Cmdable
}

func (r *redisSingleProvider) Get(ctx context.Context, key string) (string, error) {
	valueRedis := r.redisClient.Get(ctx, key)
	if valueRedis.Err() != nil {
		return "", valueRedis.Err()
	}
	return valueRedis.Val(), nil
}

func (r *redisSingleProvider) HMGet(ctx context.Context, key string, fields []string) (
	[]interface{}, error) {
	valueRedis := r.redisClient.HMGet(ctx, key, fields...)
	if valueRedis.Err() != nil {
		return nil, valueRedis.Err()
	}
	return valueRedis.Val(), nil
}

func (r *redisSingleProvider) Set(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	err := r.redisClient.Set(ctx, key, value, duration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *redisSingleProvider) HSet(ctx context.Context, key string, value map[string]interface{}) error {
	err := r.redisClient.HSet(ctx, key, value).Err()
	if err != nil {
		return err
	}
	return nil
}

type Config struct {
	Address   string
	Port      int
	Password  string
	DefaultDB int
}

func NewRedisSingleProvider(config Config) (IRedisProvider, func()) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Address, config.Port),
		Password: config.Password,  // no password set
		DB:       config.DefaultDB, // use default DB
	})
	return &redisSingleProvider{
			redisClient: rdb,
		}, func() {
			_ = rdb.Close()
		}
}

type IRedisProvider interface {
	Get(ctx context.Context, key string) (string, error)
	HMGet(ctx context.Context, key string, fields []string) (
		[]interface{}, error)
	Set(ctx context.Context, key string, value interface{}, duration time.Duration) error
	HSet(ctx context.Context, key string, value map[string]interface{}) error
}
