package manage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/common/request"
	"main.go/model/common/response"
	manageReq "main.go/model/manage/request"
)

type ManageInviteCodeApi struct {
}

func (m *ManageInviteCodeApi) CreateInviteCode(c *gin.Context) {
	var req manageReq.MallInviteCodeAddParam
	_ = c.ShouldBindJSON(&req)
	if len(req.InviteCode) < 4 || len(req.InviteCode) > 6 {
		response.FailWithMessage("创建失败,邀请码长度为4-6", c)
		return
	}
	if err := mallInviteCodeService.CreateInviteCode(req); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.OkWithMessage("创建失败"+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (m *ManageInviteCodeApi) DeleteInviteCode(c *gin.Context) {
	var ids request.IdsReq
	_ = c.ShouldBindJSON(&ids)
	if err := mallInviteCodeService.DeleteInviteCode(ids); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败"+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// GetInviteCodeList 分页获取邀请码列表
func (m *ManageInviteCodeApi) GetInviteCodeList(c *gin.Context) {
	var pageInfo manageReq.MallInviteCodeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := mallInviteCodeService.GetInviteCodeList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!"+err.Error(), zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:       list,
			TotalCount: total,
			CurrPage:   pageInfo.PageNumber,
			PageSize:   pageInfo.PageSize,
		}, "获取成功", c)
	}
}
