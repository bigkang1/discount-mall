package mall

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/common/response"
	mallReq "main.go/model/mall/request"
	"strconv"
)

type MallUserCurrencyRecordApi struct {
}

func (m *MallUserCurrencyRecordApi) CreateUserCurrencyRecord(c *gin.Context) {
	var req mallReq.AddUserCurrencyRecord
	token := c.GetHeader("token")
	_ = c.ShouldBindJSON(&req)
	if req.CurrencyAmount < 0 { //|| (req.CurrencyType != 0 && req.CurrencyType != 1)
		response.FailWithMessage("创建失败,参数格式有误", c)
		return
	}
	if err := mallUserCurrencyRecordSvc.CreateUserCurrencyRecord(token, req); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败"+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (m *MallUserCurrencyRecordApi) UserCurrencyRecordList(c *gin.Context) {
	token := c.GetHeader("token")
	pageNumber, _ := strconv.Atoi(c.Query("pageNumber"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	if pageNumber <= 0 || pageSize <= 0 {
		pageNumber = 1
		pageSize = 5
	}
	if err, list, total := mallUserCurrencyRecordSvc.UserCurrencyRecordList(token, pageNumber, pageSize); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败"+err.Error(), c)
	} else if len(list) < 1 {
		// 前端项目这里有一个取数逻辑，如果数组为空，数组需要为[] 不能是Null
		response.OkWithDetailed(response.PageResult{
			List:       make([]interface{}, 0),
			TotalCount: total,
			CurrPage:   pageNumber,
			PageSize:   pageSize,
		}, "SUCCESS", c)
	} else {
		var totalPage int
		if int(total)%pageSize > 0 {
			totalPage = int(total)/pageSize + 1
		} else {
			totalPage = int(total) / pageSize
		}
		response.OkWithDetailed(response.PageResult{
			List:       list,
			TotalCount: total,
			CurrPage:   pageNumber,
			TotalPage:  totalPage,
			PageSize:   pageSize,
		}, "SUCCESS", c)
	}

}

func (m *MallUserCurrencyRecordApi) IsExistPayPassword(c *gin.Context) {
	token := c.GetHeader("token")
	if isExistPayPassword, err := mallUserCurrencyRecordSvc.IsExistPayPassword(token); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(isExistPayPassword, "SUCCESS", c)
	}
}
