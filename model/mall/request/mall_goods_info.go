package request

import (
	"main.go/model/common/request"
	"main.go/model/mall"
)

type GoodsSearchParams struct {
	Keyword         string `form:"keyword"`
	GoodsCategoryId int    `form:"goodsCategoryId"`
	OrderBy         string `form:"orderBy"`
	PageNumber      int    `form:"pageNumber"`
}

// 回收添加
type GoodsRecovery struct {
	mall.Recovery_history
	PayPsswd string `json:"payPsswd" form:"payPsswd"`
}

// 查询回收
type GetGoodsInfo struct {
	UserId int `json:"userId" form:"userId"`
	request.PageInfo
}
