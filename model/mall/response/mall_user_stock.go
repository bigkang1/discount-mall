package response

import "main.go/model/common"

type MallUserStockResponse struct {
	UserStockId   int             `json:"userStockId" `
	UserId        int             `json:"userId" `
	GoodsId       int             `json:"goodsId"`
	GoodsCount    int             `json:"goodsCount"`
	CreateTime    common.JSONTime `json:"createTime"`
	GoodsName     string          `json:"goodsName"`
	GoodsCoverImg string          `json:"goodsCoverImg"`
	SellingPrice  float64         `json:"sellingPrice"`
}
