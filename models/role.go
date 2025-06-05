package models

import "time"

type Role struct {
	Id          int `gorm:"primary_key"`
	KeyName     string
	Name        string
	Description string
	IsActive    int       `gorm:"default:1,bigint"`
	CreatedAt   time.Time `gorm:"timestamp"`
	UpdatedAt   time.Time `gorm:"timestamp"`
}

func GetRoleById(Id int) (role Role) {
	db.Where("id = ?", Id).First(&role)
	return
}
