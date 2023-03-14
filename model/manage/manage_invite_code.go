package manage

import "main.go/model/common"

type MallInviteCode struct {
	InviteCodeId int             `json:"inviteCodeId" form:"inviteCodeId" gorm:"primarykey;AUTO_INCREMENT"`
	InviteCode   string          `json:"inviteCode" form:"inviteCode" gorm:"column:invite_code;comment:邀请码;type:varchar(6);"`
	CreateTime   common.JSONTime `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;type:datetime"`
}

// TableName MallCarousel 表名
func (MallInviteCode) TableName() string {
	return "invite_code"
}
