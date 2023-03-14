package response

type CartItemResponse struct {
	CartItemId int `json:"cartItemId"`

	GoodsId int `json:"goodsId"`

	GoodsCount float64 `json:"goodsCount"`

	GoodsName string `json:"goodsName"`

	GoodsCoverImg string `json:"goodsCoverImg"`

	SellingPrice float64 `json:"sellingPrice"`
}
