package models

import "time"

type Wallet struct {
	Id      int `gorm:"primary_key"`
	Uuid    string
	Balance float64
	Pow     float64
	Pos     float64
}

type WalletAction struct {
	Id        int `gorm:"primary_key"`
	Uuid      string
	ActionId  int `gorm:"foreignkey:ActionId;references:ID;onDelete:SET NULL"`
	SAddress  string
	DAddress  string
	Amount    float64
	Status    int
	CreatedAt time.Time `gorm:"timestamp"`
	UpdatedAt time.Time `gorm:"timestamp"`
}

func GetWalletByUuid(uuid string) (*Wallet, error) {
	var wallet Wallet
	result := db.Model(&Wallet{}).Where("uuid = ?", uuid).First(&wallet)
	if result.Error != nil {
		return nil, result.Error
	}
	return &wallet, nil
}

func GetActionsByUuid(uuid string) ([]WalletAction, error) {
	var walletActions []WalletAction
	result := db.Model(&WalletAction{}).Where("uuid = ?", uuid).Find(&walletActions)
	if result.Error != nil {
		return []WalletAction{}, result.Error
	}
	return walletActions, nil
}

func GetActionCount(uuid string) int {
	var count int
	db.Model(&WalletAction{}).Where("uuid = ?", uuid).Count(&count)
	return count
}
