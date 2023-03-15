package entities

import (
	"fmt"
	"time"
)

type User struct {
	Id        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreateOn  time.Time `json:"createOn"`
	LastLogin time.Time `json:"lastLogin"`
}

func (user User) ToString() string {
	return fmt.Sprintf("id: %d\nfirstName: %s\nlastName: %s\nPassword: %s\nemail: %s\ncreate_on:%s\nlast_login:%s\n",
		user.Id, user.FirstName, user.LastName, user.Password, user.Email, user.CreateOn, user.LastLogin)
}
