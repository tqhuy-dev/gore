package api

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tqhuy-dev/gore/templates/applications/pkg"
	"github.com/tqhuy-dev/gore/templates/applications/pkg/constants"
)

type ISampleService interface {
	SampleMethod() string
}
type BaseApiApp struct {
	sampleService ISampleService
}

func (b *BaseApiApp) Ping(c echo.Context) error {
	return c.String(200, b.sampleService.SampleMethod())
}

func NewBaseApiApp(sampleService ISampleService) BaseApiApp {
	return BaseApiApp{
		sampleService: sampleService,
	}
}

func RunAPIServer() {
	signalCtx, signalCtxStop := pkg.GetSignalCtx()
	defer signalCtxStop()

	baseCtx, baseCtxStop := context.WithCancel(context.Background())
	//-------------------------------------------
	e := echo.New()
	app, clean, err := NewApiApp()
	if err != nil {
		panic(err)
	}
	defer clean()
	//---------------------------------------------
	RegisterPingRoute(e.Group("ping"), app)
	server := http.Server{
		Addr: ":1323",
		BaseContext: func(_ net.Listener) context.Context {
			// do not pass a signalCtx here, we don't want to cancel all ongoing requests immediately
			return baseCtx
		},
		Handler: e,
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()
	//-----------------------------------------------
	<-signalCtx.Done()

	log.Println("shutdown initiated")

	// run server as is with readiness check failing for short time
	time.Sleep(constants.ShutdownDelay)

	// give server shutdown process a deadline
	shutdownCtx, shutdownCtxStop := context.WithTimeout(context.Background(), constants.ServerShutdownTimeout)
	defer shutdownCtxStop()

	errShutdown := server.Shutdown(shutdownCtx)
	// cancel the server BaseContext
	baseCtxStop()
	if errShutdown != nil {
		log.Printf("could not shutdown ongoing requests: %v\n", err)
		time.Sleep(constants.HardShutdownWait)
	}

	log.Println("shutdown complete")

	os.Exit(0)
}

func RegisterPingRoute(g *echo.Group, app BaseApiApp) {
	g.GET("", app.Ping)
}
