package models

type User struct {
	Model
	Uid          string
	Name         string
	Phone        string
	Introduction string
	Address      string
}
