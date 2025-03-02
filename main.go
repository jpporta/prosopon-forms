package main

import (
	"github.com/jpporta/prosopon-form/internal"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.GET("/api/component", internal.GetComponent)
	e.POST("/api/login", internal.Login)
	e.Static("/", "static")
	e.Static("/css", "css")
	e.Logger.Fatal(e.Start(":8000"))
}
