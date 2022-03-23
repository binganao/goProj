package main

import (
	"dServer/settings"
	"fmt"
	"time"

	"github.com/tidwall/gjson"
)

func Start() {
	ServerStatus.room = settings.Room
	Rooms = make(map[string]*Roomstatus)
	StartServer()

	if _, ok := Rooms[ServerStatus.room]; !ok {
		Rooms[ServerStatus.room] = &Roomstatus{}
	}
	ExiprePurse()

	StartBlive(ServerStatus.room, HTML)
	StartPop(ServerStatus.room)

	if GetControl() {
	}
	<-time.After(time.Millisecond * 200)
}

func GetControl() bool {
	for {
		state := <-control
		switch state.cmd {
		case CMD_CHANGE_ROOM:
			ServerStatus.room = state.room
			return true
		case CMD_RESTART:
			//restart
			<-time.After(time.Millisecond * 100)
			return true
		case CMD_UPGRADE:
			//upgrade
			return true
		}
	}
}

func StartServer() {
	fmt.Println(ServerStatus, StatusList[ServerStatus.status])

	r := InitRouters()
	go r.Run(":" + settings.Port)
}

func StartPop(room string) {
	RunUrlWatcher("RoomPop", time.Second*30, "https://api.live.bilibili.com/xlive/web-room/v1/index/getH5InfoByRoom?room_id="+room, "", "GET", func(s string) {
		js := gjson.Get(s, "data.room_info.online")
		ServerStatus.pop = int(js.Int())
	})
}

func ExiprePurse() {
	for _, v := range Rooms {
		if time.Now().Sub(v.purseExpire) > 2*24*3600*time.Second {
			v.purse = 0
		}
		var sc []ScStruct
		for _, j := range v.superchat {
			if j.expire.Sub(time.Now()) > 0 {
				sc = append(sc, j)
			}
		}
		v.superchat = sc
	}
}
