package models

type Action struct {
	Id   int `gorm:"primary_key"`
	Name string
}

func GetActionById(id string) (Action, error) {
	var action Action
	result := db.Model(&Action{}).Where("id = ?", id).Find(&action)
	if result.Error != nil {
		return Action{}, result.Error
	}
	return action, nil
}
