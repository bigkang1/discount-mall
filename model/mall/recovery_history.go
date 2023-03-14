package mall

// 回收历史
type Recovery_history struct {
	RId      int     `json:"rId" form:"rId" gorm:"r_id;AUTO_INCREMENT"`
	UserId   int     `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;type:int"`
	GoodsId  int     `json:"goodsId" form:"goodsId" gorm:"column:goods_id;comment:商品id;type:int"`
	GoodsNum int     `json:"goodsNum" form:"goodsNum" gorm:"column:goods_num;comment:商品数量;type:int"`
	PayPrice float64 `json:"payPrice" form:"payPrice" gorm:"column:pay_price;comment:支付金额;type:float"`
	RePrice  float64 `json:"rePrice" form:"rePrice" gorm:"column:re_price;comment:返回金额;type:float"`
	ReTime   int     `json:"reTime" form:"reTime" gorm:"column:re_time;comment:回收时间;type:int"`
}

// TableName Recovery_history 表名
func (Recovery_history) TableName() string {
	return "recovery_history"
}
