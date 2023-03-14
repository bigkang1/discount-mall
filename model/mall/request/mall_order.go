package request

type PaySuccessParams struct {
	OrderNo string `json:"orderNo"`
	PayType int    `json:"payType"`
}

type OrderSearchParams struct {
	Status     string `form:"status"`
	PageNumber int    `form:"pageNumber"`
}

type SaveOrderParam struct {
	CartItemIds []int `json:"cartItemIds"`
	AddressId   int   `json:"addressId"`
}

type SaveUserStockOrderParam struct {
	CartItemIds []int `json:"cartItemIds"`
}

type AddUserStock struct {
	GoodsId    int `json:"goodsId"`
	GoodsCount int `json:"goodsCount"`
}

type DeleteUserStockParam struct {
	GoodsId    int `json:"goodsId"`
	GoodsCount int `json:"goodsCount"`
}
