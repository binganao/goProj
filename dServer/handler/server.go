package handler

import (
	"dServer/handler"
	"dServer/routers"
	"dServer/settings"
	"fmt"
	"time"
)

type RoomStatus struct {
	purse     int
	superchat []struct {
		expire  time.Time
		price   int
		content string
	}
}

var history []string
var rooms map[string]*RoomStatus
var serverStatus struct {
	unread  int
	i       int
	room    string
	pop     int
	status  int
	clients []map[string]struct {
		first    string
		interval int
		last     string
		path     []string
		reads    int
		kick     string
		platform string
		browser  string
	}
}

func addDanmu(s string) {
	history = append(history, s)
	if len(history) > 2e4 {
		history = history[2000:]
		serverStatus.i -= 2000
		if serverStatus.i < 0 {
			serverStatus.i = 0
		}
	}
	fmt.Println(s)
}

func handleDanmuEvent(ch chan *handler.DanmuEvent) {
	for {
		c := <-ch
		switch c.Event {
		case handler.EventDanmu:
			addDanmu(c.Content)
		case handler.EventGuard:
			fallthrough
		case handler.EventGift:
			rooms[serverStatus.room].purse += c.Price
			addDanmu(c.Content)
			//case handler.EventSuperchat:
		}
	}
}

func Start() {
	serverStatus.room = "545068"
	if _, ok := rooms[serverStatus.room]; !ok {
		rooms[serverStatus.room] = &RoomStatus{}
	}
	ch := make(chan *handler.DanmuEvent)
	go handler.StartBlive(serverStatus.room, ch, handler.HTML)
	go handleDanmuEvent(ch)

	controlChan := make(chan string)
	go startServer(controlChan)
}

func startServer(ch chan string) {
	statusContent := []string{
		"",
		"[SLEEP] no room (CAREFUL with s4f_: cmd)",
		"[SLEEP] & [STUCK] at que.qsize() > 5000",
		"[SLEEP] & [RESTART] pong<-",
		"[UPGRADE] it depends on network",
	}
	fmt.Println(serverStatus, statusContent[serverStatus.status])
	r := routers.InitRouters()
	r.Run(":" + settings.Port)
}
