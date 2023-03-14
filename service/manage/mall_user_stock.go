package manage

import (
	"fmt"
	"main.go/global"
	"main.go/model/common/request"
	mallRes "main.go/model/mall/response"
	"main.go/model/manage"
	managereq "main.go/model/manage/request"
	"strconv"
	"time"
)

type ManageUserStockService struct {
}

// MallOrderListBySearch 搜索订单
func (m *ManageUserStockService) MallOrderList(pageInfo request.PageInfo) (err error, list []mallRes.MallUserStockResponse, total int64) {
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.PageNumber - 1)
	var newBeeMallOrders []manage.MallUserStock
	db := global.GVA_DB.Model(&newBeeMallOrders)

	err = db.Where("goods_count > 0").Count(&total).Error

	err = db.Limit(limit).Offset(offset).Order(" user_stock_id desc").Find(&newBeeMallOrders).Error

	if total > 0 {
		for _, newBeeMallOrder := range newBeeMallOrders {
			var goodsInfo manage.MallGoodsInfo
			global.GVA_DB.Where("goods_id =?", newBeeMallOrder.GoodsId).First(&goodsInfo)
			var userStockResponse mallRes.MallUserStockResponse
			var sellingprice float64
			if goodsInfo.IsDiscount == 1 && int(time.Now().Unix()) < goodsInfo.DiscountEndTime {
				sellingprice, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", goodsInfo.SellingPrice*goodsInfo.OriginalPrice), 64)
			} else {
				sellingprice, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", goodsInfo.OriginalPrice), 64)
			}
			userStockResponse.UserStockId = newBeeMallOrder.UserStockId
			userStockResponse.UserId = newBeeMallOrder.UserId
			userStockResponse.GoodsCount = newBeeMallOrder.GoodsCount
			userStockResponse.CreateTime = newBeeMallOrder.CreateTime
			userStockResponse.GoodsId = goodsInfo.GoodsId
			userStockResponse.SellingPrice = sellingprice
			userStockResponse.GoodsName = goodsInfo.GoodsName
			userStockResponse.GoodsCoverImg = goodsInfo.GoodsCoverImg
			list = append(list, userStockResponse)
		}
	}
	return err, list, total
}

func (m *ManageUserStockService) MallOrderListBySearch(req managereq.MallUserStockSearch) (err error, list []mallRes.MallUserStockResponse, total int64) {
	limit := req.PageSize
	offset := req.PageSize * (req.PageNumber - 1)
	// 根据搜索条件查询
	var newBeeMallOrders []manage.MallUserStock
	db := global.GVA_DB.Model(&newBeeMallOrders)

	var users []manage.MallUser
	err = global.GVA_DB.Where("nick_name like ?", "%"+req.SearchName+"%").Find(&users).Error
	if err != nil {
		return nil, list, total
	}
	var userIDs []int
	for _, user := range users {
		userIDs = append(userIDs, user.UserId)
	}

	err = db.Where("user_id in ? and goods_count > 0", userIDs).Count(&total).Error

	err = db.Limit(limit).Offset(offset).Order(" user_stock_id desc").Find(&newBeeMallOrders).Error

	if total > 0 {
		for _, newBeeMallOrder := range newBeeMallOrders {
			var goodsInfo manage.MallGoodsInfo
			global.GVA_DB.Where("goods_id =?", newBeeMallOrder.GoodsId).First(&goodsInfo)
			var userStockResponse mallRes.MallUserStockResponse
			var sellingprice float64
			if goodsInfo.IsDiscount == 1 && int(time.Now().Unix()) < goodsInfo.DiscountEndTime {
				sellingprice, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", goodsInfo.SellingPrice*goodsInfo.OriginalPrice), 64)
			} else {
				sellingprice, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", goodsInfo.OriginalPrice), 64)
			}
			userStockResponse.UserStockId = newBeeMallOrder.UserStockId
			userStockResponse.UserId = newBeeMallOrder.UserId
			userStockResponse.GoodsCount = newBeeMallOrder.GoodsCount
			userStockResponse.CreateTime = newBeeMallOrder.CreateTime
			userStockResponse.GoodsId = goodsInfo.GoodsId
			userStockResponse.SellingPrice = sellingprice
			userStockResponse.GoodsName = goodsInfo.GoodsName
			userStockResponse.GoodsCoverImg = goodsInfo.GoodsCoverImg
			list = append(list, userStockResponse)
		}
	}
	return err, list, total
}

func (m *ManageUserStockService) MallOrderListBySearchUserID(req managereq.MallUserStockSearchUserId) (err error, list []mallRes.MallUserStockResponse, total int64) {
	limit := req.PageSize
	offset := req.PageSize * (req.PageNumber - 1)
	// 根据搜索条件查询
	var newBeeMallOrders []manage.MallUserStock
	db := global.GVA_DB.Model(&newBeeMallOrders)

	err = db.Where("user_id = ? and goods_count > 0", req.SearchUserId).Count(&total).Error

	err = db.Limit(limit).Offset(offset).Order(" user_stock_id desc").Find(&newBeeMallOrders).Error

	if total > 0 {
		for _, newBeeMallOrder := range newBeeMallOrders {
			var goodsInfo manage.MallGoodsInfo
			global.GVA_DB.Where("goods_id =?", newBeeMallOrder.GoodsId).First(&goodsInfo)
			var userStockResponse mallRes.MallUserStockResponse
			var sellingprice float64
			if goodsInfo.IsDiscount == 1 && int(time.Now().Unix()) < goodsInfo.DiscountEndTime {
				sellingprice, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", goodsInfo.SellingPrice*goodsInfo.OriginalPrice), 64)
			} else {
				sellingprice, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", goodsInfo.OriginalPrice), 64)
			}
			userStockResponse.UserStockId = newBeeMallOrder.UserStockId
			userStockResponse.UserId = newBeeMallOrder.UserId
			userStockResponse.GoodsCount = newBeeMallOrder.GoodsCount
			userStockResponse.CreateTime = newBeeMallOrder.CreateTime
			userStockResponse.GoodsId = goodsInfo.GoodsId
			userStockResponse.SellingPrice = sellingprice
			userStockResponse.GoodsName = goodsInfo.GoodsName
			userStockResponse.GoodsCoverImg = goodsInfo.GoodsCoverImg
			list = append(list, userStockResponse)
		}
	}
	return err, list, total
}
