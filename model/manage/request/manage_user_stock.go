package request

import (
	"main.go/model/common/request"
)

type MallUserStockSearch struct {
	SearchName string `json:"searchNickName" form:"searchNickName"`
	request.PageInfo
}

type MallUserStockSearchUserId struct {
	SearchUserId int `json:"searchUserId" form:"searchUserId"`
	request.PageInfo
}
