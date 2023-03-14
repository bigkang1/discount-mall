package common

import (
	"github.com/gin-gonic/gin"
	v1 "main.go/api/v1"
)

type CommonUserRouter struct {
}

func (m *CommonUserRouter) InitCommonChartRouter(Router *gin.RouterGroup) {
	commonChartRouter := Router.Group("v1")
	var commonChartApi = v1.ApiGroupApp.CommonGroup.ChartApi
	var mallAdminUserApi = v1.ApiGroupApp.ManageApiGroup.ManageAdminUserApi
	{
		commonChartRouter.GET("/chart", commonChartApi.Chat)                         //实时聊天
		commonChartRouter.GET("/clearuunread", commonChartApi.ClearUUnread)          //清理用户未读聊天
		commonChartRouter.GET("/clearaunread", commonChartApi.ClearAUnread)          //清理管理员未读聊天
		commonChartRouter.POST("/getuserchartlist", commonChartApi.GetUserChartList) //获取用户聊天列表UserGetUunRead
		commonChartRouter.GET("/usergetunread", commonChartApi.UserGetUunRead)       //用户获取未读
		commonChartRouter.GET("/gethistorychart", commonChartApi.GetHistoryChart)    //获取历史记录
		commonChartRouter.PUT("/updateLivingMsg", commonChartApi.UpdateLivingMsg)    //更新留言
		commonChartRouter.GET("/getLivingMsg", commonChartApi.GetLivingMsg)          //查找留言
	}
	{

		commonChartRouter.POST("/upload/file", mallAdminUserApi.UploadFile) //上传图片
	}
}
