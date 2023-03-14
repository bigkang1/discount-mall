package mall

import (
	"github.com/gin-gonic/gin"
	v1 "main.go/api/v1"
	"main.go/middleware"
)

type MallUserStockRouter struct {
}

func (m *MallUserStockRouter) InitMallUserStockRouter(Router *gin.RouterGroup) {
	mallOrderRouter := Router.Group("v1").Use(middleware.UserJWTAuth())

	var mallOrderRouterApi = v1.ApiGroupApp.MallApiGroup.MallUserStockApi
	{
		mallOrderRouter.GET("/userStock/list", mallOrderRouterApi.UserStockList)     //订单列表接口
		mallOrderRouter.POST("/addUserStock", mallOrderRouterApi.SaveOrder)          //下单囤货
		mallOrderRouter.POST("/deleteUserStock", mallOrderRouterApi.DeleteUserStock) //回收
	}
}
