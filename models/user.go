package models

type User struct {
	Uuid         string `gorm:"primary_key"`
	Name         string
	Phone        string
	Introduction string
	Address      string
}

func GetUserByUuid(uuid string) (user User) {
	db.Where("uuid = ?", uuid)
	return
}

func GetUserByName(name string) (user User) {
	db.Where("name = ?", name)
	return
}

func GetUserByPhone(phone string) (user User) {
	db.Where("phone = ?", phone)
	return
}
