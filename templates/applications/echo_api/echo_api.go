package main

import "github.com/labstack/echo/v4"

type BaseApiApp struct {
}

func (b *BaseApiApp) Ping(c echo.Context) error {
	return c.String(200, "pong")
}

func NewBaseApiApp() BaseApiApp {
	return BaseApiApp{}
}

func main() {
	e := echo.New()
	app, clean, err := NewApiApp()
	if err != nil {
		panic(err)
	}
	defer clean()
	RegisterPingRoute(e.Group("ping"), app)
	e.Logger.Fatal(e.Start(":8082"))
}

func RegisterPingRoute(g *echo.Group, app BaseApiApp) {
	g.GET("", app.Ping)
}
