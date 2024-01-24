package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/koolman42/go-templ-htmx/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()
err := godotenv.Load()
    if err != nil {
        log.Fatal("failed to load env", err)
    }
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	IndexHandler := handler.IndexHandler{}
	// Routes
	e.GET("/", IndexHandler.HandleIndexShow)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
