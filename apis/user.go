package apis

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"smartPOST/database"
	"smartPOST/entities"
	"strconv"
	"sync"
	"time"
)

var (
	listUsers = map[int]*entities.User{}
	seq       = 1
	lock      = sync.Mutex{}
)

func CreateUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	u := &entities.User{
		Id:        seq,
		CreateOn:  time.Now(),
		LastLogin: time.Now(),
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	listUsers[u.Id] = u
	seq++
	db := database.DB
	db.Create(&u)
	return c.JSON(http.StatusCreated, u)
}

func GetUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, listUsers[id])
}

func UpdateUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	u := new(entities.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	listUsers[id].FirstName = u.FirstName
	listUsers[id].LastName = u.LastName
	listUsers[id].Email = u.Email
	db := database.DB
	db.Save(listUsers[id])
	return c.JSON(http.StatusOK, listUsers[id])
}

func DeleteUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	delete(listUsers, id)
	db := database.DB
	db.Delete(listUsers[id], id)
	return c.JSON(http.StatusOK, "Deleted successfully")
}

func GetAllUsers(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	return c.JSON(http.StatusOK, listUsers)
}
