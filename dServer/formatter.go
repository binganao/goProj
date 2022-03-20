package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Akegarasu/blivedm-go/client"
	"github.com/Akegarasu/blivedm-go/message"
)

func supbold(s string) string {
	return fmt.Sprintf(`<span style="font-weight: bold; vertical-align: super; font-size: .8em">%s</span>`, s)
}

func bigbold(s string, sizes ...float64) string {
	var size float64
	if len(sizes) == 0 {
		size = 1.2
	} else {
		size = sizes[0]
	}
	return fmt.Sprintf(`<span style="font-weight: bold; font-size: %.2fem">%s</span>`, size, s)
}

func escapeHTML(s string) string {
	return strings.ReplaceAll(s, "<", "&lt;")
}

func parseLevel(n int64) string {
	if n >= 15 {
		return strconv.Itoa(int(n / 5))
	} else {
		return ""
	}
}

func convertCoin(p int, price chan int) int {
	v := p / 1e3
	if v > 0 {
		price <- v
	}
	return v
}

type schat struct {
	price   int
	content string
}

func HtmlFormatter(room string, ch chan string, price chan int, sc chan schat) {
	c := client.NewClient(room)
	// 弹幕事件
	c.OnDanmaku(func(danmuku *message.Danmaku) {
		ch <- fmt.Sprintf(`<span style="font-size: .64em">%s</span>%s`, escapeHTML(danmuku.Sender.Uname), bigbold("<!---->"+escapeHTML(danmuku.Content)))
	})
	// 醒目留言事件
	c.OnSuperChat(func(superChat *message.SuperChat) {
		v := convertCoin(superChat.Price*1e3, price)
		identity := string(" ᴀʙᴄ"[superChat.UserInfo.GuardLevel]) + parseLevel(int64(superChat.UserInfo.UserLevel))
		sc <- schat{price: v, content: fmt.Sprintf(`%s%s:%s`, supbold(identity), escapeHTML(superChat.UserInfo.Uname), bigbold(escapeHTML(superChat.Message)))}
	})
	// 礼物事件
	c.OnGift(func(gift *message.Gift) {
		v := convertCoin(gift.TotalCoin, price)
		if gift.CoinType != "silver" && v >= 5 {
			identity := " ᴀʙᴄ"[gift.GuardLevel]
			ch <- bigbold(fmt.Sprintf(`%v%s 赠送%sx%d#%d`, identity, escapeHTML(gift.Uname), escapeHTML(gift.GiftName), gift.Num, v), .64+math.Max(math.Pow(float64(v), 1/3)/40, 1/(1+math.Pow(math.E, -.002*float64(v)+3))))
		}
	})
	// 上舰事件
	c.OnGuardBuy(func(guardBuy *message.GuardBuy) {
		v := convertCoin(guardBuy.Price, price)
		ch <- fmt.Sprintf(`%s 成为 %s#%d`, bigbold(escapeHTML(guardBuy.Username)), bigbold(escapeHTML(guardBuy.GiftName)), v)
	})
	// 【可选】设置弹幕服务器，不设置就会从 api 获取服务器地址
	// 该函数设置服务器为 wss://broadcastlv.chat.bilibili.com/sub
	c.UseDefaultHost()
	// 启动
	err := c.ConnectAndStart()
	if err != nil {
		fmt.Println(err)
	}
	/*
		cl := blivedm.BLiveWsClient{ShortId: room, HearbeatInterval: 25 * time.Second}
		fmt.Println(cl.GetRoomInfo(), cl.GetDanmuInfo())
		cl.ConnectDanmuServer()

		cl.RegHandler(blivedm.CmdDanmaku, func(context *blivedm.Context) {
			msg, _ := context.ToDanmakuMessage()
			var identity string
			if msg.Admin {
				identity += "⚑"
			}
			identity += string(" ᴀʙᴄ"[msg.PrivilegeType]) + parseLevel(msg.UserLevel)
			ch <- fmt.Sprintf(`<span style="font-size: .64em">%s%s</span>%s`, supbold(identity), escapeHTML(msg.Uname), bigbold("<!---->"+escapeHTML(msg.Msg)))
		})
	*/
	fmt.Println("started")
}
