package manage

import (
	"main.go/global"
	"main.go/model/common"
	"main.go/model/common/request"
	"main.go/model/manage"
	manageReq "main.go/model/manage/request"
	"time"
)

type ManageInviteCodeService struct {
}

func (m *ManageInviteCodeService) CreateInviteCode(req manageReq.MallInviteCodeAddParam) (err error) {
	mallInviteCode := manage.MallInviteCode{
		InviteCode: req.InviteCode,
		CreateTime: common.JSONTime{Time: time.Now()},
	}
	err = global.GVA_DB.Create(&mallInviteCode).Error
	return err
}

func (m *ManageInviteCodeService) DeleteInviteCode(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&manage.MallInviteCode{}, "invite_code_id in ?", ids.Ids).Error
	return err
}

func (m *ManageInviteCodeService) GetInviteCodeList(info manageReq.MallInviteCodeSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNumber - 1)
	// 创建db
	db := global.GVA_DB.Model(&manage.MallInviteCode{})
	var mallInviteCodes []manage.MallInviteCode
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("invite_code_id desc").Find(&mallInviteCodes).Error
	return err, mallInviteCodes, total
}
