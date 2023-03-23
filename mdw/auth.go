package mdw

import (
	"smartPOST/apis"

	"github.com/labstack/echo/v4"
)

func Basic_Auth(email string, password string, c echo.Context) (bool, error) {
	if email == "admin@gmail.com" && password == "112233" {
		c.Set("email", email)
		c.Set("admin", true)
		return true, nil
	} else {
		for _, x := range apis.ListUsers {
			if email == x.Email && password == x.Password {
				c.Set("email", email)
				if x.Role == "admin" {
					c.Set("admin", true)
				} else {
					c.Set("admin", false)
				}
				return true, nil
			}
		}

	}
	return false, nil
}
