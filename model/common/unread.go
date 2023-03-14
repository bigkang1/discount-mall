package common

// Charts 结构体
type UnRead struct {
	Uid     int `json:"uid" form:"uid" gorm:"primarykey;"`
	UUnRead int `json:"u_un_read" form:"u_un_read" gorm:"column:u_un_read;comment:用户未读的消息;type:int;"`
	AUnRead int `json:"a_un_read" form:"a_un_read" gorm:"column:a_un_read;comment:管理员未读的消息;type:int;"`
}

// TableName Charts 表名
func (UnRead) TableName() string {
	return "unread"
}
