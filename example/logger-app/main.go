package main

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tqhuy-dev/gore/go_monitor/openmetry"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net"
	"net/http"
)

func initLogger() *zap.Logger {
	udpAddr, err := net.ResolveUDPAddr("udp", "logstash:5000") // Logstash UDP server
	if err != nil {
		fmt.Println("Failed to resolve UDP address:", err)
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println("Failed to connect to UDP server:", err)
	}

	// Create a Zap core that writes logs to UDP
	udpCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(conn), // Send logs to UDP
		zap.InfoLevel,         // Set log level
	)

	return zap.New(udpCore)
}

func main() {
	logger := initLogger() // Logs in JSON format
	defer func() {
		_ = logger.Sync()
	}()
	logger.WithOptions()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(openmetry.PrometheusMiddleware)
	e.GET("/read", func(c echo.Context) error {
		logger.Info("read service",
			zap.String("function", "read_function"),
			zap.String("hub_id", "1111"),
		)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"value": 1,
		})
	})
	e.GET("/write/:data", func(c echo.Context) error {
		data := c.Param("data")
		logger.Info("write service",
			zap.String("function", "write_function"),
			zap.String("value", data),
		)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data": data,
		})
	})
	e.GET("/error", func(c echo.Context) error {
		data := c.Param("data")
		logger.Error("error service",
			zap.String("function", "write_function"),
			zap.String("error", errors.New("Error").Error()),
		)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data": data,
		})
	})
	e.Logger.Fatal(e.Start(":8081"))
}
