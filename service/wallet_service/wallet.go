package wallet

import (
	"errors"
	"nyx_api/models"
)

func WalletInfoService(uuid string) (WalletInfoResponse, error) {
	wallet, err := models.GetWalletByUuid(uuid)
	if err != nil {
		return WalletInfoResponse{}, errors.New("从数据库中获取用户钱包数据失败" + err.Error())
	}

	return WalletInfoResponse{
		Uuid:    wallet.Uuid,
		Balance: wallet.Balance,
		Pow:     wallet.Pow,
		Pos:     wallet.Pos,
	}, nil
}

func WalletActionListService(req WalletRequest) (WalletActionListResponse, error) {
	var response WalletActionListResponse
	actions, err := models.GetActionsByUuid(req.Uuid)
	if err != nil {
		return WalletActionListResponse{}, errors.New("从数据库中获取用户钱包数据失败" + err.Error())
	}

	for _, action := range actions {
		response.Actions = append(response.Actions, WalletActionInfoResponse{
			Uuid:     action.Uuid,
			SAddress: action.SAddress,
			DAddress: action.DAddress,
			Amount:   action.Amount,
			Status:   action.Status,
		})
	}

	response.Total = models.GetActionCount(req.Uuid)

	return response, nil
}
