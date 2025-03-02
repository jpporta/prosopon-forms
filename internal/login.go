package internal

import (
	"log"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type LoginBody struct{}

func Login(c echo.Context) error {
	email := c.FormValue("email")
	pass := c.FormValue("password")
	log.Println(email, pass)
	return nil
}

func encrypt(str string) string {
	password := []byte(str)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}
func compare(str string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(str), []byte(password))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
