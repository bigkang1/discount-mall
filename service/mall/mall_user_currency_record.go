package mall

import (
	"errors"
	"main.go/global"
	"main.go/model/common"
	"main.go/model/mall"
	mallReq "main.go/model/mall/request"
	mallRes "main.go/model/mall/response"
	"main.go/model/manage"
	"time"
)

type MallUserCurrencyRecordService struct {
}

func (m *MallUserCurrencyRecordService) CreateUserCurrencyRecord(token string, req mallReq.AddUserCurrencyRecord) (err error) {
	var userToken mall.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户")
	}

	var user manage.MallUser
	if err = global.GVA_DB.Where("user_id =?", userToken.UserId).First(&user).Error; err != nil {
		global.GVA_DB.Rollback()
		return err
	}
	if user.Currency < req.CurrencyAmount {
		return errors.New("用户余额不足")
	}

	record := manage.MallUserCurrencyRecord{
		UserId:         userToken.UserId,
		CurrencyAmount: req.CurrencyAmount,
		CurrencyType:   1,
		Status:         2,
		CreateTime:     common.JSONTime{Time: time.Now()},
		UpdateTime:     common.JSONTime{Time: time.Now()},
	}

	err = global.GVA_DB.Create(&record).Error
	return err
}

func (m *MallUserCurrencyRecordService) UserCurrencyRecordList(token string, pageNumber, pageSize int) (err error,
	list []mallRes.UserCurrencyRecordListResponse, total int64) {
	var userToken mall.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户"), list, total
	}
	limit := pageSize
	offset := pageSize * (pageNumber - 1)

	var userCurrencyRecords []manage.MallUserCurrencyRecord
	db := global.GVA_DB.Model(&userCurrencyRecords)

	err = db.Where("user_id =?", userToken.UserId).Count(&total).Error

	err = db.Limit(limit).Offset(offset).Order("create_time desc").Find(&userCurrencyRecords).Error

	/*var user manage.MallUser
	err = global.GVA_DB.Where("user_id = ?", userToken.UserId).First(&user).Error
	if err != nil {
		return nil, list, total
	}*/

	if total > 0 {
		for _, userCurrencyRecord := range userCurrencyRecords {
			var res mallRes.UserCurrencyRecordListResponse
			res.UserCurrencyRecordId = userCurrencyRecord.UserCurrencyRecordId
			res.UserId = userCurrencyRecord.UserId
			//userStockResponse.NickName = user.NickName
			res.CurrencyAmount = userCurrencyRecord.CurrencyAmount
			res.CurrencyType = userCurrencyRecord.CurrencyType
			res.Status = userCurrencyRecord.Status
			res.AdminUserId = userCurrencyRecord.AdminUserId
			res.CreateTime = userCurrencyRecord.CreateTime
			res.UpdateTime = userCurrencyRecord.UpdateTime
			list = append(list, res)
		}
	}
	return err, list, total
}

func (m *MallUserCurrencyRecordService) IsExistPayPassword(token string) (isExistPayPassword bool, err error) {
	var userToken mall.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return false, errors.New("不存在的用户")
	}
	var user manage.MallUser
	err = global.GVA_DB.Model(&manage.MallUser{}).Where("user_id = ?", userToken.UserId).
		First(&user).Error
	if user.PayPasswordMd5 == "" {
		return false, nil
	}
	return true, nil
}
