package common

import (
	"fmt"
	"main.go/global"
	"main.go/model/common"
	"main.go/model/common/request"
	"main.go/model/manage"
	manageReq "main.go/model/manage/request"
	"time"
)

type CommonChartService struct {
}

// 添加消息
func (c *CommonChartService) AddMsg(cmsg *request.ChartMsg, gid, sid, isAdmi int) error {
	if isAdmi == 0 {
		gid = sid
	}
	chartmsg := &common.Charts{
		SendId:      sid,
		GetId:       gid,
		IsSendAdmin: isAdmi,
		Msg:         cmsg.Msg,
		MsgType:     cmsg.MsgType,
		SendTime:    int(time.Now().Unix()),
	}
	return global.GVA_DB.Create(chartmsg).Error
}

// 设置用户未读消息
func (c *CommonChartService) AddUUnRead(uid int) error {
	var unRead common.UnRead
	n := global.GVA_DB.Where("uid = ?", uid).First(&unRead).RowsAffected
	if n == 0 {
		unSetRead := &common.UnRead{
			Uid:     uid,
			UUnRead: 1,
			AUnRead: 0,
		}
		global.GVA_DB.Create(unSetRead)
	} else {
		global.GVA_DB.Model(&common.UnRead{}).Where("uid = ?", uid).Update("u_un_read", unRead.UUnRead+1)
	}
	return nil
}

// 设置管理员未读消息
func (c *CommonChartService) AddAUnRead(uid int) error {
	var unRead common.UnRead
	n := global.GVA_DB.Where("uid = ?", uid).First(&unRead).RowsAffected
	if n == 0 {
		unSetRead := &common.UnRead{
			Uid:     uid,
			UUnRead: 0,
			AUnRead: 1,
		}
		global.GVA_DB.Create(unSetRead)
	} else {
		global.GVA_DB.Model(&common.UnRead{}).Where("uid = ?", uid).Update("a_un_read", unRead.AUnRead+1)
	}
	return nil
}

// 清理管理员未读
func (c *CommonChartService) ClearAUnRead(uid int) error {
	global.GVA_DB.Model(&common.UnRead{}).Where("uid = ?", uid).Update("a_un_read", 0)
	return nil
}

// 清理用户未读
func (c *CommonChartService) ClearUUnRead(uid int) error {
	global.GVA_DB.Model(&common.UnRead{}).Where("uid = ?", uid).Update("u_un_read", 0)
	return nil
}

// 管理端查询用户聊天列表
func (c *CommonChartService) GetUserList(info manageReq.MallUserSearch) (error, []manage.MallUser, int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNumber - 1)
	var mallUsers []manage.MallUser
	var total int64

	lName := fmt.Sprintf("%%%s%%", info.LoginName)
	nName := fmt.Sprintf("%%%s%%", info.NickName)
	// 如果有条件搜索 下方会自动创建搜索语句
	err := global.GVA_DB.Model(&manage.MallUser{}).Where("login_name like ? and nick_name like ?", lName, nName).Count(&total).Error
	if err != nil {
		return err, nil, 0
	}
	err = global.GVA_DB.Select("user_id", "nick_name", "login_name").Where("login_name like ? and nick_name like ?", lName, nName).Offset(offset).Limit(limit).Find(&mallUsers).Error
	return err, mallUsers, total
}

func (c *CommonChartService) GetUserunread(uid int) *common.UnRead {
	var unRead common.UnRead
	global.GVA_DB.Select("a_un_read").Where("uid = ?", uid).First(&unRead)
	return &unRead
}
func (c *CommonChartService) UserGetunread(uid int) (error, int) {
	unRead := common.UnRead{}
	err := global.GVA_DB.Select("u_un_read").Where("uid = ?", uid).First(&unRead).Error
	return err, unRead.UUnRead
}

// 获取用户端聊天记录
func (c *CommonChartService) GetHistoryUserChart(uid int) (error, []common.Charts) {
	var charts []common.Charts
	err := global.GVA_DB.Where("sendId = ? or getId=?", uid, uid).Order("sendTime desc").Find(&charts).Error
	return err, charts
}

// 聊天记录
func (c *CommonChartService) ClearHistory() error {
	var charts common.Charts
	now := time.Now().AddDate(0, 0, -7).Unix()
	return global.GVA_DB.Where("sendTime < ", now).Unscoped().Delete(&charts).Error
}

// 返回管理员姓名
func (c *CommonChartService) GetAdminName(aid int) string {
	var charts manage.MallAdminUser
	global.GVA_DB.Select("nick_name").Where("admin_user_id = ?", aid).First(&charts)
	return charts.NickName
}
func (a *CommonChartService) UpdateLivingMsg(info common.AdminLivingMsg) error {

	return global.GVA_DB.Select("livemsg", "start_time", "ent_time", "is_use").Where("is_use = ? or is_use = ?", 1, 0).Updates(&info).Error
}
func (a *CommonChartService) GetLivingMsg() common.AdminLivingMsg {
	var msg common.AdminLivingMsg
	global.GVA_DB.First(&msg)
	return msg
}
