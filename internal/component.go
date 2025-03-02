package internal

import (
	"log"

	"github.com/jpporta/prosopon-form/templates"
	"github.com/labstack/echo/v4"
)

func GetComponent(c echo.Context) error {
	token, err := c.Cookie("token")
	log.Println("Request of component")
	if err != nil {
		log.Println("Error getting token cookie", err)
		return CreateLoginComponent(c)
	}
	if token.Valid() != nil {
		return CreateLoginComponent(c)
	}
	return nil
}

func CreateLoginComponent(c echo.Context) error {
	login := templates.LoginForm()
	err := login.Render(c.Request().Context(), c.Response().Writer)
	if err != nil {
		log.Println("Failed to render login component")
	}
	return nil
}
