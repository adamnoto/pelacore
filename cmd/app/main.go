package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/saveav/pelacore/renderings"
)

func views(path string) string {
	return "views/pelacore-ui/build/" + path
}

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Static("/static", views("static"))
	e.File("/", views("index.html"))
	e.File("/manifest.json", views("manifest.json"))
	e.File("/favicon.ico", views("favicon.ico"))

	e.Renderer = renderings.Renderer

	e.Logger.Fatal(e.Start(":8888"))
}
