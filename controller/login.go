package controller

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"smartPOST/models"
	"time"
)

func Login(c echo.Context) error {
	email := c.Get("email").(string)

	log.Printf("Login with email %s\n", email)
	admin := c.Get("admin").(bool)
	log.Printf("Login with admin %v\n", admin)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = email
	claims["admin"] = admin
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("mysecretkey"))

	if err != nil {
		log.Printf("Signed token err %v\n", err)
		return err
	}

	return c.JSON(http.StatusOK, &models.LoginResponse{
		Token: t,
	})
}
