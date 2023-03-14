package common

// Charts 结构体
type Charts struct {
	MsgId       int    `json:"msgId" form:"msgId" gorm:"primarykey;AUTO_INCREMENT"`
	SendId      int    `json:"sendId" form:"sendId" gorm:"column:sendId;comment:发送信息的uid;type:int;"`
	GetId       int    `json:"getId" form:"getId" gorm:"column:getId;comment:接受信息的uid;type:int;"`
	IsSendAdmin int    `json:"isSendAdmin"form:"isSendAdmin" gorm:"column:isSendAdmin;comment:是否是管理员发送的消息,0:管理员发生的，1：用户发送的;type:int;"`
	SendTime    int    `json:"sendTime" form:"sendTime" gorm:"column:sendTime;comment:发送消息的时间;type:int;"`
	MsgType     int    `json:"msgType" form:"msgType" gorm:"column:msgType;comment:发送消息的类型，0:文本，1：文件;type:int"`
	Msg         string `json:"msg" form:"msg" gorm:"column:msg;comment:发送的消息，认为文件，则是路径;type:blob"`
}

// TableName Charts 表名
func (Charts) TableName() string {
	return "charts"
}
