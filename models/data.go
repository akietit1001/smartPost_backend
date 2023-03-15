package models

type LoginRequest struct {
	Email    string `json:"email" form:"email" xml:"email" query:"email"`
	Password string `json:"password" form:"password" xml:"password" query:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type Content struct {
	Text string `json:"text"`
}
