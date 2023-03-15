package main

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"smartPOST/apis"
	"smartPOST/controller"
	"smartPOST/database"
	"smartPOST/mdw"
)

func main() {
	server := echo.New()
	database.DBConnection()
	isLoggedIn := echojwt.JWT([]byte("mysecretkey"))
	isAdmin := mdw.IsAdminMdw
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	server.GET("/", controller.Hello, isLoggedIn)
	server.POST("/login", controller.Login, middleware.BasicAuth(mdw.Basic_Auth))

	server.POST("/create", apis.CreateUser)

	server.GET("/admin", controller.Hello, isLoggedIn, isAdmin)

	groupUser := server.Group("/api/user", isLoggedIn)
	groupUser.GET("/get/:id", apis.GetUser)
	groupUser.GET("/getall", apis.GetAllUsers, isAdmin)
	groupUser.PUT("/update/:id", apis.UpdateUser)
	groupUser.DELETE("/delete/:id", apis.DeleteUser, isAdmin)
	//database.DB.Find(entities.User{})

	server.Logger.Fatal(server.Start(":8080"))
}
