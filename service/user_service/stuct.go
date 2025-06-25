package user

import (
	wallet "nyx_api/service/wallet_service"
)

type CreateRequest struct {
	Username string `json:"username" valid:"Required"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

type UpdateRequest struct {
	Uuid         string `json:"uuid" valid:"Required"`
	Name         string `json:"name"`
	Avatar       string `json:"avatar"`
	Introduction string `json:"introduction"`
	RoleId       int    `json:"roleId"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
}

type DeleteRequest struct {
	Uuid string `json:"uuid" valid:"Required"`
}

type LoginRequest struct {
	Username string `json:"username" valid:"Required"`
	Password string `json:"password" valid:"Required"`
}

type UserListRequest struct {
	Page  string `form:"page" valid:"Required"`
	Limit string `form:"limit" valid:"Required"`
}

type UserListResponse struct {
	Total int            `json:"total"`
	Users []InfoResponse `json:"users"`
}

type LogoutRequest struct {
	Token string `json:"token" valid:"Required"`
}

type InfoResponse struct {
	Uuid         string                    `json:"uuid"`
	Username     string                    `json:"username"`
	Name         string                    `json:"name"`
	Avatar       string                    `json:"avatar"`
	Introduction string                    `json:"introduction"`
	Roles        []string                  `json:"roles"`
	Phone        string                    `json:"phone"`
	Address      string                    `json:"address"`
	WorkersCount int                       `json:"workersCount"`
	WalletInfo   wallet.WalletInfoResponse `json:"walletInfo,omitempty"`
}
