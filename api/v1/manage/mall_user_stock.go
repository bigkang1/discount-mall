package manage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/common/request"
	"main.go/model/common/response"
	managereq "main.go/model/manage/request"
)

type ManageUserStockApi struct {
}

func (m *ManageUserStockApi) SearchMallUserStockList(c *gin.Context) {
	var pageInfo managereq.MallUserStockSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if pageInfo.PageNumber <= 0 {
		pageInfo.PageNumber = 1
	}
	if err, list, total := mallUserStockSercvice.MallOrderListBySearch(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else if len(list) < 1 {
		// 前端项目这里有一个取数逻辑，如果数组为空，数组需要为[] 不能是Null
		response.OkWithDetailed(response.PageResult{
			List:       make([]interface{}, 0),
			TotalCount: total,
			CurrPage:   pageInfo.PageNumber,
			PageSize:   pageInfo.PageSize,
		}, "SUCCESS", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:       list,
			TotalCount: total,
			CurrPage:   pageInfo.PageNumber,
			PageSize:   pageInfo.PageSize,
		}, "SUCCESS", c)
	}
}

func (m *ManageUserStockApi) UserStockList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	if pageInfo.PageNumber <= 0 {
		pageInfo.PageNumber = 1
	}
	if err, list, total := mallUserStockSercvice.MallOrderList(pageInfo); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败"+err.Error(), c)
	} else if len(list) < 1 {
		// 前端项目这里有一个取数逻辑，如果数组为空，数组需要为[] 不能是Null
		response.OkWithDetailed(response.PageResult{
			List:       make([]interface{}, 0),
			TotalCount: total,
			CurrPage:   pageInfo.PageNumber,
			PageSize:   pageInfo.PageSize,
		}, "SUCCESS", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:       list,
			TotalCount: total,
			CurrPage:   pageInfo.PageNumber,
			PageSize:   pageInfo.PageSize,
		}, "SUCCESS", c)
	}

}

func (m *ManageUserStockApi) SearchMallUserStockListByUserId(c *gin.Context) {
	var pageInfo managereq.MallUserStockSearchUserId
	_ = c.ShouldBindQuery(&pageInfo)
	if pageInfo.PageNumber <= 0 {
		pageInfo.PageNumber = 1
	}
	if err, list, total := mallUserStockSercvice.MallOrderListBySearchUserID(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else if len(list) < 1 {
		// 前端项目这里有一个取数逻辑，如果数组为空，数组需要为[] 不能是Null
		response.OkWithDetailed(response.PageResult{
			List:       make([]interface{}, 0),
			TotalCount: total,
			CurrPage:   pageInfo.PageNumber,
			PageSize:   pageInfo.PageSize,
		}, "SUCCESS", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:       list,
			TotalCount: total,
			CurrPage:   pageInfo.PageNumber,
			PageSize:   pageInfo.PageSize,
		}, "SUCCESS", c)
	}
}
