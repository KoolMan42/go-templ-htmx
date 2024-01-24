package handler

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/koolman42/go-templ-htmx/models"
	"github.com/koolman42/go-templ-htmx/view/todos"
	"github.com/labstack/echo/v4"
)

type TodosHandler struct {
  Db *sqlx.DB
}


func(h TodosHandler) HandleTodosShow(c echo.Context) error {
  todoList := []models.Todo{
		{
			Title:       "Complete Golang tutorial",
			Description: "Finish the tutorial on Go programming language.",
			DueDate:     time.Now().AddDate(0, 0, 7), // Due in 7 days
			IsCompleted: false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Title:       "Read a book",
			Description: "Start reading a new book of your choice.",
			DueDate:     time.Now().AddDate(0, 0, 14), // Due in 14 days
			IsCompleted: false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Title:       "Exercise",
			Description: "Go for a jog or do a workout routine.",
			DueDate:     time.Now().AddDate(0, 0, 2), // Due in 2 days
			IsCompleted: false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}


  return render(c,todos.Show(todoList))
}
