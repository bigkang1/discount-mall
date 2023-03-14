package manage

import (
	"github.com/gin-gonic/gin"
	v1 "main.go/api/v1"
	"main.go/middleware"
)

type ManageInviteCodeRouter struct {
}

func (r *ManageInviteCodeRouter) InitManageInviteCodeRouter(Router *gin.RouterGroup) {
	mallInviteCodeRouter := Router.Group("v1").Use(middleware.AdminJWTAuth())
	var mallInviteCodeApi = v1.ApiGroupApp.ManageApiGroup.ManageInviteCodeApi
	{
		mallInviteCodeRouter.POST("inviteCodes", mallInviteCodeApi.CreateInviteCode) // 新建邀请码
		mallInviteCodeRouter.GET("inviteCodes", mallInviteCodeApi.GetInviteCodeList)
		mallInviteCodeRouter.DELETE("inviteCodes", mallInviteCodeApi.DeleteInviteCode) //删除邀请码
	}
}
