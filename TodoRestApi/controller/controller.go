package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"todoRestApi/database"
	"todoRestApi/models"
)

func GetTodos(c echo.Context) error {
	var todos []models.Todo             //slice for data
	database.DB.Find(&todos)            //find all
	return c.JSON(http.StatusOK, todos) //return all data
}

func CreateTodo(c echo.Context) error {
	var todo models.Todo
	var err error
	if err = c.Bind(&todo); err != nil { //bind data to struct
		return err
	}

	database.DB.Create(&todo)               //create new todo
	return c.JSON(http.StatusCreated, todo) //return new data
}

func UpdateTodo(c echo.Context) error {
	var todo models.Todo
	if err := c.Bind(&todo); err != nil { //bind data to struct
		return err
	}
	database.DB.Save(&todo)                 //new data by id
	return c.JSON(http.StatusCreated, todo) //return new data
}

func DeleteTodo(c echo.Context) error {
	id := c.Param("id")                    //get id parameter
	database.DB.Delete(&models.Todo{}, id) //delete todo by id
	return c.NoContent(http.StatusNoContent)
}
