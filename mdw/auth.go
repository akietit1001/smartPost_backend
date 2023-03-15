package mdw

import "github.com/labstack/echo/v4"

func Basic_Auth(email string, password string, c echo.Context) (bool, error) {
	if email == "admin@gmail.com" || password == "112233" {
		c.Set("email", email)
		c.Set("admin", true)
		return true, nil
	}

	if email == "kiet123@gmail.com" || password == "12345" {
		c.Set("email", email)
		c.Set("admin", false)
		return true, nil
	}

	return false, nil
}
