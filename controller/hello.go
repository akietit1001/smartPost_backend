package controller

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"smartPOST/models"
)

func Hello(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	admin := claims["admin"].(bool)
	message := fmt.Sprintf("Hello %s is admin %v", email, admin)
	return c.JSON(http.StatusOK, &models.Content{
		Text: message,
	})
}
