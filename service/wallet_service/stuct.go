package wallet

type WalletRequest struct {
	Uuid string `form:"uuid" valid:"Required"`
}

type WalletInfoResponse struct {
	Uuid    string  `json:"uuid"`
	Balance float64 `json:"balance"`
	Pow     float64 `json:"pow"`
	Pos     float64 `json:"pos"`
}

type WalletActionInfoResponse struct {
	Uuid     string  `json:"uuid"`
	SAddress string  `json:"s_address"`
	DAddress string  `json:"d_address"`
	Amount   float64 `json:"amount"`
	Status   int     `json:"status"`
}

type WalletActionListResponse struct {
	Total   int                        `json:"total"`
	Actions []WalletActionInfoResponse `json:"actions"`
}
