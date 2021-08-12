package main

import (
	"github.com/5xxxx/echo-pprof"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	// automatically add routers for net/http/pprof
	// e.g. /debug/pprof, /debug/pprof/heap, etc.
	echopprof.Wrap(e)

	// echopprof also plays well with *echo.Group
	// prefix := "/debug/pprof"
	// group := e.Group(prefix)
	// echopprof.WrapGroup(prefix, group)

	err := e.Start(":8080")
	if err != nil {
		panic(err)
	}
}
