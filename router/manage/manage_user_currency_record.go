package manage

import (
	"github.com/gin-gonic/gin"
	v1 "main.go/api/v1"
	"main.go/middleware"
)

type MallUserCurrencyRecordRouter struct {
}

func (m *MallUserCurrencyRecordRouter) InitMallUserCurrencyRecordRouter(Router *gin.RouterGroup) {
	mallUserRouter := Router.Group("v1").Use(middleware.AdminJWTAuth())
	var mallRecordApi = v1.ApiGroupApp.ManageApiGroup.ManageUserCurrencyRecordApi
	{
		mallUserRouter.POST("/userCurrencyRecords/status", mallRecordApi.UpdateUserCurrencyRecordStatus)
		mallUserRouter.GET("/userCurrencyRecords", mallRecordApi.UserCurrencyRecordList)
		mallUserRouter.GET("/userCurrencyRecords/userName", mallRecordApi.UserCurrencyRecordListByUserName)
		mallUserRouter.GET("/userCurrencyRecords/userLoginName", mallRecordApi.UserCurrencyRecordListByUserLoginName)
		mallUserRouter.GET("/userCurrencyRecords/status", mallRecordApi.UserCurrencyRecordListByStatus)
		mallUserRouter.POST("/userCurrencyRecords/accessApply", mallRecordApi.AccessUserCurrency)
	}

}
