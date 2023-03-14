package common

type AdminLivingMsg struct {
	LiveMsg   string `json:"liveMsg" form:"liveMsg" gorm:"column:livemsg;comment:留言;type:string;"`
	StartTime int    `json:"startTime" form:"startTime" gorm:"force;column:start_time;comment:留言开始时间;type:int;"`
	EntTime   int    `json:"endTime" form:"endTime" gorm:"force;column:ent_time;comment:留言结束时间;type:int;"`
	IsUse     int    `json:"isUse" form:"isUse" gorm:"column:is_use;comment:0，使用，1，不使用;type:int;"`
}

// TableName Charts 表名
func (AdminLivingMsg) TableName() string {
	return "admin_live_msg"
}
