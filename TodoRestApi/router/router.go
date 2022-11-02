package router

import (
	"github.com/labstack/echo/v4"
	"todoRestApi/controller"
)

func SetRouter(e *echo.Echo) {
	e.POST("/todo/", controller.CreateTodo)
	e.GET("/todo/", controller.GetTodos)
	e.PUT("/todo/:id", controller.UpdateTodo)
	e.DELETE("/todo/:id", controller.DeleteTodo)
}
