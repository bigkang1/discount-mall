package mall

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/common/response"
	mallReq "main.go/model/mall/request"
	response2 "main.go/model/mall/response"
	"main.go/utils"
	"strconv"
)

type MallUserApi struct {
}

func (m *MallUserApi) UserRegister(c *gin.Context) {
	var req mallReq.RegisterUserParam
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.MallUserRegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := mallUserService.RegisterUser(req); err != nil {
		global.GVA_LOG.Error("创建失败", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (m *MallUserApi) UserInfoUpdate(c *gin.Context) {
	var req mallReq.UpdateUserInfoParam
	_ = c.ShouldBindJSON(&req)
	token := c.GetHeader("token")
	if err := mallUserService.UpdateUserInfo(token, req); err != nil {
		global.GVA_LOG.Error("更新用户信息失败", zap.Error(err))
		response.FailWithMessage("更新用户信息失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (m *MallUserApi) GetUserInfo(c *gin.Context) {
	token := c.GetHeader("token")
	if err, userDetail := mallUserService.GetUserDetail(token); err != nil {
		global.GVA_LOG.Error("未查询到记录", zap.Error(err))
		response.FailWithMessage("未查询到记录", c)
	} else {
		response.OkWithData(userDetail, c)
	}
}

func (m *MallUserApi) UserLogin(c *gin.Context) {
	var req mallReq.UserLoginParam
	_ = c.ShouldBindJSON(&req)
	if err, _, adminToken := mallUserService.UserLogin(req); err != nil {
		response.FailWithMessage("登陆失败", c)
	} else {
		response.OkWithData(adminToken.Token, c)
	}
}

func (m *MallUserApi) UserLogout(c *gin.Context) {
	token := c.GetHeader("token")
	if err := mallUserTokenService.DeleteMallUserToken(token); err != nil {
		response.FailWithMessage("登出失败", c)
	} else {
		response.OkWithMessage("登出成功", c)
	}

}

// 获取用户余额
func (m *MallUserApi) GetUserCurrecy(c *gin.Context) {
	uidstr := c.Query("uid")
	uid, err := strconv.Atoi(uidstr)
	if err != nil {
		response.FailWithMessage("用户信息错误", c)
		return
	}
	currecy := mallUserService.GetUser(uid)
	response.OkWithData(currecy, c)
}

// 更新用户银行卡号和支付密码
func (m *MallUserApi) UpdateUserBankCard(c *gin.Context) {
	var req mallReq.UpdateUserBankCard
	_ = c.ShouldBindJSON(&req)
	token := c.GetHeader("token")
	if req.BankCard == "" || req.Cardhilder == "" {
		response.FailWithMessage("银行卡号或持卡人信息为空！", c)
		return
	}
	if err := mallUserService.UpdateUserBankCard(token, req); err != nil {
		global.GVA_LOG.Error("更新用户银行卡号失败", zap.Error(err))
		response.FailWithMessage("更新用户银行卡号失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (m *MallUserApi) ValidateUserPayPassword(c *gin.Context) {
	var req mallReq.ValidateUserPayPassword
	_ = c.ShouldBindJSON(&req)
	token := c.GetHeader("token")
	if len(req.PayPasswordMd5) <= 0 {
		response.FailWithMessage("支付密码格式错误！", c)
	}
	if err := mallUserService.ValidateUserPayPassword(token, req); err != nil {
		global.GVA_LOG.Error("验证失败", zap.Error(err))
		response.FailWithMessage("验证失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("验证成功", c)
}

// 返回用户银行卡号
func (m *MallUserApi) GetBankCard(c *gin.Context) {
	uidstr := c.Query("uid")
	uid, _ := strconv.Atoi(uidstr)
	bankCard, cardholder := mallUserService.GetBankCard(uid)
	if bankCard == "" {
		response.FailWithMessage("银行卡信息未设置", c)
		return
	}
	res := response2.MallCardInfoRes{
		BankCard:   bankCard,
		Cardhilder: cardholder,
	}
	response.OkWithData(res, c)
}
