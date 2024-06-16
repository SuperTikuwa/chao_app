package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/supertikuwa/chao_app/api/handler"
)

func main() {
	e := echo.New()

	e.GET("hc", func(c echo.Context) error {
		return c.String(200, "healthy")
	})

	e.GET("roll", handler.Roll)
	fmt.Println("Starting server on port 80")
	e.Logger.Fatal(e.Start(":80"))
}
