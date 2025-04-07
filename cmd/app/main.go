package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"test-task-echo/internal/handler"
	"test-task-echo/internal/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.RoleCheckMiddleware())
	e.GET("/time", handler.TimeHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
