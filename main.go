package main

import (
	"todo-list/database"
	"todo-list/todos"

	"github.com/labstack/echo/v4"
)

func main() {
	db := database.InitDb()
	err := db.Ping()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	e.GET("/todos", todos.GetTodosHandler(db))
	e.POST("/todos", todos.CreateTodosHandler(db))
	e.DELETE("/todos/:id", todos.DeleteTodosHandler(db))
	e.PATCH("/todos/:id", todos.TodosUpdateHandler(db))
	e.PATCH("/todos/:id/check", todos.CheckUncheckHandler(db))

	e.Start(":8080")
}
