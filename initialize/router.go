package initialize

import (
	"github.com/gin-gonic/gin"
	"main.go/global"
	"main.go/middleware"
	"main.go/router"
	"net/http"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	//Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.GVA_LOG.Info("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors()) // 如需跨域可以打开
	global.GVA_LOG.Info("use middleware cors")
	// 方便统一添加路由组前缀 多服务器上线使用
	//商城后管路由
	manageRouter := router.RouterGroupApp.Manage
	ManageGroup := Router.Group("admin-api")
	PublicGroup := Router.Group("")

	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{
		//商城后管路由初始化
		manageRouter.InitManageAdminUserRouter(ManageGroup)     //管理员用户
		manageRouter.InitManageGoodsCategoryRouter(ManageGroup) //分类
		manageRouter.InitManageGoodsInfoRouter(ManageGroup)     //商品
		manageRouter.InitManageCarouselRouter(ManageGroup)
		manageRouter.InitManageIndexConfigRouter(ManageGroup)
		manageRouter.InitManageOrderRouter(ManageGroup)
		manageRouter.InitManageInviteCodeRouter(ManageGroup)
		manageRouter.InitManageUserStockRouter(ManageGroup)
		manageRouter.InitMallUserCurrencyRecordRouter(ManageGroup)
	}
	//商城前端路由
	mallRouter := router.RouterGroupApp.Mall
	MallGroup := Router.Group("front-api")
	{
		// 商城前端路由
		mallRouter.InitMallCarouselIndexRouter(MallGroup)
		mallRouter.InitMallGoodsInfoIndexRouter(MallGroup)
		mallRouter.InitMallGoodsCategoryIndexRouter(MallGroup)
		mallRouter.InitMallUserRouter(MallGroup)
		mallRouter.InitMallUserAddressRouter(MallGroup)
		mallRouter.InitMallShopCartRouter(MallGroup)
		mallRouter.InitMallOrderRouter(MallGroup)
		mallRouter.InitMallUserStockRouter(MallGroup)
		mallRouter.InitMallUserCurrencyRecordRouter(MallGroup)
	}

	//公共路由
	commmonRouter := router.RouterGroupApp.Common
	commonGroup := Router.Group("common")
	{
		commmonRouter.InitCommonChartRouter(commonGroup)
	}
	global.GVA_LOG.Info("router register success")
	return Router
}
