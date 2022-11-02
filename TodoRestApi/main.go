package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"todoRestApi/database"
	"todoRestApi/router"
)

func main() {
	e := echo.New()

	database.ConnectDatabase() // initialize and configure database connection

	e.GET("/", func(c echo.Context) error { //welcome page
		return c.String(http.StatusOK, "Welcome to REST API")
	})
	router.SetRouter(e) //set urls of api
	e.Logger.Fatal(e.Start(":1323"))
}
