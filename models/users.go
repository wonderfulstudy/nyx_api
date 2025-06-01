package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Uid          string `gorm:"primary_key"`
	Name         string
	Phone        string
	Introduction string
	Address      string
}
