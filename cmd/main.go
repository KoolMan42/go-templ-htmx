package main

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/koolman42/go-templ-htmx/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()
	if os.Getenv("env") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("failed to load env", err)
		}
	}

	db, err := sqlx.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	IndexHandler := handler.IndexHandler{}
	TodosHandler := handler.TodosHandler{
		Db: db,
	}
	// Routes
	e.GET("/", IndexHandler.HandleIndexShow)
	e.GET("/todos", TodosHandler.HandleTodosShow)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

