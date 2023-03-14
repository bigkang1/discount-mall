package manage

import (
	"github.com/gin-gonic/gin"
	v1 "main.go/api/v1"
	"main.go/middleware"
)

type ManageUserStockRouter struct {
}

func (r *ManageUserStockRouter) InitManageUserStockRouter(Router *gin.RouterGroup) {
	mallOrderRouter := Router.Group("v1").Use(middleware.AdminJWTAuth())
	var mallOrderApi = v1.ApiGroupApp.ManageApiGroup.ManageUserStockApi
	{
		mallOrderRouter.GET("userStocks", mallOrderApi.UserStockList)
		mallOrderRouter.GET("userStocks/search", mallOrderApi.SearchMallUserStockList)
		mallOrderRouter.GET("userStocks/searchByUserId", mallOrderApi.SearchMallUserStockListByUserId)
	}
}
