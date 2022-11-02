package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"todoRestApi/database"
	"todoRestApi/models"
)

func GetTodos(c echo.Context) error {
	var todos []models.Todo
	database.DB.Find(&todos)
	return c.JSON(http.StatusOK, todos)
}

func CreateTodo(c echo.Context) error {
	var todo models.Todo
	var err error
	if err = c.Bind(&todo); err != nil {
		return err
	}

	database.DB.Create(&todo)
	return c.JSON(http.StatusCreated, todo)
}

func UpdateTodo(c echo.Context) error {
	var todo models.Todo
	if err := c.Bind(&todo); err != nil {
		return err
	}
	database.DB.Save(&todo)
	return c.JSON(http.StatusCreated, todo)
}

func DeleteTodo(c echo.Context) error {
	id := c.Param("id")
	database.DB.Delete(&models.Todo{}, id)
	return c.NoContent(http.StatusNoContent)
}
