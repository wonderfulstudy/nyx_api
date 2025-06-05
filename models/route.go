package models

import "time"

type Route struct {
	Id        int `gorm:"primary_key"`
	Path      string
	Name      string
	Component string
	Redirect  string
	Meta      string `gorm:"json"`
	ParentId  int
	SortOrder int
	CreatedAt time.Time `gorm:"timestamp"`
	UpdatedAt time.Time `gorm:"timestamp"`
}

func GetRouteById(Id int) (route Route) {
	db.Where("id = ?", Id).First(&route)
	return
}
