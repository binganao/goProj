package handler

import (
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

func Start() {
	serverStatus.room = "545068"
	rooms = make(map[string]*RoomStatus)
	if _, ok := rooms[serverStatus.room]; !ok {
		rooms[serverStatus.room] = &RoomStatus{}
	}
	go StartBlive(serverStatus.room, HTML)

	controlChan := make(chan string)
	go StartServer(controlChan)
}

func StartServer(ch chan string) {
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
