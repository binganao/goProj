package main

import (
	"fmt"
	"html"
	"math"

	"github.com/Akegarasu/blivedm-go/client"
	"github.com/Akegarasu/blivedm-go/message"
)

type DanmuEvent struct {
	Event   int
	Price   int
	Content string
}

func StartBlive(room string, f func(c *client.Client)) {
	f(client.NewClient(room))
}

func HTML(c *client.Client) {
	// 弹幕事件
	c.OnDanmaku(func(danmuku *message.Danmaku) {
		addDanmu(fmt.Sprintf(`<span style="font-size: .64em">%s</span>%s`, html.EscapeString(danmuku.Sender.Uname), bigbold("<!---->"+html.EscapeString(danmuku.Content))))
	})
	// 醒目留言事件 UNdONE
	c.OnSuperChat(func(superChat *message.SuperChat) {
		identity := string(" ᴀʙᴄ"[superChat.UserInfo.GuardLevel]) + parseLevel(int64(superChat.UserInfo.UserLevel))
		addPurse(superChat.Price)
		addDanmu(fmt.Sprintf(`%s%s:%s`, supbold(identity), html.EscapeString(superChat.UserInfo.Uname), bigbold(html.EscapeString(superChat.Message))))
	})
	// 礼物事件
	c.OnGift(func(gift *message.Gift) {
		v := gift.TotalCoin / 1e3
		if gift.CoinType != "silver" && v >= 5 {
			identity := " ᴀʙᴄ"[gift.GuardLevel]
			addPurse(v)
			addDanmu(bigbold(fmt.Sprintf(`%v%s 赠送%sx%d#%d`, identity, html.EscapeString(gift.Uname), html.EscapeString(gift.GiftName), gift.Num, v), .64+math.Max(math.Pow(float64(v), 1/3)/40, 1/(1+math.Pow(math.E, -.002*float64(v)+3)))))
		}
	})
	// 上舰事件
	c.OnGuardBuy(func(guardBuy *message.GuardBuy) {
		v := guardBuy.Price / 1e3
		addPurse(v)
		addDanmu(fmt.Sprintf(`%s 成为 %s#%d`, bigbold(html.EscapeString(guardBuy.Username)), bigbold(html.EscapeString(guardBuy.GiftName)), v))
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
		cl.RegHandler(blivedm.CmdDanmaku, func(context *blivedm.Context) {
			msg, _ := context.ToDanmakuMessage()
			var identity string
			if msg.Admin {
				identity += "⚑"
			}
			identity += string(" ᴀʙᴄ"[msg.PrivilegeType]) + parseLevel(msg.UserLevel)
			ch <- fmt.Sprintf(`<span style="font-size: .64em">%s%s</span>%s`, supbold(identity), html.EscapeString(msg.Uname), bigbold("<!---->"+html.EscapeString(msg.Msg)))
		})
	*/
	fmt.Println("started")
}

func addDanmu(s string) {
	History = append(History, s)
	if len(History) > 2e4 {
		History = History[2000:]
		ServerStatus.i -= 2000
		if ServerStatus.i < 0 {
			ServerStatus.i = 0
		}
	}
}

func addPurse(price int) {
	Rooms[ServerStatus.room].purse += price
}
