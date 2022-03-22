package main

import (
	"dServer/settings"
	"fmt"
	"time"
)

type Roomstatus struct {
	purse     int
	superchat []struct {
		expire  time.Time
		price   int
		content string
	}
}

var History []string
var Rooms map[string]*Roomstatus
var ServerStatus struct {
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
var StatusList []string

func Start() {
	ServerStatus.room = "545068"
	Rooms = make(map[string]*Roomstatus)
	if _, ok := Rooms[ServerStatus.room]; !ok {
		Rooms[ServerStatus.room] = &Roomstatus{}
	}
	go StartBlive(ServerStatus.room, HTML)

	controlChan := make(chan string)
	go StartServer(controlChan)
	GetControl(controlChan)
}

func GetControl(ch chan string) {
	<-ch
}

func StartServer(ch chan string) {
	StatusList = []string{
		"",
		"[SLEEP] no room (CAREFUL with s4f_: cmd)",
		"[SLEEP] & [STUCK] at que.qsize() > 5000",
		"[SLEEP] & [RESTART] pong<-",
		"[UPGRADE] it depends on network",
	}
	fmt.Println(ServerStatus, StatusList[ServerStatus.status])

	r := InitRouters()
	r.Run(":" + settings.Port)
}
