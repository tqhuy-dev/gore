package graceful_shutdown

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
)

const (
	shutdownDelay         = 5 * time.Second
	serverShutdownTimeout = 10 * time.Second
	hardShutdownWait      = 5 * time.Second
)

var (
	shutdownInProgress atomic.Bool
)
var globalCount = 0

func RunGracefulShutdown() {
	e := echo.New()
	signalCtx, signalCtxStop := signal.NotifyContext(context.Background(),
		syscall.SIGINT,  // Ctrl+C
		syscall.SIGQUIT, // Ctrl+\
		syscall.SIGTERM, // the normal way to politely ask a program to terminate
	)
	defer signalCtxStop()

	e.GET("/", func(c echo.Context) error {
		if shutdownInProgress.Load() {
			fmt.Println("shutdown in progress")
			return echo.NewHTTPError(http.StatusServiceUnavailable)
		}

		fmt.Println("ok")
		return c.String(http.StatusOK, "ok")
	})

	e.GET("inc", func(c echo.Context) error {
		// simulate some delay
		globalCount++
		return c.String(http.StatusOK, "ok")
	})
	e.GET("long", func(c echo.Context) error {
		// simulate some delay
		time.Sleep(time.Duration(rand.Intn(7000)) * time.Millisecond)
		globalCount++
		return c.String(http.StatusOK, "ok")
	})
	// readiness check
	baseCtx, baseCtxStop := context.WithCancel(context.Background())
	server := http.Server{
		Addr: ":1323",
		BaseContext: func(_ net.Listener) context.Context {
			// do not pass a signalCtx here, we don't want to cancel all ongoing requests immediately
			return baseCtx
		},
		Handler: e,
	}

	// run server in a goroutine
	go func() {
		err := server.ListenAndServe()
		if err != nil {

		}
	}()

	// listen for the interrupt signal
	<-signalCtx.Done()

	log.Println("shutdown initiated")
	shutdownInProgress.Store(true)

	// run server as is with readiness check failing for short time
	time.Sleep(shutdownDelay)

	// give server shutdown process a deadline
	shutdownCtx, shutdownCtxStop := context.WithTimeout(context.Background(), serverShutdownTimeout)
	defer shutdownCtxStop()

	// shutdown server
	err := server.Shutdown(shutdownCtx)
	// cancel the server BaseContext
	baseCtxStop()
	if err != nil {
		log.Printf("could not shutdown ongoing requests: %v\n", err)
		time.Sleep(hardShutdownWait)
	}

	log.Println("shutdown complete")

	os.Exit(0)
}
