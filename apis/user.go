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
	//ListUsers = map[int]*entities.User{}
	ListUsers = make([]*entities.User, 0)
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
	for i := 1; i <= len(ListUsers); i++ {
		if i == id {
			ListUsers[id-1].FirstName = u.FirstName
			ListUsers[id-1].LastName = u.LastName
			db.Save(ListUsers[id-1])
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
	for i := 1; i <= len(ListUsers); i++ {
		if i == id {
			if ListUsers[id-1].Password != u.Password {
				ListUsers[id-1].Password = u.Password
				db.Save(ListUsers[id-1])
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
	for i := 1; i <= len(ListUsers); i++ {
		if i == id {
			ListUsers[id-1].Email = u.Email
			db.Save(ListUsers[id-1])
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
	k := 1
	db := database.DB
	for i := 1; i <= len(ListUsers); i++ {
		if i != id {
			ListUsers[i-1] = ListUsers[k-1]
			k++
		} else {
			k++
			db.Delete(ListUsers, id)
		}
		i++
	}
	return c.JSON(http.StatusOK, "Deleted successfully")
}

func GetAllUsers(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	return c.JSON(http.StatusOK, ListUsers)
}
