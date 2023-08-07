package Handlers

type LoginForm struct {
	Username	string	`json:"username"`
	Password	string	`json:"password"`
}

type RegisterForm struct {
	Username	string `json:"username"`
	Password	string	`json:"password"`
}

type CreateRoomForm struct {
	Name	string 	`json:name`
}