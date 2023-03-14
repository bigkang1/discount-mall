package manage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/common/request"
	"main.go/model/common/response"
	request2 "main.go/model/mall/request"
	response2 "main.go/model/mall/response"
	"main.go/model/manage"
	manageReq "main.go/model/manage/request"
	"strconv"
)

type ManageGoodsInfoApi struct {
}

func (m *ManageGoodsInfoApi) CreateGoodsInfo(c *gin.Context) {
	var mallGoodsInfo manageReq.GoodsInfoAddParam
	_ = c.ShouldBindJSON(&mallGoodsInfo)
	fmt.Println("接收到的信息：", mallGoodsInfo)
	if mallGoodsInfo.SellingPrice > 1 || mallGoodsInfo.SellingPrice < 0 {
		response.FailWithMessage("折扣不合理!", c)
		return
	}
	if err := mallGoodsInfoService.CreateMallGoodsInfo(mallGoodsInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败!"+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMallGoodsInfo 删除MallGoodsInfo
func (m *ManageGoodsInfoApi) DeleteGoodsInfo(c *gin.Context) {
	var mallGoodsInfo manage.MallGoodsInfo
	_ = c.ShouldBindJSON(&mallGoodsInfo)
	if err := mallGoodsInfoService.DeleteMallGoodsInfo(mallGoodsInfo); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败"+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// ChangeMallGoodsInfoByIds 批量删除MallGoodsInfo
func (m *ManageGoodsInfoApi) ChangeGoodsInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	sellStatus := c.Param("status")
	if err := mallGoodsInfoService.ChangeMallGoodsInfoByIds(IDS, sellStatus); err != nil {
		global.GVA_LOG.Error("修改商品状态失败!", zap.Error(err))
		response.FailWithMessage("修改商品状态失败"+err.Error(), c)
	} else {
		response.OkWithMessage("修改商品状态成功", c)
	}
}

// UpdateMallGoodsInfo 更新MallGoodsInfo
func (m *ManageGoodsInfoApi) UpdateGoodsInfo(c *gin.Context) {
	var mallGoodsInfo manageReq.GoodsInfoUpdateParam
	_ = c.ShouldBindJSON(&mallGoodsInfo)
	fmt.Println("mallGoodsInfo:", mallGoodsInfo)
	if err := mallGoodsInfoService.UpdateMallGoodsInfo(mallGoodsInfo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败"+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMallGoodsInfo 用id查询MallGoodsInfo
func (m *ManageGoodsInfoApi) FindGoodsInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err, goodsInfo := mallGoodsInfoService.GetMallGoodsInfo(id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败"+err.Error(), c)
		return
	}
	goodsInfoRes := make(map[string]interface{})
	goodsInfoRes["goods"] = goodsInfo
	if _, thirdCategory := mallGoodsCategoryService.SelectCategoryById(goodsInfo.GoodsCategoryId); thirdCategory != (manage.MallGoodsCategory{}) {
		goodsInfoRes["thirdCategory"] = thirdCategory
		if _, secondCategory := mallGoodsCategoryService.SelectCategoryById(thirdCategory.ParentId); secondCategory != (manage.MallGoodsCategory{}) {
			goodsInfoRes["secondCategory"] = secondCategory
			if _, firstCategory := mallGoodsCategoryService.SelectCategoryById(secondCategory.ParentId); firstCategory != (manage.MallGoodsCategory{}) {
				goodsInfoRes["firstCategory"] = firstCategory
			}
		}
	}
	response.OkWithData(goodsInfoRes, c)

}

// GetMallGoodsInfoList 分页获取MallGoodsInfo列表
func (m *ManageGoodsInfoApi) GetGoodsInfoList(c *gin.Context) {
	var pageInfo manageReq.MallGoodsInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := mallGoodsInfoService.GetMallGoodsInfoInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:       list,
			TotalCount: total,
			CurrPage:   pageInfo.PageNumber,
			PageSize:   pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// 商品回收查找
func (m *ManageGoodsInfoApi) GetRecoveryInfo(c *gin.Context) {
	var GRinfo request2.GetGoodsInfo
	c.ShouldBindQuery(&GRinfo)
	//分页查看
	list, total := mallGoodsInfoService.GetRecoveryInfo(GRinfo)
	if GRinfo.PageSize == 0 {
		fmt.Println("Grpinf是:", GRinfo)
		response.FailWithMessage("pagesize位0", c)
		return
	}
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
