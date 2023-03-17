package models

import "time"

type LoginRequest struct {
	Email    string `json:"email" form:"email" xml:"email" query:"email"`
	Password string `json:"password" form:"password" xml:"password" query:"password"`
}

type LoginResponse struct {
	Id        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreateOn  time.Time `json:"create_on"`
	LastLogin time.Time `json:"last_login"`
	Role      string    `json:"role"`
	Token     string    `json:"token"`
	Admin     bool      `json:"admin"`
}

type Content struct {
	Text string `json:"text"`
}
