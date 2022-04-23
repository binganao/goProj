package service

import (
	"strconv"
	"strings"

	"mall/global"
	"mall/models"
)

type CartService struct{}

func (c *CartService) Add(param models.WebCartAddParam) int64 {
	key := strings.Join([]string{"user", param.UserId, "cart"}, ":")
	pid := strconv.Itoa(int(param.ProductId))
	if v := global.Rdb.HSetNX(ctx, key, pid, param.ProductCount).Val(); v {
		return 1
	} else {
		return 0
	}
}

func (c *CartService) Delete(param models.WebCartDeleteParam) int64 {
	key := strings.Join([]string{"user", param.UserId, "cart"}, ":")
	return global.Rdb.HDel(ctx, key, param.ProductId).Val()
}

func (c *CartService) Clear(param models.WebCartClearParam) int64 {
	key := strings.Join([]string{"user", param.UserId, "cart"}, ":")
	pidsAndCounts := global.Rdb.HGetAll(ctx, key).Val()
	var rows int64
	for id := range pidsAndCounts {
		rows += global.Rdb.HDel(ctx, key, id).Val()
	}
	return rows
}

func (c *CartService) GetInfo(param models.WebCartQueryParam) models.WebCartInfo {
	var cartInfo models.WebCartInfo
	key := strings.Join([]string{"user", param.UserId, "cart"}, ":")
	productIdsAndCounts := global.Rdb.HGetAll(ctx, key).Val()
	productIds := make([]uint, 0)
	idsAndCounts := make(map[uint64]int, 0)
	for pid, pcount := range productIdsAndCounts {
		id, _ := strconv.Atoi(pid)
		count, _ := strconv.Atoi(pcount)
		productIds = append(productIds, uint(id))
		idsAndCounts[uint64(id)] = count
	}
	if len(productIds) > 0 {
		global.Db.Table("product").Find(&cartInfo.CartItem, productIds)
		for index, item := range cartInfo.CartItem {
			cartInfo.CartItem[index].Count = idsAndCounts[item.Id]
			cartInfo.TotalPrice = cartInfo.TotalPrice + item.Price*float64(idsAndCounts[item.Id])
		}
		return cartInfo
	}
	return cartInfo
}
