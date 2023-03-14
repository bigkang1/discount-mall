package request

import (
	"main.go/model/common/request"
	"main.go/model/manage"
)

type MallUserSearch struct {
	manage.MallUser
	request.PageInfo
}

type AddUserCurrency struct {
	UserId      int     `json:"userId"`
	AddCurrency float64 `json:"addCurrency"`
}

type DeleteUserCurrency struct {
	UserId         int     `json:"userId"`
	DeleteCurrency float64 `json:"deleteCurrency"`
}

type ResetUserPassword struct {
	UserId int `json:"userId"`
}

type UpdateUserBankCard struct {
	BankCard   string `json:"bankCard"`
	Cardhilder string `json:"cardhilder"`
	UserId     int    `json:"userId"`
}
