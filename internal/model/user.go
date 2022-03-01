package model

type Response struct {
	Text    string `json:"text"`
	Message string `json:"message"`
}

type User struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	UUID     string  `json:"uuid"`
	Account  Account `json:"account"`
}