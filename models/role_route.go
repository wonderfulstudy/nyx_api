package models

type RoleRoute struct {
	RoleId  int
	RouteId int
}

func GetRoleRouteById(Id int) (roleRoute RoleRoute) {
	db.Where("id = ?", Id).First(&roleRoute)
	return
}
