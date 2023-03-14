package manage

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"main.go/global"
	"main.go/model/common/request"
	"main.go/model/mall"
	"main.go/model/manage"
	manageReq "main.go/model/manage/request"
	"main.go/utils"
)

type ManageUserService struct {
}

// LockUser 修改用户状态
func (m *ManageUserService) LockUser(idReq request.IdsReq, lockStatus int) (err error) {
	if lockStatus != 0 && lockStatus != 1 {
		return errors.New("操作非法！")
	}
	//更新字段为0时，不能直接UpdateColumns
	err = global.GVA_DB.Model(&manage.MallUser{}).Where("user_id in ?", idReq.Ids).Update("locked_flag", lockStatus).Error
	return err
}

// GetMallUserInfoList 分页获取商城注册用户列表
func (m *ManageUserService) GetMallUserInfoList(info manageReq.MallUserSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNumber - 1)
	// 创建db
	db := global.GVA_DB.Model(&manage.MallUser{})
	var mallUsers []manage.MallUser
	lName := fmt.Sprintf("%%%s%%", info.LoginName)
	nName := fmt.Sprintf("%%%s%%", info.NickName)
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Where("login_name like ? and nick_name like ?", lName, nName).Count(&total).Error
	if err != nil {
		return
	}
	err = db.Where("login_name like ? and nick_name like ?", lName, nName).Limit(limit).Offset(offset).Order("create_time desc").Find(&mallUsers).Error
	return err, mallUsers, total
}

// UpdateUserCurrency 修改用户余额
func (m *ManageUserService) AddUserCurrency(addUserCurrency manageReq.AddUserCurrency) (err error) {
	if addUserCurrency.AddCurrency <= 0 || addUserCurrency.UserId <= 0 {
		return errors.New("操作非法！")
	}

	err = global.GVA_DB.Model(&manage.MallUser{}).Where("user_id = ?", addUserCurrency.UserId).
		Update("currency", gorm.Expr("currency + ?", addUserCurrency.AddCurrency)).Error
	return err
}

func (m *ManageUserService) DeleteUserCurrency(delUserCurrency manageReq.DeleteUserCurrency) (err error) {
	if delUserCurrency.DeleteCurrency <= 0 || delUserCurrency.UserId <= 0 {
		return errors.New("操作非法！")
	}

	var userCurrency manage.MallUser
	global.GVA_DB.Where("user_id =?", delUserCurrency.UserId).First(&userCurrency)
	if userCurrency.Currency < delUserCurrency.DeleteCurrency {
		return errors.New("用户余额不足！")
	}

	if err = global.GVA_DB.Where("user_id = ?", delUserCurrency.UserId).
		Updates(manage.MallUser{Currency: userCurrency.Currency - delUserCurrency.DeleteCurrency}).Error; err != nil {
		return errors.New("扣款失败！")
	}

	/*err = global.GVA_DB.Model(&manage.MallUser{}).Where("user_id = ?", delUserCurrency.UserId).
	Set("currency", userCurrency.Currency-delUserCurrency.DeleteCurrency).Error*/
	return nil
}

func (m *ManageUserService) ResetUserPassword(delUserCurrency manageReq.ResetUserPassword) (err error) {
	if delUserCurrency.UserId <= 0 {
		return errors.New("操作非法！")
	}

	err = global.GVA_DB.Model(&manage.MallUser{}).Where("user_id = ?", delUserCurrency.UserId).
		Update("password_md5", utils.MD5V([]byte("123456"))).Error
	return err
}

func (m *ManageUserService) ResetUserPayPassword(delUserCurrency manageReq.ResetUserPassword) (err error) {
	if delUserCurrency.UserId <= 0 {
		return errors.New("操作非法！")
	}

	err = global.GVA_DB.Model(&manage.MallUser{}).Where("user_id = ?", delUserCurrency.UserId).
		Update("pay_password_md5", utils.MD5V([]byte("123456"))).Error
	return err
}

func (m *ManageUserService) UpdateUserBankCard(req manageReq.UpdateUserBankCard) (err error) {
	var userInfo mall.MallUser
	err = global.GVA_DB.Where("user_id =?", req.UserId).First(&userInfo).Error
	if err != nil {
		return errors.New("不存在的用户")
	}
	if req.BankCard != "" {
		userInfo.BankCard = req.BankCard
	}
	if req.Cardhilder != "" {
		userInfo.Cardhilder = req.BankCard
	}
	err = global.GVA_DB.Where("user_id =?", req.UserId).UpdateColumns(&userInfo).Error
	return
}
