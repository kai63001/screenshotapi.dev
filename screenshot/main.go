package main

import (
	api "screenshot/api"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/v1/screenshot", func(c echo.Context) error {
		return api.TakeScreenshot(c)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
