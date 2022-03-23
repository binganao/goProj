package main

import (
	"dServer/settings"
	"fmt"
	"time"

	//grmon "github.com/bcicen/grmon/agent"
	"github.com/tidwall/gjson"
)

func Start() {
	//grmon.Start()
	ServerStatus.room = settings.Room
	Rooms = make(map[string]*Roomstatus)
	StartServer()
	var t Watcher

	for {
		if _, ok := Rooms[ServerStatus.room]; !ok {
			Rooms[ServerStatus.room] = &Roomstatus{Superchat: []ScStruct{}}
		}
		ExiprePurse()

		c := StartBlive(ServerStatus.room, HTML)
		t = StartPop(ServerStatus.room, t)

		if GetControl() {
			// fork from original blivedm repo
			// changes to Danmuku struct, Stop, Log.Fatal
			c.Stop()
			continue
		}
		<-time.After(time.Millisecond * 50)
	}
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
			return false
		case CMD_UPGRADE:
			//upgrade
			return false
		}
	}
}

func StartServer() {
	fmt.Println(ServerStatus, StatusList[ServerStatus.status])

	r := InitRouters()
	go r.Run(":" + settings.Port)
}

func StartPop(room string, t Watcher) Watcher {
	t.Stop()
	return StartUrlWatcher(time.Second*30, "https://api.live.bilibili.com/xlive/web-room/v1/index/getH5InfoByRoom?room_id="+room, "", "GET", func(s string) {
		js := gjson.Get(s, "data.room_info.online")
		ServerStatus.pop = int(js.Int())
	})
}

func ExiprePurse() {
	for _, v := range Rooms {
		if time.Now().Sub(v.PurseExpire) > 2*24*time.Hour {
			v.Purse = 0
		}
		var sc []ScStruct
		for _, j := range v.Superchat {
			if j.Expire.Sub(time.Now()) > 0 {
				sc = append(sc, j)
			}
		}
		v.Superchat = sc
	}
}
