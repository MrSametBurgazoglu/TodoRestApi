package router

import (
	"github.com/labstack/echo/v4"
	"todoRestApi/controller"
)

func SetRouter(e *echo.Echo) {
	e.POST("/todo/", controller.CreateTodo)      //create new todo
	e.GET("/todo/", controller.GetTodos)         //get all todos
	e.PUT("/todo/:id", controller.UpdateTodo)    //update todo with id
	e.DELETE("/todo/:id", controller.DeleteTodo) //delete todo with id
}
