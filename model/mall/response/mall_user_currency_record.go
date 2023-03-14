package response

import "main.go/model/common"

type UserCurrencyRecordListResponse struct {
	UserCurrencyRecordId int             `json:"userCurrencyRecordId"`
	UserId               int             `json:"userId"`
	NickName             string          `json:"nickName"`
	LoginName            string          `json:"loginName"`
	CurrencyAmount       float64         `json:"currencyAmount"`
	CurrencyType         int             `json:"currencyType"`
	AdminUserId          int             `json:"adminUserId"`
	AdminNickName        string          `json:"adminNickName"`
	Status               int             `json:"status"`
	CreateTime           common.JSONTime `json:"createTime" `
	UpdateTime           common.JSONTime `json:"updateTime"`
}
