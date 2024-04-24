package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", HomeHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

func HomeHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!1")
}
