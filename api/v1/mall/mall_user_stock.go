package mall

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/common/response"
	mallReq "main.go/model/mall/request"
	"main.go/utils"
	"strconv"
)

type MallUserStockApi struct {
}

func (m *MallUserStockApi) SaveOrder(c *gin.Context) {
	var saveOrderParam mallReq.SaveUserStockOrderParam
	_ = c.ShouldBindJSON(&saveOrderParam)
	if err := utils.Verify(saveOrderParam, utils.SaveUserStockOrderParamVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	token := c.GetHeader("token")

	priceTotal := 0.00
	err, itemsForSave := mallShopCartService.GetCartItemsForSettle(token, saveOrderParam.CartItemIds)
	if len(itemsForSave) < 1 {
		response.FailWithMessage("无数据:"+err.Error(), c)
	} else {
		//总价
		for _, newBeeMallShoppingCartItemVO := range itemsForSave {
			priceTotal = priceTotal + newBeeMallShoppingCartItemVO.GoodsCount*newBeeMallShoppingCartItemVO.SellingPrice
		}
		if priceTotal <= 0 {
			response.FailWithMessage("价格异常", c)
			return
		}

		if err := mallUserStockService.SaveOrder(token, itemsForSave); err != nil {
			global.GVA_LOG.Error("囤货失败", zap.Error(err))
			response.FailWithMessage("囤货失败:"+err.Error(), c)
		} else {
			response.OkWithMessage("囤货成功", c)
		}
	}
}

func (m *MallUserStockApi) UserStockList(c *gin.Context) {
	token := c.GetHeader("token")
	pageNumber, _ := strconv.Atoi(c.Query("pageNumber"))
	if pageNumber <= 0 {
		pageNumber = 1
	}
	if err, list, total := mallUserStockService.MallOrderListBySearch(token, pageNumber); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败"+err.Error(), c)
	} else if len(list) < 1 {
		// 前端项目这里有一个取数逻辑，如果数组为空，数组需要为[] 不能是Null
		response.OkWithDetailed(response.PageResult{
			List:       make([]interface{}, 0),
			TotalCount: total,
			CurrPage:   pageNumber,
			PageSize:   5,
		}, "SUCCESS", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:       list,
			TotalCount: total,
			CurrPage:   pageNumber,
			PageSize:   5,
		}, "SUCCESS", c)
	}

}

func (m *MallUserStockApi) DeleteUserStock(c *gin.Context) {
	var saveOrderParam mallReq.DeleteUserStockParam
	_ = c.ShouldBindJSON(&saveOrderParam)
	if saveOrderParam.GoodsId <= 0 || saveOrderParam.GoodsCount <= 0 {
		response.FailWithMessage("GoodsId、GoodsCount格式错误", c)
	}
	token := c.GetHeader("token")

	if err := mallUserStockService.DeleteUserStock(token, saveOrderParam.GoodsId, saveOrderParam.GoodsCount); err != nil {
		global.GVA_LOG.Error("回收失败", zap.Error(err))
		response.FailWithMessage("回收失败:"+err.Error(), c)
	} else {
		response.OkWithMessage("回收成功", c)
	}

}
