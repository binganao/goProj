package main

import (
	"dServer/settings"
	"fmt"
	"os"
	"syscall"
	"time"

	//grmon "github.com/bcicen/grmon/agent"
	"github.com/gin-gonic/gin"
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
		case CMD_UPGRADE:
			//upgrade
			//return false
			fallthrough
		case CMD_RESTART:
			//restart
			<-time.After(time.Millisecond * 100)
			self, err := os.Executable()
			if err != nil {
				fmt.Println("FAILED restart: ", err)
			}
			syscall.Exec(self, os.Args, os.Environ())
			return false
		}
	}
}

func StartServer() {
	if settings.Debug {
		fmt.Println(ServerStatus, StatusList[ServerStatus.status])
	} else {
		fmt.Println("Run :" + settings.Port)
		gin.SetMode(gin.ReleaseMode)
	}

	r := InitRouters()
	go r.Run(":" + settings.Port)
}

func StartPop(room string, t Watcher) Watcher {
	t.Stop()
	return StartWatcher(time.Minute, func() {
		s := CorsAccess("https://api.live.bilibili.com/xlive/web-room/v1/index/getH5InfoByRoom?room_id="+room, "", "GET")
		js := gjson.Get(s, "data.room_info.online")
		updatePop(int(js.Int()))
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
