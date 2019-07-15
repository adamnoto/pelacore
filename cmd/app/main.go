package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	e.Logger.Fatal(e.Start(":8888"))
}
