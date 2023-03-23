package models

import "time"

type LoginRequest struct {
	Email    string `json:"email" form:"email" xml:"email" query:"email"`
	Password string `json:"password" form:"password" xml:"password" query:"password"`
}

type LoginResponse struct {
	Id        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreateOn  time.Time `json:"createOn"`
	LastLogin time.Time `json:"lastLogin"`
	Role      string    `json:"role"`
	Token     string    `json:"token"`
	Admin     bool      `json:"admin"`
}

type Content struct {
	Text string `json:"text"`
}
