package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/common"
	"main.go/model/common/response"
	manageReq "main.go/model/manage/request"
	"main.go/utils"
	"net/http"
	"strconv"
)

type ChartApi struct {
}

var upgrader = websocket.Upgrader{
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (a *ChartApi) Chat(c *gin.Context) {
	sidstr := c.Query("sendId")
	sid, _ := strconv.Atoi(sidstr)
	gidstr := c.Query("gedId")
	gid, _ := strconv.Atoi(gidstr)
	isSendAdmin := c.Query("isSendAdmin")
	isAdmin, _ := strconv.Atoi(isSendAdmin)

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("upgrade:", err)
		return
	}

	//接收消息群组
	var chatnode ChartNode
	if isAdmin == 0 {
		//判断身份,用户

		chatnode = InitUserChat(sid, ws)
	} else {
		//判断身份，管理员

		chatnode = InitAdminChat(sid, gid, ws)
	}
	defer chatnode.CloseChat(sid)
	//接受消息
	go chatnode.SendProc(gid, sid, isAdmin)
	//发送消息
	go chatnode.RecvProc()

	chatnode.UnLock()
}

func (a *ChartApi) ClearAUnread(c *gin.Context) {
	uidstr := c.Query("uid")
	uid, _ := strconv.Atoi(uidstr)
	commonChartService.ClearAUnRead(uid)
	response.OkWithMessage("清理成功", c)
}
func (a *ChartApi) ClearUUnread(c *gin.Context) {
	uidstr := c.Query("uid")
	uid, _ := strconv.Atoi(uidstr)
	commonChartService.ClearUUnRead(uid)
	response.OkWithMessage("清理成功", c)
}

// 获取用户聊天列表
func (a *ChartApi) GetUserChartList(c *gin.Context) {
	var pageInfo manageReq.MallUserSearch
	_ = c.ShouldBindJSON(&pageInfo)
	err, userList, total := commonChartService.GetUserList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	resUserChartList := make([]*response.ResUserChartList, 0)
	for _, v := range userList {
		unread := commonChartService.GetUserunread(v.UserId)
		resChart := &response.ResUserChartList{
			Uid:         v.UserId,
			NickName:    v.NickName,
			LoginName:   v.LoginName,
			AdminUnRead: unread.AUnRead,
		}
		resUserChartList = append(resUserChartList, resChart)
	}
	//排序
	utils.UpSort(resUserChartList)
	response.OkWithDetailed(response.PageResult{
		List:       resUserChartList,
		TotalCount: total,
		CurrPage:   pageInfo.PageNumber,
		PageSize:   pageInfo.PageSize,
	}, "获取成功", c)
}

// 用户端检查未读消息
func (a *ChartApi) UserGetUunRead(c *gin.Context) {
	uidstr := c.Query("uid")
	uid, _ := strconv.Atoi(uidstr)
	err, n := commonChartService.UserGetunread(uid)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.OkWithDetailed(0, "没有聊天数据", c)
		return
	}
	response.OkWithDetailed(n, "获取成功", c)
}

// 获取聊天记录
func (a *ChartApi) GetHistoryChart(c *gin.Context) {
	uidstr := c.Query("uid")
	uid, _ := strconv.Atoi(uidstr)
	commonChartService.ClearHistory() //清理历史聊天记录
	err, history := commonChartService.GetHistoryUserChart(uid)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	reqlist := make([]response.ResHistoryList, 0)
	for _, v := range history {
		if (v.SendId == uid && v.IsSendAdmin == 0) || v.GetId == uid && v.IsSendAdmin == 1 {
			var adminName string
			if v.IsSendAdmin == 1 {
				adminName = commonChartService.GetAdminName(v.SendId)
			}
			rev := response.ResHistoryList{
				Chart:     v,
				AdminName: adminName,
			}
			reqlist = append(reqlist, rev)
		}
	}
	response.OkWithDetailed(reqlist, "获取成功", c)
}

// 修改留言
func (a *ChartApi) UpdateLivingMsg(c *gin.Context) {
	var msgInfo common.AdminLivingMsg
	if msgInfo.StartTime < 0 || msgInfo.StartTime > 23 || msgInfo.EntTime < 0 || msgInfo.EntTime > 23 {
		global.GVA_LOG.Error("留言时间必须在0-23小时直接")
		return
	}
	err := c.ShouldBindJSON(&msgInfo)
	err = commonChartService.UpdateLivingMsg(msgInfo)
	if err != nil {
		global.GVA_LOG.Error("留言更新失败!", zap.Error(err))
		response.FailWithMessage("留言更新失败", c)
		return
	}
	response.OkWithMessage("留言更新成功", c)
}

// 获取留言
func (a *ChartApi) GetLivingMsg(c *gin.Context) {
	msginfo := commonChartService.GetLivingMsg()

	response.OkWithDetailed(msginfo, "获取聊天消息", c)
}
