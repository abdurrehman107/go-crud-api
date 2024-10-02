package model

type User struct {
	Id      int    `json:"Id"`
	Name    string `json:"Name"`
	Age     int    `json:"Age"`
	Address string `json:"Address"`
}
