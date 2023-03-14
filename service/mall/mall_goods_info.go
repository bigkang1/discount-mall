package mall

import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"main.go/global"
	"main.go/model/mall"
	"main.go/model/mall/request"
	mallRes "main.go/model/mall/response"
	"main.go/model/manage"
	"main.go/utils"
	"strconv"
	"time"
)

type MallGoodsInfoService struct {
}

// MallGoodsListBySearch 商品搜索分页
func (m *MallGoodsInfoService) MallGoodsListBySearch(pageNumber int, isdiscount int, goodsCategoryId int, keyword string, orderBy string) (err error, searchGoodsList []mallRes.GoodsSearchResponse, total int64) {
	// 根据搜索条件查询
	var goodsList []manage.MallGoodsInfo
	db := global.GVA_DB.Model(&manage.MallGoodsInfo{})
	if keyword != "" {
		db.Where("goods_name like ? or goods_intro like ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	if isdiscount == 1 {
		db.Where("is_discount = ? and discount_end_time > ?", isdiscount, int(time.Now().Unix()))
	}
	if goodsCategoryId > 0 {
		db.Where("goods_category_id= ?", goodsCategoryId)
	}
	err = db.Count(&total).Error
	switch orderBy {
	case "new":
		db.Order("goods_id desc")
	case "price":
		db.Order("selling_price asc")
	default:
		db.Order("stock_num desc")
	}
	limit := 10
	offset := 10 * (pageNumber - 1)
	err = db.Limit(limit).Offset(offset).Find(&goodsList).Error
	// 返回查询结果
	for _, goods := range goodsList {
		sellingprice := goods.OriginalPrice
		if goods.IsDiscount == 1 && int(time.Now().Unix()) < goods.DiscountEndTime {
			sellingprice, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", goods.SellingPrice*goods.OriginalPrice), 64)
		} else {
			sellingprice, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", goods.OriginalPrice), 64)
		}
		searchGoods := mallRes.GoodsSearchResponse{
			GoodsId:       goods.GoodsId,
			GoodsName:     utils.SubStrLen(goods.GoodsName, 28),
			GoodsIntro:    utils.SubStrLen(goods.GoodsIntro, 28),
			GoodsCoverImg: goods.GoodsCoverImg,
			OriginalPrice: goods.OriginalPrice,
			SellingPrice:  sellingprice,
		}
		searchGoodsList = append(searchGoodsList, searchGoods)
	}
	return
}

// GetMallGoodsInfo 获取商品信息
func (m *MallGoodsInfoService) GetMallGoodsInfo(id int) (err error, res mallRes.GoodsInfoDetailResponse) {
	var mallGoodsInfo manage.MallGoodsInfo
	err = global.GVA_DB.Where("goods_id = ?", id).First(&mallGoodsInfo).Error
	if mallGoodsInfo.GoodsSellStatus != 0 {
		return errors.New("商品已下架"), mallRes.GoodsInfoDetailResponse{}
	}
	err = copier.Copy(&res, &mallGoodsInfo)
	if err != nil {
		return err, mallRes.GoodsInfoDetailResponse{}
	}
	var list []string
	list = append(list, mallGoodsInfo.GoodsCarousel)

	if mallGoodsInfo.IsDiscount == 1 && int(time.Now().Unix()) < mallGoodsInfo.DiscountEndTime {
		res.SellingPrice, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", mallGoodsInfo.SellingPrice*mallGoodsInfo.OriginalPrice), 64)
	} else {
		res.SellingPrice, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", mallGoodsInfo.OriginalPrice), 64)
	}
	if len(list) == 0 {
		list = append(list, mallGoodsInfo.GoodsCoverImg)
	}
	res.GoodsCarouselList = list

	return
}

// 查询商品名
func (m *MallGoodsInfoService) GetGoodName(gid int) string {
	var mallGoodsInfo manage.MallGoodsInfo
	global.GVA_DB.Where("goods_id = ?", gid).First(&mallGoodsInfo)
	return mallGoodsInfo.GoodsName
}

// 添加商品回收信息
func (m *MallGoodsInfoService) AddRecoveryInfo(recque request.GoodsRecovery) (float64, float64, int, error) {
	var mallGoodsInfo manage.MallGoodsInfo
	err := global.GVA_DB.Where("goods_id = ?", recque.GoodsId).First(&mallGoodsInfo).Error
	var sellingPrice float64
	originalPrice, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", mallGoodsInfo.OriginalPrice), 64)
	if mallGoodsInfo.IsDiscount == 1 && mallGoodsInfo.DiscountEndTime > int(time.Now().Unix()) {
		sellingPrice, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", mallGoodsInfo.SellingPrice*mallGoodsInfo.OriginalPrice), 64)
	} else {
		sellingPrice = originalPrice
	}

	recoverinfo := mall.Recovery_history{
		UserId:   recque.UserId,
		GoodsId:  recque.GoodsId,
		PayPrice: recque.PayPrice,
		RePrice:  originalPrice - sellingPrice,
		GoodsNum: recque.GoodsNum,
		ReTime:   int(time.Now().Unix()),
	}
	err = global.GVA_DB.Create(recoverinfo).Error
	return sellingPrice * float64(recque.GoodsNum), originalPrice * float64(recque.GoodsNum), mallGoodsInfo.DiscountEndTime, err
}

// 删除商品回收
func (m *MallGoodsInfoService) DeleteRecoveryInfo(uid int) {
	var RecoveryInfo mall.Recovery_history
	global.GVA_DB.Where("user_id = ? and re_time < ?", uid, int(time.Now().AddDate(0, 0, -15).Unix())).Delete(&RecoveryInfo)
}

// 分页查找用户回收信息
func (m *MallGoodsInfoService) GetRecoveryInfo(info request.GetGoodsInfo) (recoveryInfos []mall.Recovery_history, total int64) {
	db := global.GVA_DB.Model(&mall.Recovery_history{})
	db.Where("user_id = ?", info.UserId)
	db.Count(&total)
	limit := info.PageSize
	offset := info.PageSize * (info.PageNumber - 1)
	db.Limit(limit).Offset(offset).Order("re_time desc").Find(&recoveryInfos)
	return recoveryInfos, total
}
