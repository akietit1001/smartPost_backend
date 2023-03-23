package apis

import (
	"net/http"
	"smartPOST/database"
	"smartPOST/entities"
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

var (
	//ListUsers = map[int]*entities.User{}
	ListUsers = make([]*entities.User, 0)
	// ListUsers = database.DB.Select("*").Find(&entities.User{})
	seq  = 1
	lock = sync.Mutex{}
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
	ListUsers = append(ListUsers, u)
	seq++
	db := database.DB
	db.Create(&u)
	return c.JSON(http.StatusCreated, u)
}

func GetUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, ListUsers[id-1])
}

func UpdateNameUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	u := new(entities.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))

	db := database.DB
	for index, user := range ListUsers {
		if user.Id == id {
			ListUsers[index].FirstName = u.FirstName
			ListUsers[index].LastName = u.LastName
			db.Save(ListUsers[index])
			break
		}
	}
	return c.JSON(http.StatusOK, ListUsers[id-1])
}

func UpdatePasswordUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	u := new(entities.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	db := database.DB
	for index, user := range ListUsers {
		if user.Id == id {
			if user.Password != u.Password {
				user.Password = u.Password
				db.Save(ListUsers[index])
			} else {
				return c.JSON(http.StatusBadRequest, "Password is existed")
			}
			break
		}
	}
	return c.JSON(http.StatusOK, ListUsers[id-1])
}

func UpdateEmailUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	u := new(entities.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))

	db := database.DB
	for index, user := range ListUsers {
		if user.Id == id {
			user.Email = u.Email
			db.Save(ListUsers[index])
			break
		}
	}
	return c.JSON(http.StatusOK, ListUsers[id-1])
}

func DeleteUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	//delete(ListUsers, id)
	db := database.DB
	for index, user := range ListUsers {
		if user.Id == id {
			copy(ListUsers[index:], ListUsers[index+1:])
			ListUsers[len(ListUsers)-1] = &entities.User{}
			ListUsers = ListUsers[:len(ListUsers)-1]
			db.Delete(ListUsers, id)
			break
		}
	}
	return c.JSON(http.StatusOK, "Deleted successfully")
}

func GetAllUsers(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	return c.JSON(http.StatusOK, ListUsers)
}
