package models

type User struct {
	Id       uint   `json:"id"`
	Fistname string `json:"first_name"`
	Lastname string `json:"last_name"`
	Email    string `json:"email"`
	Password []byte `json:"-"`
	Phone    string `json:"phone"`
}
