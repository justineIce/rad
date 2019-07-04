package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Static("/", "dist")
	e.File("/", "dist/index.html")
	e.Logger.Fatal(e.Start(":8089"))
}
