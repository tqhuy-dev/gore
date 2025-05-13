package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/redis/go-redis/v9"
	"github.com/tqhuy-dev/gore/go_monitor/openmetry"
	"net/http"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewFailoverClusterClient(&redis.FailoverOptions{
		MasterName:     "sentinel-cluster",
		SentinelAddrs:  []string{"redis-sentinel1:26379", "redis-sentinel2:26379", "redis-sentinel3:26379"},
		RouteByLatency: true,
	})
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(openmetry.PrometheusMiddleware)
	key := "example"
	e.GET("/read", func(c echo.Context) error {
		value := rdb.Get(ctx, key).Val()
		return c.JSON(http.StatusOK, map[string]interface{}{
			"value": value,
		})
	})
	e.GET("/write/:data", func(c echo.Context) error {
		data := c.Param("data")
		err := rdb.Set(ctx, key, data, 0).Err()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data": data,
		})
	})
	e.Logger.Fatal(e.Start(":8081"))
}
