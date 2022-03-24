package main

import (
	"dServer/settings"
	"fmt"
	"os"
	"os/exec"
	"time"

	grmon "github.com/bcicen/grmon/agent"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/wmillers/blivedm-go/client"
)

func Start() {
	if settings.Debug {
		grmon.Start()
	}

	StartServer()

	var t Watcher
	for {
		Rooms.RWMutex.Lock()
		if _, ok := Rooms.Value[ServerStatus.Room]; !ok {
			Rooms.Value[ServerStatus.Room] = &Roomstatus{Superchat: []Superchat{}}
		}
		Rooms.RWMutex.Unlock()
		ExiprePurse()

		t = StartPop(ServerStatus.Room, t)

		if GetControl(StartBlive(ServerStatus.Room, HTML)) {
			// fork from original blivedm repo
			// changes to Danmuku struct, Stop, Log.Fatal
			continue
		} else {
			break
		}
	}

	fmt.Println("[QUIT]")
}

func GetControl(c *client.Client) bool {
	for {
		state := <-control
		switch state.cmd {
		case CMD_CHANGE_ROOM:
			ServerStatus.Room = state.room
			c.Stop()
			return true
		case CMD_UPGRADE:
			//upgrade
			//return false
			fallthrough
		case CMD_RESTART:
			RestartServer(c)
			return false
		}
	}
}

func RestartServer(c *client.Client) {
	c.Stop()
	self, err := os.Executable()
	if err != nil {
		fmt.Println("FAILED restart: ", err)
	}

	cmd := exec.Command(self, Args())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Env = os.Environ()
	cmd.Run()
}

func StartServer() {
	if settings.Debug {
		fmt.Println(ServerStatus, StatusList[ServerStatus.Status])
	} else {
		fmt.Println("#"+Args()+"\nRun"+ServerStatus.Room+" :"+settings.Port+settings.Path, time.Now().Format("2006-01-02 15:04:05.0-07"))
		gin.SetMode(gin.ReleaseMode)
	}

	r := InitRouters()
	go r.Run(":" + settings.Port)
}

func StartPop(room string, t Watcher) Watcher {
	t.Stop()
	return StartWatcher(time.Minute, func() {
		s := CorsAccess("https://api.live.bilibili.com/xlive/web-room/v1/index/getH5InfoByRoom?room_id="+room, "", "GET")
		js := gjson.Get(s, "data.room_info.live_status")
		if js.Int() > 0 {
			js = gjson.Get(s, "data.room_info.online")
			updatePop(int(js.Int()))
		} else {
			updatePop(1)
		}
	})
}

func ExiprePurse() {
	Rooms.RWMutex.Lock()
	for _, v := range Rooms.Value {
		if v.PurseExpire.Sub(time.Now()) < 0 {
			v.Purse = 0
		}
		var sc []Superchat
		for _, j := range v.Superchat {
			if j.Expire.Sub(time.Now()) > 0 {
				sc = append(sc, j)
			}
		}
		v.Superchat = sc
	}
	Rooms.RWMutex.Unlock()
}

func Args() string {
	return fmt.Sprintf(`-path "%v" -port %v -debug=%v -room %v -store "%v" -timeout %v`,
		settings.Path, settings.Port, settings.Debug, ServerStatus.Room, ServerStatus.Store, settings.Timeout)
}
