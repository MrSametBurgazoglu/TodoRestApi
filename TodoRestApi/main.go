package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"todoRestApi/database"
	"todoRestApi/router"
)

func main() {
	e := echo.New()

	database.ConnectDatabase()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	router.SetRouter(e)
	e.Logger.Fatal(e.Start(":1323"))
}
