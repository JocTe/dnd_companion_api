package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
				type Response struct {
					Message string
				}
        return c.JSON(http.StatusOK, Response{"Hello world!"})
    })
    e.Logger.Fatal(e.Start(":1323"))
}
