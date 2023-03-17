package controller

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"smartPOST/apis"
	"smartPOST/entities"
	"smartPOST/models"
	"time"
)

func Login(c echo.Context) error {
	email := c.Get("email").(string)
	log.Printf("Login with email %s\n", email)
	admin := c.Get("admin").(bool)
	log.Printf("Login with admin %v\n", admin)
	var user = entities.User{}
	for _, x := range apis.ListUsers {
		if x.Email == email {
			user.Id = x.Id
			user.FirstName = x.FirstName
			user.LastName = x.LastName
			user.Email = x.Email
			user.Password = x.Password
			if admin {
				user.Role = "admin"
			} else {
				user.Role = "user"
			}
			user.CreateOn = x.CreateOn
			user.LastLogin = x.LastLogin
		}
	}
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
	//sendEmail.SendEmail()
	return c.JSON(http.StatusOK, &models.LoginResponse{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		CreateOn:  user.CreateOn,
		LastLogin: user.LastLogin,
		Role:      user.Role,
		Token:     t,
		Admin:     admin,
	})
}
