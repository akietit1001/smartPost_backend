package main

import (
	"net/http"
	"smartPOST/apis"
	"smartPOST/controller"
	"smartPOST/database"
	"smartPOST/mdw"
	"smartPOST/sendEmail"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Create server
	server := echo.New()

	// Connect database
	database.DBConnection()

	// Send email
	sendEmail.SendEmailWithGmail()

	// Middleware
	isLoggedIn := echojwt.JWT([]byte("mysecretkey"))
	isAdmin := mdw.IsAdminMdw
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	//Login
	server.POST("/login", controller.Login, middleware.BasicAuth(mdw.Basic_Auth))

	//Home
	server.GET("/", controller.Hello, isLoggedIn)

	//User
	server.POST("/api/user/create", apis.CreateUser)
	groupUser := server.Group("/api/user", isLoggedIn)
	groupUser.GET("/get/:id", apis.GetUser)
	groupUser.GET("/getall", apis.GetAllUsers)
	groupUser.PUT("/update-name/:id", apis.UpdateNameUser)
	groupUser.PUT("/update-password/:id", apis.UpdatePasswordUser)
	groupUser.PUT("/update-email/:id", apis.UpdateEmailUser)
	groupUser.DELETE("/delete/:id", apis.DeleteUser, isAdmin)

	//Run port: 8080
	server.Logger.Fatal(server.Start(":8080"))
}
