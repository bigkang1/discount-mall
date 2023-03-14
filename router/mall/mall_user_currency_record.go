package mall

import (
	"github.com/gin-gonic/gin"
	v1 "main.go/api/v1"
	"main.go/middleware"
)

type MallUserCurrencyRecordRouter struct {
}

func (m *MallUserCurrencyRecordRouter) InitMallUserCurrencyRecordRouter(Router *gin.RouterGroup) {
	mallUserAddressRouter := Router.Group("v1").Use(middleware.UserJWTAuth())
	var mallRecordApi = v1.ApiGroupApp.MallApiGroup.MallUserCurrencyRecordApi
	{
		mallUserAddressRouter.POST("/userCurrencyRecords", mallRecordApi.CreateUserCurrencyRecord)
		mallUserAddressRouter.GET("/userCurrencyRecords", mallRecordApi.UserCurrencyRecordList)
		mallUserAddressRouter.GET("/isExistPayPassword", mallRecordApi.IsExistPayPassword)
	}

}
