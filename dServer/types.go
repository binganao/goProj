package main

import "time"

type Roomstatus struct {
	purse       int
	purseExpire time.Time
	superchat   []struct {
		expire  time.Time
		price   int
		content string
	}
}

var History []string
var Rooms map[string]*Roomstatus

type ClientsStruct struct {
	First    string
	Interval int
	Last     string
	Path     []string
	Reads    int
	Kick     string
	//platform string
	//browser  string
}

var ServerStatus struct {
	i          int
	room       string
	other_room string
	pop        int
	status     int
	store      string
	clients    map[string]*ClientsStruct
}
var StatusList []string

type ControlStruct struct {
	cmd  int
	room string
}

var control chan ControlStruct

const (
	CMD_CHANGE_ROOM = iota
	CMD_RESTART
	CMD_UPGRADE
)

func init() {
	StatusList = []string{
		"",
		"[SLEEP] no room (CAREFUL with s4f_: cmd)",
		"[SLEEP] & [STUCK] at que.qsize() > 5000",
		"[SLEEP] & [RESTART] pong<-",
		"[UPGRADE] it depends on network",
	}
	ServerStatus.clients = make(map[string]*ClientsStruct)
}
