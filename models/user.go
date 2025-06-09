package models

import (
	"nyx_api/middleware/aes"
	"nyx_api/pkg/setting"
	"time"
)

type User struct {
	Id           int `gorm:"primary_key"`
	Uuid         string
	Username     string
	Password     string
	Avatar       string
	Name         string
	Introduction string
	Token        string
	RoleId       int
	Phone        string
	Address      string
	CreatedAt    time.Time `gorm:"timestamp"`
	UpdatedAt    time.Time `gorm:"timestamp"`
}

func GetUserByUuid(uuid string) (user User) {
	db.Where("uuid = ?", uuid).First(&user)
	user.Password, _ = aes.AesDecryptCBCBase64(user.Password)
	return
}

func GetUserByName(userName string) (user User) {
	db.Where("username = ?", userName).First(&user)
	user.Password, _ = aes.AesDecryptCBCBase64(user.Password)
	return
}

func GetUsersByToken(token string) (users []User) {
	db.Where("token = ?", token).Find(&users)
	for i := 0; i < len(users); i++ {
		users[i].Password, _ = aes.AesDecryptCBCBase64(users[i].Password)
	}
	return
}

func GetUserByToken(token string) (user User) {
	db.Where("token = ?", token).First(&user)
	user.Password, _ = aes.AesDecryptCBCBase64(user.Password)
	return
}

func GetUserList(page, limit int) (user []User) {
	db.Where("status = ?", 1).Offset((page - 1) * setting.PageSize).Limit(limit).Find(&user)
	return
}

func GetUserCount() (count int) {
	db.Where("status = ?", 1).Model(&User{}).Count(&count)
	return
}

func AddUser(user *User) {
	db.Create(&user)
	return
}

func DeleteUser(user *User) {
	db.Table("nyx_user").Where("uuid = ?", user.Uuid).Update("status", 0)
	return
}

func UpdateUser(user *User) {
	db.Table("nyx_user").Where("uuid = ?", user.Uuid).Update(user)
	return
}
