package manage

import (
	"errors"
	"main.go/global"
	mallRes "main.go/model/mall/response"
	"main.go/model/manage"
	manageReq "main.go/model/manage/request"
	"strconv"
)

type MallUserCurrencyRecordService struct {
}

func (m *MallUserCurrencyRecordService) UserCurrencyRecordList(pageNumber, pageSize int) (err error,
	list []mallRes.UserCurrencyRecordListResponse, total int64) {

	limit := pageSize
	offset := pageSize * (pageNumber - 1)

	var userCurrencyRecords []manage.MallUserCurrencyRecord
	db := global.GVA_DB.Model(&userCurrencyRecords)

	err = db.Count(&total).Error

	err = db.Limit(limit).Offset(offset).Order("create_time desc").Find(&userCurrencyRecords).Error

	/*var userIds []int
	for _, userCurrencyRecord := range userCurrencyRecords {
		userIds = append(userIds,userCurrencyRecord.UserId)
	}
	var user manage.MallUser
	err = global.GVA_DB.Where("user_id in ?", userIds).First(&user).Error
	if err != nil {
		return nil, list, total
	}*/

	if total > 0 {
		for _, userCurrencyRecord := range userCurrencyRecords {
			var res mallRes.UserCurrencyRecordListResponse
			res.UserCurrencyRecordId = userCurrencyRecord.UserCurrencyRecordId
			res.UserId = userCurrencyRecord.UserId
			var user manage.MallUser
			_ = global.GVA_DB.Where("user_id = ?", userCurrencyRecord.UserId).First(&user).Error
			res.NickName = user.NickName
			res.LoginName = user.LoginName
			res.CurrencyAmount = userCurrencyRecord.CurrencyAmount
			res.CurrencyType = userCurrencyRecord.CurrencyType
			res.Status = userCurrencyRecord.Status
			res.AdminUserId = userCurrencyRecord.AdminUserId
			var adminUser manage.MallAdminUser
			_ = global.GVA_DB.Where("admin_user_id = ?", userCurrencyRecord.AdminUserId).First(&adminUser).Error
			res.AdminNickName = adminUser.NickName
			res.CreateTime = userCurrencyRecord.CreateTime
			res.UpdateTime = userCurrencyRecord.UpdateTime
			list = append(list, res)
		}
	}
	return err, list, total
}

func (m *MallUserCurrencyRecordService) UserCurrencyRecordListByUserName(searchName string, status, pageNumber, pageSize int) (err error,
	list []mallRes.UserCurrencyRecordListResponse, total int64) {

	limit := pageSize
	offset := pageSize * (pageNumber - 1)

	var userCurrencyRecords []manage.MallUserCurrencyRecord
	db := global.GVA_DB.Model(&userCurrencyRecords)

	var users []manage.MallUser
	err = global.GVA_DB.Where("nick_name like ? or login_name like ?", "%"+searchName+"%", "%"+searchName+"%").
		Find(&users).Error
	if err != nil {
		return nil, list, total
	}
	var userIDs []int
	for _, user := range users {
		userIDs = append(userIDs, user.UserId)
	}

	if status >= 0 {
		db.Where("status = ?", status)
	}

	err = db.Where("user_id in ?", userIDs).Count(&total).Error

	err = db.Limit(limit).Offset(offset).Order("create_time desc").Find(&userCurrencyRecords).Error

	if total > 0 {
		for _, userCurrencyRecord := range userCurrencyRecords {
			var res mallRes.UserCurrencyRecordListResponse
			res.UserCurrencyRecordId = userCurrencyRecord.UserCurrencyRecordId
			res.UserId = userCurrencyRecord.UserId
			var user manage.MallUser
			_ = global.GVA_DB.Where("user_id = ?", userCurrencyRecord.UserId).First(&user).Error
			res.NickName = user.NickName
			res.LoginName = user.LoginName
			res.CurrencyAmount = userCurrencyRecord.CurrencyAmount
			res.CurrencyType = userCurrencyRecord.CurrencyType
			res.Status = userCurrencyRecord.Status
			res.AdminUserId = userCurrencyRecord.AdminUserId
			var adminUser manage.MallAdminUser
			_ = global.GVA_DB.Where("admin_user_id = ?", userCurrencyRecord.AdminUserId).First(&adminUser).Error
			res.AdminNickName = adminUser.NickName
			res.CreateTime = userCurrencyRecord.CreateTime
			res.UpdateTime = userCurrencyRecord.UpdateTime
			list = append(list, res)
		}
	}
	return err, list, total
}

func (m *MallUserCurrencyRecordService) UserCurrencyRecordListByUserLoginName(searchLoginName string, pageNumber, pageSize int) (err error,
	list []mallRes.UserCurrencyRecordListResponse, total int64) {

	limit := pageSize
	offset := pageSize * (pageNumber - 1)

	var userCurrencyRecords []manage.MallUserCurrencyRecord
	db := global.GVA_DB.Model(&userCurrencyRecords)

	var users []manage.MallUser
	err = global.GVA_DB.Where("login_name like ?", "%"+searchLoginName+"%").Find(&users).Error
	if err != nil {
		return nil, list, total
	}
	var userIDs []int
	for _, user := range users {
		userIDs = append(userIDs, user.UserId)
	}

	err = db.Where("user_id in ?", userIDs).Count(&total).Error

	err = db.Limit(limit).Offset(offset).Order("create_time desc").Find(&userCurrencyRecords).Error

	if total > 0 {
		for _, userCurrencyRecord := range userCurrencyRecords {
			var res mallRes.UserCurrencyRecordListResponse
			res.UserCurrencyRecordId = userCurrencyRecord.UserCurrencyRecordId
			res.UserId = userCurrencyRecord.UserId
			var user manage.MallUser
			_ = global.GVA_DB.Where("user_id = ?", userCurrencyRecord.UserId).First(&user).Error
			res.NickName = user.NickName
			res.LoginName = user.LoginName
			res.CurrencyAmount = userCurrencyRecord.CurrencyAmount
			res.CurrencyType = userCurrencyRecord.CurrencyType
			res.Status = userCurrencyRecord.Status
			res.AdminUserId = userCurrencyRecord.AdminUserId
			var adminUser manage.MallAdminUser
			_ = global.GVA_DB.Where("admin_user_id = ?", userCurrencyRecord.AdminUserId).First(&adminUser).Error
			res.AdminNickName = adminUser.NickName
			res.CreateTime = userCurrencyRecord.CreateTime
			res.UpdateTime = userCurrencyRecord.UpdateTime
			list = append(list, res)
		}
	}
	return err, list, total
}

func (m *MallUserCurrencyRecordService) UserCurrencyRecordListByStatus(status, pageNumber, pageSize int) (err error,
	list []mallRes.UserCurrencyRecordListResponse, total int64) {
	limit := pageSize
	offset := pageSize * (pageNumber - 1)

	var userCurrencyRecords []manage.MallUserCurrencyRecord
	db := global.GVA_DB.Model(&userCurrencyRecords)

	err = db.Where("status = ?", status).Count(&total).Error

	err = db.Limit(limit).Offset(offset).Order("create_time desc").Find(&userCurrencyRecords).Error

	if total > 0 {
		for _, userCurrencyRecord := range userCurrencyRecords {
			var res mallRes.UserCurrencyRecordListResponse
			res.UserCurrencyRecordId = userCurrencyRecord.UserCurrencyRecordId
			res.UserId = userCurrencyRecord.UserId
			var user manage.MallUser
			_ = global.GVA_DB.Where("user_id = ?", userCurrencyRecord.UserId).First(&user).Error
			res.NickName = user.NickName
			res.LoginName = user.LoginName
			res.CurrencyAmount = userCurrencyRecord.CurrencyAmount
			res.CurrencyType = userCurrencyRecord.CurrencyType
			res.Status = userCurrencyRecord.Status
			res.AdminUserId = userCurrencyRecord.AdminUserId
			var adminUser manage.MallAdminUser
			_ = global.GVA_DB.Where("admin_user_id = ?", userCurrencyRecord.AdminUserId).First(&adminUser).Error
			res.AdminNickName = adminUser.NickName
			res.CreateTime = userCurrencyRecord.CreateTime
			res.UpdateTime = userCurrencyRecord.UpdateTime
			list = append(list, res)
		}
	}
	return err, list, total
}

func (m *MallUserCurrencyRecordService) UpdateUserCurrencyRecordStatus(req manageReq.UpdateCurrencyRecordStatus) (err error) {
	if req.Status != 0 && req.Status != 1 && req.Status != 2 {
		return errors.New("status字段错误")
	}
	err = global.GVA_DB.Model(&manage.MallUserCurrencyRecord{}).
		Where(" user_currency_record_id = ?", req.UserCurrencyRecordId).
		Set("status", req.Status).Error
	return err
}

func (m *MallUserCurrencyRecordService) AccessUserCurrency(req manageReq.AccessUserCurrency) (loginName, nickName string, err error) {
	if req.UserCurrencyRecordId <= 0 {
		return "", "", errors.New("userCurrencyRecordId字段错误")
	}

	var userCurrencyRecord manage.MallUserCurrencyRecord
	err = global.GVA_DB.Where("user_currency_record_id = ?", req.UserCurrencyRecordId).First(&userCurrencyRecord).Error
	if err != nil {
		return "", "", errors.New("查询不到该记录：" + strconv.Itoa(req.UserCurrencyRecordId))
	}

	var user manage.MallUser
	if err = global.GVA_DB.Where("user_id =?", userCurrencyRecord.UserId).First(&user).Error; err != nil {
		return "", "", err
	}

	tx := global.GVA_DB.Begin()
	defer func() {
		if r := recover(); r != any(nil) {
			tx.Rollback()
		}
	}()

	if user.Currency < userCurrencyRecord.CurrencyAmount {
		err = tx.Where("user_currency_record_id = ? and currency_type = 1", req.UserCurrencyRecordId).
			Updates(&manage.MallUserCurrencyRecord{Status: 3}).Error
		if err != nil {
			tx.Rollback()
			return "", "", err
		}
		return user.LoginName, user.NickName, errors.New("用户余额不足")
	}
	// 添加用户余额
	if err = tx.Where("user_id = ?", userCurrencyRecord.UserId).
		Updates(&manage.MallUser{Currency: user.Currency - userCurrencyRecord.CurrencyAmount}).Error; err != nil {
		tx.Rollback()
		return user.LoginName, user.NickName, err
	}

	err = tx.Where(" user_currency_record_id = ? and currency_type = 1", req.UserCurrencyRecordId).
		Updates(&manage.MallUserCurrencyRecord{Status: 1}).Error
	if err != nil {
		tx.Rollback()
		return user.LoginName, user.NickName, err
	}
	tx.Commit()
	return user.LoginName, user.NickName, err
}
