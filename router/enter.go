package router

import (
	"main.go/router/common"
	"main.go/router/mall"
	"main.go/router/manage"
)

type RouterGroup struct {
	Manage manage.ManageRouterGroup
	Mall   mall.MallRouterGroup
	Common common.CommonRouterGroup
}

var RouterGroupApp = new(RouterGroup)
