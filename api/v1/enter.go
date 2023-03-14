package v1

import (
	"main.go/api/v1/common"
	"main.go/api/v1/mall"
	"main.go/api/v1/manage"
)

type ApiGroup struct {
	ManageApiGroup manage.ManageGroup
	MallApiGroup   mall.MallGroup
	CommonGroup    common.CommonGroup
}

var ApiGroupApp = new(ApiGroup)
