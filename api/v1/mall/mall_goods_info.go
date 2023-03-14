package mall

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/common/response"
	"main.go/model/mall/request"
	response2 "main.go/model/mall/response"
	"strconv"
)

type MallGoodsInfoApi struct {
}

// 商品搜索
func (m *MallGoodsInfoApi) GoodsSearch(c *gin.Context) {
	pageNumber, _ := strconv.Atoi(c.Query("pageNumber"))
	goodsCategoryId, _ := strconv.Atoi(c.Query("goodsCategoryId"))
	keyword := c.Query("keyword")
	orderBy := c.Query("orderBy")
	isdiscount := c.Query("isdiscount")
	isCount, _ := strconv.Atoi(isdiscount)
	if err, list, total := mallGoodsInfoService.MallGoodsListBySearch(pageNumber, isCount, goodsCategoryId, keyword, orderBy); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:       list,
			TotalCount: total,
			CurrPage:   pageNumber,
			PageSize:   10,
		}, "获取成功", c)
	}
}

func (m *MallGoodsInfoApi) GoodsDetail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err, goodsInfo := mallGoodsInfoService.GetMallGoodsInfo(id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败"+err.Error(), c)
		return
	}
	response.OkWithData(goodsInfo, c)
}

// 商品进行回收
func (m *MallGoodsInfoApi) GoodsRecovery(c *gin.Context) {
	var req request.GoodsRecovery
	c.ShouldBindJSON(&req)
	//验证支付密码
	if err := mallUserService.VerificatePayPasswd(req.UserId, req.PayPsswd); err != nil {
		response.FailWithMessage("密码错误", c)
		return
	}
	//验证用户能有足够金额进行支付
	err := mallUserService.VerificateMoney(req.UserId, req.PayPrice)
	if err != nil {
		response.FailWithMessage("支付金额不够", c)
		return
	}
	//添加回收记录，返回折扣价、原价
	sellingPrince, OriginPrince, endTime, err := mallGoodsInfoService.AddRecoveryInfo(req)
	if err != nil {
		response.FailWithMessage("回收失败", c)
		return
	}
	//扣除并添加用户余额
	err = mallUserService.RecoverMoney(req.UserId, sellingPrince, OriginPrince, endTime)
	if err != nil {
		response.FailWithMessage("支付金额不够", c)
		return
	}
	//删除商品回收记录
	mallGoodsInfoService.DeleteRecoveryInfo(req.UserId)
	response.OkWithMessage("回收成功，等活动结束余额会自动到位", c)
}

// 商品回收查找
func (m *MallGoodsInfoApi) GetRecoveryInfo(c *gin.Context) {
	var GRinfo request.GetGoodsInfo
	c.ShouldBindQuery(&GRinfo)
	//分页查看
	list, total := mallGoodsInfoService.GetRecoveryInfo(GRinfo)
	totalPages := int(total) / GRinfo.PageSize
	if int(total)%GRinfo.PageSize != 0 {
		totalPages++
	}
	resRecoveryinfos := make([]response2.ReqRecoveryInfos, 0)
	for _, v := range list {
		goodname := mallGoodsInfoService.GetGoodName(v.GoodsId)
		recInfo := response2.ReqRecoveryInfos{
			RId:      v.RId,
			UserId:   v.UserId,
			GoodsId:  v.GoodsId,
			GoodsNum: v.GoodsNum,
			PayPrice: v.PayPrice,
			RePrice:  v.RePrice,
			ReTime:   v.ReTime,
			GoodName: goodname,
		}
		resRecoveryinfos = append(resRecoveryinfos, recInfo)
	}
	getRecoveryInfos := &response2.GetRecoveryInfos{
		List:       resRecoveryinfos,
		TotalCount: int(total),
		TotalPage:  totalPages,
		PageNumber: GRinfo.PageNumber,
		PageSize:   GRinfo.PageSize,
	}
	response.OkWithData(getRecoveryInfos, c)
}
