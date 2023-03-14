package mall

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"main.go/global"
	"main.go/model/common"
	"main.go/model/common/enum"
	"main.go/model/mall"
	mallRes "main.go/model/mall/response"
	"main.go/model/manage"
	"strconv"
	"time"
)

type MallUserStockService struct {
}

// SaveOrder 保存订单
func (m *MallUserStockService) SaveOrder(token string, myShoppingCartItems []mallRes.CartItemResponse) (err error) {
	var userToken mall.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户")
	}
	var itemIdList []int
	var goodsIds []int
	for _, cartItem := range myShoppingCartItems {
		itemIdList = append(itemIdList, cartItem.CartItemId)
		goodsIds = append(goodsIds, cartItem.GoodsId)
	}
	var newBeeMallGoods []manage.MallGoodsInfo
	global.GVA_DB.Where("goods_id in ? ", goodsIds).Find(&newBeeMallGoods)
	//检查是否包含已下架商品
	for _, mallGoods := range newBeeMallGoods {
		if mallGoods.GoodsSellStatus != enum.GOODS_UNDER.Code() {
			return errors.New("已下架，无法生成订单")
		}
	}
	newBeeMallGoodsMap := make(map[int]manage.MallGoodsInfo)
	for _, mallGoods := range newBeeMallGoods {
		newBeeMallGoodsMap[mallGoods.GoodsId] = mallGoods
	}
	//判断商品库存
	for _, shoppingCartItemVO := range myShoppingCartItems {
		//查出的商品中不存在购物车中的这条关联商品数据，直接返回错误提醒
		if _, ok := newBeeMallGoodsMap[shoppingCartItemVO.GoodsId]; !ok {
			return errors.New("购物车数据异常！")
		}
		if shoppingCartItemVO.GoodsCount > float64(newBeeMallGoodsMap[shoppingCartItemVO.GoodsId].StockNum) {
			return errors.New("库存不足！" + shoppingCartItemVO.GoodsName)
		}
	}
	//删除购物项
	if len(itemIdList) > 0 && len(goodsIds) > 0 {
		// 遍历购物车，囤货
		for _, newBeeMallShoppingCartItemVO := range myShoppingCartItems {
			tx := global.GVA_DB.Begin()
			defer func() {
				if r := recover(); r != any(nil) {
					tx.Rollback()
				}
			}()
			if err = tx.Where("cart_item_id = ?", newBeeMallShoppingCartItemVO.CartItemId).Updates(mall.MallShoppingCartItem{IsDeleted: 1}).Error; err != nil {
				tx.Rollback()
				return errors.New("购物车删除失败！")
			}

			var userCurrency manage.MallUser
			tx.Where("user_id =?", userToken.UserId).First(&userCurrency)
			if userCurrency.Currency <= newBeeMallShoppingCartItemVO.GoodsCount*newBeeMallShoppingCartItemVO.SellingPrice {
				tx.Rollback()
				return errors.New("余额不足！")
			}
			// 扣除用户余额
			if err = tx.Where("user_id = ?", userToken.UserId).
				Updates(&manage.MallUser{Currency: userCurrency.Currency - newBeeMallShoppingCartItemVO.GoodsCount*newBeeMallShoppingCartItemVO.SellingPrice}).Error; err != nil {
				tx.Rollback()
				return err
			}

			// 减去商品库存
			var goodsInfo manage.MallGoodsInfo
			tx.Where("goods_id =?", newBeeMallShoppingCartItemVO.GoodsId).First(&goodsInfo)
			if err = tx.Where("goods_id =? and stock_num>= ? and goods_sell_status = 0", newBeeMallShoppingCartItemVO.GoodsId, newBeeMallShoppingCartItemVO.GoodsCount).
				Updates(manage.MallGoodsInfo{StockNum: goodsInfo.StockNum - int(newBeeMallShoppingCartItemVO.GoodsCount)}).Error; err != nil {
				tx.Rollback()
				return errors.New("库存不足！")
			}

			var userStock manage.MallUserStock
			err = tx.Where("user_id = ? and goods_id = ? and goods_count >= 0 ", userToken.UserId, newBeeMallShoppingCartItemVO.GoodsId).First(&userStock).Error
			if err != nil {
				// 不存在该商品囤货记录，添加一条
				if errors.Is(err, gorm.ErrRecordNotFound) {
					if err = tx.Save(&manage.MallUserStock{
						UserId:     userToken.UserId,
						GoodsId:    newBeeMallShoppingCartItemVO.GoodsId,
						GoodsCount: int(newBeeMallShoppingCartItemVO.GoodsCount),
						CreateTime: common.JSONTime{Time: time.Now()},
					}).Error; err != nil {
						tx.Rollback()
						return err
					}
				} else {
					tx.Rollback()
					return errors.New("囤货失败")
				}
			} else {
				// 有该商品囤货记录，添加商品数量即可
				if err = tx.Where("user_id = ? and goods_id = ? and goods_count >= 0 ", userToken.UserId, newBeeMallShoppingCartItemVO.GoodsId).
					Updates(&manage.MallUserStock{GoodsCount: userStock.GoodsCount + int(newBeeMallShoppingCartItemVO.GoodsCount)}).Error; err != nil {
					tx.Rollback()
					return err
				}
			}
			tx.Commit()
		}

	}
	return
}

// MallOrderListBySearch 搜索订单
func (m *MallUserStockService) MallOrderListBySearch(token string, pageNumber int) (err error, list []mallRes.MallUserStockResponse, total int64) {
	var userToken mall.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户"), list, total
	}
	// 根据搜索条件查询
	var newBeeMallOrders []manage.MallUserStock
	db := global.GVA_DB.Model(&newBeeMallOrders)

	err = db.Where("user_id =? and goods_count > 0", userToken.UserId).Count(&total).Error
	//这里前段没有做滚动加载，直接显示全部订单
	//limit := 5
	offset := 5 * (pageNumber - 1)
	err = db.Offset(offset).Order(" user_stock_id desc").Find(&newBeeMallOrders).Error

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

func (m *MallUserStockService) DeleteUserStock(token string, goodsId, goodsCount int) (err error) {
	var userToken mall.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户")
	}

	var newBeeMallGoods manage.MallGoodsInfo
	global.GVA_DB.Where("goods_id = ? ", goodsId).Find(&newBeeMallGoods)
	//检查是否包含已下架商品
	if newBeeMallGoods.GoodsSellStatus != enum.GOODS_UNDER.Code() {
		return errors.New("已下架，无法生成订单回收")
	}

	//回收
	tx := global.GVA_DB.Begin()
	defer func() {
		if r := recover(); r != any(nil) {
			tx.Rollback()
		}
	}()

	var userCurrency manage.MallUser
	tx.Where("user_id =?", userToken.UserId).First(&userCurrency)

	// 添加用户余额
	if err = tx.Where("user_id = ?", userToken.UserId).
		Updates(&manage.MallUser{Currency: userCurrency.Currency + float64(goodsCount)*newBeeMallGoods.OriginalPrice}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 添加商品库存
	var goodsInfo manage.MallGoodsInfo
	tx.Where("goods_id =?", goodsId).First(&goodsInfo)
	if err = tx.Where("goods_id =? and stock_num>= ? and goods_sell_status = 0", goodsId, goodsCount).
		Updates(manage.MallGoodsInfo{StockNum: goodsInfo.StockNum + int(goodsCount)}).Error; err != nil {
		tx.Rollback()
		return errors.New("回收失败，添加商品库存失败！")
	}

	var userStock manage.MallUserStock
	err = tx.Where("user_id = ? and goods_id = ? and goods_count > 0 ", userToken.UserId, goodsId).First(&userStock).Error
	if err != nil {
		tx.Rollback()
		return errors.New("无商品囤货记录，回收失败")
	}

	if userStock.GoodsCount < int(goodsCount) {
		tx.Rollback()
		return errors.New("你的货不足，回收失败！")
	} else if userStock.GoodsCount > int(goodsCount) {
		// 有该商品囤货记录，减少商品数量即可
		if err = tx.Where("user_id = ? and goods_id = ? and goods_count > 0 ", userToken.UserId, goodsId).
			Updates(&manage.MallUserStock{GoodsCount: userStock.GoodsCount - int(goodsCount)}).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		// 用户货物数量=回收数量，将数量改为0
		if err = tx.Model(&manage.MallUserStock{}).Where("user_id = ? and goods_id = ? ", userToken.UserId, goodsId).
			Update("goods_count", 0).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return
}
