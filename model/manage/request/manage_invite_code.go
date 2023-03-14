package request

import (
	"main.go/model/common/request"
	"main.go/model/manage"
)

type MallInviteCodeSearch struct {
	manage.MallInviteCode
	request.PageInfo
}

type MallInviteCodeAddParam struct {
	InviteCode string `json:"inviteCode"`
}
