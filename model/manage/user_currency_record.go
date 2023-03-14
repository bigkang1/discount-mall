package manage

import "main.go/model/common"

type MallUserCurrencyRecord struct {
	UserCurrencyRecordId int             `json:"userCurrencyRecordId" form:"userCurrencyRecordId" gorm:"primarykey;AUTO_INCREMENT"`
	UserId               int             `json:"userId" form:"userId" gorm:"column:user_id;comment:用户主键id;type:bigint"`
	CurrencyAmount       float64         `json:"currencyAmount" form:"currencyAmount" gorm:"column:currency_amount;comment:交易金额;type:float"`
	CurrencyType         int             `json:"currencyType" form:"currencyType" gorm:"column:currency_type;comment:交易类型;type:int"`
	AdminUserId          int             `json:"adminUserId" form:"adminUserId" gorm:"column:admin_user_id;type:bigint"`
	Status               int             `json:"status" form:"status" gorm:"column:status;comment:交易状态: 0-失败 1-成功 2-处理中;type:tinyint"`
	CreateTime           common.JSONTime `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;type:datetime"`
	UpdateTime           common.JSONTime `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:修改时间;type:datetime"`
}

func (MallUserCurrencyRecord) TableName() string {
	return "user_currency_record"
}
