package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/saveav/pelacore/handlers"
	"github.com/saveav/pelacore/renderings"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	e.Renderer = renderings.Renderer

	e.GET("/", handlers.GetIndex)

	e.Logger.Fatal(e.Start(":8888"))
}
