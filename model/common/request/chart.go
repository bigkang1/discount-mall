package request

// 消息结构体
type ChartMsg struct {
	Msg     string `json:"msg"`
	MsgType int    `json:"msgType"`
}
