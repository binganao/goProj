package main

import (
	"log"
	"main/global"
	"main/model"
	"strconv"
	"strings"
)

func main() {
	test("This is a model test for redis mysql elastic")
	param := model.AppCartAddParam{
		ProductId:    9,
		ProductCount: 2,
		UserId:       "abcd",
	}
	log.Println(redisAdd(param))
	log.Println(global.Rdb.HGet(global.Ctx, redisKey(param.UserId), strconv.Itoa(int(param.ProductId))).Val())
	log.Println(redisShow("abcd"))
	var res float64
	global.Db.Raw("select SUM(total_price) from `order` where created like ?", "2021-11-15%").Find(&res)
	log.Println(res)
}

func test[T any](a T) {
	log.Println(a)
}

func redisAdd(param model.AppCartAddParam) bool {
	key := redisKey(param.UserId)
	pid := strconv.Itoa(int(param.ProductId))
	return global.Rdb.HSetNX(global.Ctx, key, pid, param.ProductCount).Val()
	//global.Rdb.HDel(global.Ctx, key, param.ProductId)
}

func redisShow(userId string) (result model.AppCartInfo) {
	key := redisKey(userId)
	strProducts := global.Rdb.HGetAll(global.Ctx, key).Val()
	productIds := make([]int, 0)
	products := make(map[int]int)
	for id, count := range strProducts {
		intId, _ := strconv.Atoi(id)
		intCount, _ := strconv.Atoi(count)
		productIds = append(productIds, intId)
		products[intId] = intCount
	}
	if len(products) > 0 {
		global.Db.Table("product").Find(&result.CartItem, productIds)
		for _, v := range result.CartItem {
			result.TotalPrice += v.Price * float64(products[int(v.Id)])
		}
	}
	return
}

func redisKey(id string) string {
	return strings.Join([]string{"user", id, "cart"}, ":")
}
