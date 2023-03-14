package response

import "main.go/model/common"

type ResUserChartList struct {
	Uid         int    `json:"uid"`
	NickName    string `json:"nickName"`
	LoginName   string `json:"loginName"`
	AdminUnRead int    `json:"adminUnRead"`
}

type ChartMsg struct {
	Uid      int    `json:"uid"`
	SendTime int    `json:"sendTime"`
	Msg      string `json:"msg"`
}

type ResHistoryList struct {
	Chart     common.Charts
	AdminName string `json:"adminName"`
}
