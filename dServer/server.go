package main

import (
	"dServer/settings"
	"fmt"
	"time"
)

func Start() {
	ServerStatus.room = "545068"
	Rooms = make(map[string]*Roomstatus)
	go StartServer()
	ch := make(chan bool)

	for {
		if _, ok := Rooms[ServerStatus.room]; !ok {
			Rooms[ServerStatus.room] = &Roomstatus{}
		}
		go StartBlive(ServerStatus.room, HTML, ch)

		if !GetControl() {
			break
		}
		ch <- true
	}
}

func GetControl() bool {
	for {
		state := <-control
		switch state.cmd {
		case CMD_CHANGE_ROOM:
			//some changes to purse (only remain 3days)
			ServerStatus.room = state.room
			return true
		case CMD_RESTART:
			//restart
			<-time.After(time.Second)
			return true
		case CMD_UPGRADE:
			//upgrade
		}
	}
}

func StartServer() {
	fmt.Println(ServerStatus, StatusList[ServerStatus.status])

	r := InitRouters()
	r.Run(":" + settings.Port)
}
