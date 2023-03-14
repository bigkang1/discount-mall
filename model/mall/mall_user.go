package mall

import (
	"main.go/model/common"
)

type MallUser struct {
	UserId         int             `json:"userId" form:"userId" gorm:"primarykey;AUTO_INCREMENT"`
	NickName       string          `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:用户昵称;type:varchar(50);"`
	LoginName      string          `json:"loginName" form:"loginName" gorm:"column:login_name;comment:登陆名称(默认为手机号);type:varchar(11);"`
	PasswordMd5    string          `json:"passwordMd5" form:"passwordMd5" gorm:"column:password_md5;comment:MD5加密后的密码;type:varchar(32);"`
	IntroduceSign  string          `json:"introduceSign" form:"introduceSign" gorm:"column:introduce_sign;comment:个性签名;type:varchar(100);"`
	IsDeleted      int             `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:注销标识字段(0-正常 1-已注销);type:tinyint"`
	LockedFlag     int             `json:"lockedFlag" form:"lockedFlag" gorm:"column:locked_flag;comment:锁定标识字段(0-未锁定 1-已锁定);type:tinyint"`
	CreateTime     common.JSONTime `json:"createTime" form:"createTime" gorm:"column:create_time;comment:注册时间;type:datetime"`
	Currency       float64         `json:"currency" form:"currency" gorm:"column:currency;comment:充值币;type:float"`
	BankCard       string          `json:"bankCard" form:"bankCard" gorm:"column:bank_card;comment:银行卡号;type:varchar(20)"`
	Cardhilder     string          `json:"cardhilder" form:"cardhilder" gorm:"column:cardhilder;comment:银行卡此有者;type:varchar(30)"`
	PayPasswordMd5 string          `json:"payPasswordMd5" form:"payPasswordMd5" gorm:"column:pay_password_md5;comment:支付密码;type:varchar(32);"`
}

// TableName MallUser 表名
func (MallUser) TableName() string {
	return "tb_newbee_mall_user"
}
