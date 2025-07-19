package models

type User struct {
	ID        int    `json:"ID"`
	Username  string `json:"Username"`
	Email     string `json:"Email"`
	Password  string `json:"Password"`
	Validated bool   `json:"Validated"`
}
