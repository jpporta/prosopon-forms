package main

import (
	"context"

	"github.com/jpporta/prosopon-form/templates"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// Component

	e.GET("/", func(c echo.Context) error {
		component := templates.Hello("Joao")
		return component.Render(context.Background(), c.Response().Writer)

	})
	e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":8000"))
}
