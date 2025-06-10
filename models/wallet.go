package models

import "time"

type Wallet struct {
	Id      int     `gorm:"primary_key" json:"id"`
	Uuid    string  `json:"uuid"`
	Balance float64 `json:"balance"`
	Pow     float64 `json:"pow"`
	Pos     float64 `json:"pos"`
}

type WalletAction struct {
	Id        int       `gorm:"primary_key" json:"id"`
	Uuid      string    `json:"uuid"`
	ActionId  int       `json:"actionId"`
	SAddress  string    `json:"s_address"`
	DAddress  string    `json:"d_address"`
	Amount    float32   `json:"amount"`
	Status    int       `json:"status"`
	CreatedAt time.Time `gorm:"timestamp" json:"createdAt"`
	UpdatedAt time.Time `gorm:"timestamp" json:"updateAt"`
}

func GetWalletByUuid(uuid string) (wallet Wallet) {
	db.Where("uuid = ?", uuid).First(&wallet)
	return
}

func GetActionByUuid(uuid string) (walletAction []WalletAction) {
	db.Table("nyx_wallet_action").Where("uuid = ?", uuid).Find(&walletAction)
	return
}
