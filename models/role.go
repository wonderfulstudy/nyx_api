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

func GetRoleById(id int) (Role, error) {
	var role Role
	result := db.Model(&Role{}).Where("id = ?", id).First(&role)
	if result.Error != nil {
		return Role{}, result.Error
	}

	return role, nil
}
