package mdw

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func IsAdminMdw(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)

		admin := claims["admin"].(bool)

		if admin {
			next(c)
		}
		return echo.ErrUnauthorized
	}
}
