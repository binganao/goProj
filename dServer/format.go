package main

import (
	"fmt"
	"html"
	"math"
	"strconv"

	"github.com/Akegarasu/blivedm-go/client"
	"github.com/Akegarasu/blivedm-go/message"
)

type DanmuEvent struct {
	event   int
	price   int
	content string
}

const (
	EventDanmu int = iota
	EventSuperchat
	EventGift
	EventGuard
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

func parseLevel(n int64) string {
	if n >= 15 {
		return strconv.Itoa(int(n / 5))
	} else {
		return ""
	}
}

func StartBlive(room string, ch chan *DanmuEvent, f func(c *client.Client, ch chan *DanmuEvent)) {
	f(client.NewClient(room), ch)
}

func HTML(c *client.Client, ch chan *DanmuEvent) {
	// 弹幕事件
	c.OnDanmaku(func(danmuku *message.Danmaku) {
		ch <- &DanmuEvent{
			event:   EventDanmu,
			content: fmt.Sprintf(`<span style="font-size: .64em">%s</span>%s`, html.EscapeString(danmuku.Sender.Uname), bigbold("<!---->"+html.EscapeString(danmuku.Content))),
		}
	})
	// 醒目留言事件
	c.OnSuperChat(func(superChat *message.SuperChat) {
		identity := string(" ᴀʙᴄ"[superChat.UserInfo.GuardLevel]) + parseLevel(int64(superChat.UserInfo.UserLevel))
		ch <- &DanmuEvent{
			event:   EventSuperchat,
			price:   superChat.Price,
			content: fmt.Sprintf(`%s%s:%s`, supbold(identity), html.EscapeString(superChat.UserInfo.Uname), bigbold(html.EscapeString(superChat.Message))),
		}
	})
	// 礼物事件
	c.OnGift(func(gift *message.Gift) {
		v := gift.TotalCoin / 1e3
		if gift.CoinType != "silver" && v >= 5 {
			identity := " ᴀʙᴄ"[gift.GuardLevel]
			ch <- &DanmuEvent{
				event:   EventGift,
				price:   v,
				content: bigbold(fmt.Sprintf(`%v%s 赠送%sx%d#%d`, identity, html.EscapeString(gift.Uname), html.EscapeString(gift.GiftName), gift.Num, v), .64+math.Max(math.Pow(float64(v), 1/3)/40, 1/(1+math.Pow(math.E, -.002*float64(v)+3)))),
			}
		}
	})
	// 上舰事件
	c.OnGuardBuy(func(guardBuy *message.GuardBuy) {
		v := guardBuy.Price / 1e3
		ch <- &DanmuEvent{
			event:   EventGuard,
			price:   v,
			content: fmt.Sprintf(`%s 成为 %s#%d`, bigbold(html.EscapeString(guardBuy.Username)), bigbold(html.EscapeString(guardBuy.GiftName)), v),
		}
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
