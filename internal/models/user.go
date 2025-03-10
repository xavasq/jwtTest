package models

type User struct {
	ID       uint   `json:"ID"`
	Name     string `json:"Name"`
	Password string `json:"Password"`
}
