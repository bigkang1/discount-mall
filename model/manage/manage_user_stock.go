package manage

import "main.go/model/common"

type MallUserStock struct {
	UserStockId int             `json:"userStockId" form:"userStockId" gorm:"primarykey;AUTO_INCREMENT"`
	UserId      int             `json:"userId" form:"userId" gorm:"column:user_id;comment:用户主键id;type:bigint"`
	GoodsId     int             `json:"goodsId" form:"goodsId" gorm:"column:goods_id;comment:关联商品id;type:bigint"`
	GoodsCount  int             `json:"goodsCount" form:"goodsCount" gorm:"column:goods_count;comment:数量;type:int"`
	CreateTime  common.JSONTime `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;type:datetime"`
}

func (MallUserStock) TableName() string {
	return "user_stock"
}
