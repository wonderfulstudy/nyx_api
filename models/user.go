package models

type User struct {
	Uuid         string `gorm:"primary_key"`
	Name         string
	Phone        string
	Password     string
	Introduction string
	Address      string
}

func GetUserByUuid(uuid string) (user User) {
	db.Where("uuid = ?", uuid).First(&user)
	return
}

func GetUserByName(name string) (user User) {
	db.Where("name = ?", name).First(&user)
	return
}

func GetUserByPhone(phone string) (user User) {
	db.Where("phone = ?", phone).First(&user)
	return
}
