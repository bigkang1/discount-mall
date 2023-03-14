package mall

import (
	"errors"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"main.go/global"
	"main.go/model/common"
	"main.go/model/mall"
	mallReq "main.go/model/mall/request"
	mallRes "main.go/model/mall/response"
	"main.go/model/manage"
	"main.go/utils"
	"strconv"
	"strings"
	"time"
)

type MallUserService struct {
}

// RegisterUser 注册用户
func (m *MallUserService) RegisterUser(req mallReq.RegisterUserParam) (err error) {
	//判断邀请码是否存在
	err = global.GVA_DB.Where("invite_code =?", req.InviteCode).First(&manage.MallInviteCode{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("邀请码不存在")
		}
		return err
	}

	if !errors.Is(global.GVA_DB.Where("login_name =?", req.LoginName).First(&mall.MallUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("该用户名已注册")
	}

	if !errors.Is(global.GVA_DB.Where("nick_name =?", req.NickNam).First(&mall.MallUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同用户名")
	}

	return global.GVA_DB.Create(&mall.MallUser{
		NickName:      req.NickNam,
		LoginName:     req.LoginName,
		PasswordMd5:   utils.MD5V([]byte(req.Password)),
		IntroduceSign: "请添加自己的介绍",
		CreateTime:    common.JSONTime{Time: time.Now()},
	}).Error

}

func (m *MallUserService) UpdateUserInfo(token string, req mallReq.UpdateUserInfoParam) (err error) {
	var userToken mall.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户")
	}
	var userInfo mall.MallUser
	err = global.GVA_DB.Where("user_id =?", userToken.UserId).First(&userInfo).Error
	// 若密码为空字符，则表明用户不打算修改密码，使用原密码保存
	if !(req.PasswordMd5 == "") {
		if userInfo.PasswordMd5 != utils.MD5V([]byte(req.RawPassword)) {
			return errors.New("原密码不正确")
		}
		userInfo.PasswordMd5 = utils.MD5V([]byte(req.PasswordMd5))
	}
	userInfo.NickName = req.NickName
	userInfo.IntroduceSign = req.IntroduceSign
	err = global.GVA_DB.Where("user_id =?", userToken.UserId).UpdateColumns(&userInfo).Error
	return
}

func (m *MallUserService) GetUserDetail(token string) (err error, userDetail mallRes.MallUserDetailResponse) {
	var userToken mall.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户"), userDetail
	}
	var userInfo mall.MallUser
	err = global.GVA_DB.Where("user_id =?", userToken.UserId).First(&userInfo).Error
	if err != nil {
		return errors.New("用户信息获取失败"), userDetail
	}
	err = copier.Copy(&userDetail, &userInfo)
	return
}

func (m *MallUserService) UserLogin(params mallReq.UserLoginParam) (err error, user mall.MallUser, userToken mall.MallUserToken) {
	err = global.GVA_DB.Where("login_name=? AND password_md5=?", params.LoginName, params.PasswordMd5).First(&user).Error
	if user != (mall.MallUser{}) {
		token := getNewToken(time.Now().UnixNano()/1e6, int(user.UserId))
		global.GVA_DB.Where("user_id", user.UserId).First(&token)
		nowDate := time.Now()
		// 48小时过期
		expireTime, _ := time.ParseDuration("48h")
		expireDate := nowDate.Add(expireTime)
		// 没有token新增，有token 则更新
		if userToken == (mall.MallUserToken{}) {
			userToken.UserId = user.UserId
			userToken.Token = token
			userToken.UpdateTime = nowDate
			userToken.ExpireTime = expireDate
			if err = global.GVA_DB.Save(&userToken).Error; err != nil {
				return
			}
		} else {
			userToken.Token = token
			userToken.UpdateTime = nowDate
			userToken.ExpireTime = expireDate
			if err = global.GVA_DB.Save(&userToken).Error; err != nil {
				return
			}
		}
	}
	return err, user, userToken
}

func getNewToken(timeInt int64, userId int) (token string) {
	var build strings.Builder
	build.WriteString(strconv.FormatInt(timeInt, 10))
	build.WriteString(strconv.Itoa(userId))
	build.WriteString(utils.GenValidateCode(6))
	return utils.MD5V([]byte(build.String()))
}

func (m *MallUserService) GetUser(uid int) float64 {
	var userinfo manage.MallUser
	global.GVA_DB.Select("currency").Where("user_id = ?", uid).First(&userinfo)
	return userinfo.Currency
}

func (m *MallUserService) UpdateUserBankCard(token string, req mallReq.UpdateUserBankCard) (err error) {
	var userToken mall.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户")
	}

	var userInfo mall.MallUser
	err = global.GVA_DB.Where("user_id =?", userToken.UserId).First(&userInfo).Error
	if !(req.PayPassword == "") {
		userInfo.PayPasswordMd5 = utils.MD5V([]byte(req.PayPassword))
	}
	if req.BankCard != "" {
		userInfo.BankCard = req.BankCard
	}
	if req.Cardhilder != "" {
		userInfo.Cardhilder = req.BankCard
	}
	err = global.GVA_DB.Where("user_id =?", userToken.UserId).UpdateColumns(&userInfo).Error
	return
}

func (m *MallUserService) ValidateUserPayPassword(token string, req mallReq.ValidateUserPayPassword) (err error) {
	var userToken mall.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户")
	}

	var userInfo mall.MallUser
	err = global.GVA_DB.Where("user_id = ? and pay_password_md5 = ?", userToken.UserId, req.PayPasswordMd5).
		First(&userInfo).Error
	return
}

// 支付密码验证
func (m *MallUserService) VerificatePayPasswd(uid int, payPasswd string) error {
	var user mall.MallUser
	err := global.GVA_DB.Select("pay_password_md5").Where("user_id =?", uid).First(&user).Error
	if err != nil || user.PayPasswordMd5 != payPasswd {
		return errors.New("用户支付密码验证失败")
	}
	return nil
}

func (m *MallUserService) VerificateMoney(uid int, money float64) error {
	corrent := m.GetUser(uid)
	if corrent < money {
		return errors.New("用户支付余额不够")
	}
	return nil
}

// 回收修改用户余额
func (m *MallUserService) RecoverMoney(uid int, sellingPrice, orginPrice float64, endTime int) error {
	var user mall.MallUser
	corrent := m.GetUser(uid)
	if corrent < sellingPrice {
		return errors.New("用户支付余额不够")
	}
	global.GVA_DB.Select("currency").Where("user_id =?", uid).First(&user)
	err := global.GVA_DB.Model(&mall.MallUser{}).Where("user_id =?", uid).Update("currency", sellingPrice).Error
	//添加定时
	time.AfterFunc(time.Duration(endTime-int(time.Now().Unix()))*time.Second, func() {
		global.GVA_DB.Model(&mall.MallUser{}).Where("user_id =?", uid).Update("currency", orginPrice)
	})
	return err
}

// 获取用户银行卡号
func (m *MallUserService) GetBankCard(uid int) (string, string) {
	var user mall.MallUser
	global.GVA_DB.Select("bank_card,cardholder").Where("user_id =?", uid).First(&user)
	return user.BankCard, user.Cardhilder
}
