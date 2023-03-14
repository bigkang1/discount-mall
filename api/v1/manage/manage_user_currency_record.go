package manage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/common/response"
	manageReq "main.go/model/manage/request"
	"strconv"
)

type ManageUserCurrencyRecordApi struct {
}

func (m *ManageUserCurrencyRecordApi) UserCurrencyRecordList(c *gin.Context) {
	pageNumber, _ := strconv.Atoi(c.Query("pageNumber"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	if pageNumber <= 0 || pageSize <= 0 {
		pageNumber = 1
		pageSize = 5
	}
	if err, list, total := mallUserCurrencyRecordSvc.UserCurrencyRecordList(pageNumber, pageSize); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败"+err.Error(), c)
	} else if len(list) < 1 {
		// 前端项目这里有一个取数逻辑，如果数组为空，数组需要为[] 不能是Null
		response.OkWithDetailed(response.PageResult{
			List:       make([]interface{}, 0),
			TotalCount: total,
			CurrPage:   pageNumber,
			TotalPage:  0,
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

func (m *ManageUserCurrencyRecordApi) UserCurrencyRecordListByUserName(c *gin.Context) {
	pageNumber, _ := strconv.Atoi(c.Query("pageNumber"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	searchName := c.Query("searchName")
	status, _ := strconv.Atoi(c.Query("status"))
	if pageNumber <= 0 || pageSize <= 0 {
		pageNumber = 1
		pageSize = 5
	}
	if err, list, total := mallUserCurrencyRecordSvc.UserCurrencyRecordListByUserName(searchName, status, pageNumber, pageSize); err != nil {
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

func (m *ManageUserCurrencyRecordApi) UserCurrencyRecordListByUserLoginName(c *gin.Context) {
	pageNumber, _ := strconv.Atoi(c.Query("pageNumber"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	searchLoginName := c.Query("searchLoginName")
	if pageNumber <= 0 || pageSize <= 0 {
		pageNumber = 1
		pageSize = 5
	}
	if err, list, total := mallUserCurrencyRecordSvc.UserCurrencyRecordListByUserLoginName(searchLoginName, pageNumber, pageSize); err != nil {
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

func (m *ManageUserCurrencyRecordApi) UserCurrencyRecordListByStatus(c *gin.Context) {
	pageNumber, _ := strconv.Atoi(c.Query("pageNumber"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	status, _ := strconv.Atoi(c.Query("status"))
	if pageNumber <= 0 || pageSize <= 0 {
		pageNumber = 1
		pageSize = 5
	}
	if err, list, total := mallUserCurrencyRecordSvc.UserCurrencyRecordListByStatus(status, pageNumber, pageSize); err != nil {
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

func (m *ManageUserCurrencyRecordApi) UpdateUserCurrencyRecordStatus(c *gin.Context) {
	var req manageReq.UpdateCurrencyRecordStatus
	_ = c.ShouldBindJSON(&req)
	if err := mallUserCurrencyRecordSvc.UpdateUserCurrencyRecordStatus(req); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

type AccessUserCurrencyRes struct {
	NickName  string `json:"nickName"`
	LoginName string `json:"loginName"`
}

// 同意用户申请提现
func (m *ManageUserCurrencyRecordApi) AccessUserCurrency(c *gin.Context) {
	var req manageReq.AccessUserCurrency
	_ = c.ShouldBindJSON(&req)
	if loginName, nickName, err := mallUserCurrencyRecordSvc.AccessUserCurrency(req); err != nil {
		global.GVA_LOG.Error("同意申请失败!", zap.Error(err))
		response.FailWithDetailed(AccessUserCurrencyRes{
			NickName:  nickName,
			LoginName: loginName,
		}, "同意申请失败", c)
		return
	} else {
		response.OkWithDetailed(AccessUserCurrencyRes{
			NickName:  nickName,
			LoginName: loginName,
		}, "同意申请成功", c)
		return
	}
}
