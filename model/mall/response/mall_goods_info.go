package response

type GoodsSearchResponse struct {
	GoodsId       int     `json:"goodsId"`
	GoodsName     string  `json:"goodsName"`
	GoodsIntro    string  `json:"goodsIntro"`
	GoodsCoverImg string  `json:"goodsCoverImg"`
	SellingPrice  float64 `json:"sellingPrice"`
	OriginalPrice float64 `json:"originalPrice"`
}

type GoodsInfoDetailResponse struct {
	GoodsId            int      `json:"goodsId"`
	GoodsName          string   `json:"goodsName"`
	GoodsIntro         string   `json:"goodsIntro"`
	GoodsCoverImg      string   `json:"goodsCoverImg"`
	SellingPrice       float64  `json:"sellingPrice"`
	GoodsDetailContent string   `json:"goodsDetailContent"  `
	OriginalPrice      int      `json:"originalPrice" `
	Tag                string   `json:"tag" form:"tag" `
	GoodsCarouselList  []string `json:"goodsCarouselList" `
}

type ReqRecoveryInfos struct {
	RId      int     `json:"rId" `
	UserId   int     `json:"userId"`
	GoodsId  int     `json:"goodsId"`
	GoodsNum int     `json:"goodsNum"`
	PayPrice float64 `json:"payPrice"`
	RePrice  float64 `json:"rePrice" `
	ReTime   int     `json:"reTime"`
	GoodName string  `json:"goodName"`
}

// 分页查询回收返回
type GetRecoveryInfos struct {
	List       []ReqRecoveryInfos `json:"list"`
	TotalCount int                `json:"totalCount"`
	TotalPage  int                `json:"totalPage"`
	PageNumber int                `json:"pageNumber" form:"pageNumber"` // 页码
	PageSize   int                `json:"pageSize" form:"pageSize"`     // 每页大小
}
