package model



type LoginModel struct {
	Email    string `json:"email";validate:"required,email"`
	Password string `json:"password";validate:"required,min=3,max=15"`
}
