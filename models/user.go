package models

import (
	"nyx_api/pkg/setting"
	"time"
)

type User struct {
	Id           int       `gorm:"primary_key"`
	Uuid         string    `json:"uuid"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Name         string    `json:"name"`
	Avatar       string    `json:"avatar"`
	Introduction string    `json:"introduction"`
	RoleId       int       `json:"roleId"`
	Phone        string    `json:"phone"`
	Address      string    `json:"address"`
	Status       int       `json:"status"`
	CreatedAt    time.Time `gorm:"timestamp"`
	UpdatedAt    time.Time `gorm:"timestamp"`
}

func GetByUuid(uuid string) (User, error) {
	var user User
	result := db.Model(&User{}).Where("uuid = ?", uuid).First(&user)
	if result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}

func GetByUsername(userName string) (User, error) {
	var user User
	result := db.Model(&User{}).Where("username = ?", userName).First(&user)
	if result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}

func ListUsers(page, limit int) ([]User, error) {
	var users []User
	result := db.Model(&User{}).
		Offset((page - 1) * setting.PageSize).
		Limit(limit).
		Find(&users)
	if result.Error != nil {
		return []User{}, result.Error
	}
	return users, nil
}

func Count() (int, error) {
	var count int
	result := db.Model(&User{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func AddUser(user *User) error {
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(uuid string) error {
	result := db.Model(&User{}).Where("uuid = ?", uuid).Update("status", 0)
	return result.Error
}

func UpdateUser(user *User) error {
	result := db.Model(&User{}).Where("uuid = ?", user.Uuid).Update(user)
	return result.Error
}
