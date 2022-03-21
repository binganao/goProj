package main

/*
cw.htm -> :8080 -> unreceived danmu / command
                          ^-ch- formatter <- blive <- bilibili
*/
import (
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

func handleDanmuEvent(c *DanmuEvent) {
	switch (*c).event {
	case EventDanmu:
		addDanmu(c.content)
	case EventGuard:
		fallthrough
	case EventGift:
		rooms[serverStatus.room].purse += c.price
		addDanmu(c.content)
		//case EventSuperchat:
	}
}

func main() {
	statusContent := []string{
		"",
		"[SLEEP] no room (CAREFUL with s4f_: cmd)",
		"[SLEEP] & [STUCK] at que.qsize() > 5000",
		"[SLEEP] & [RESTART] pong<-",
		"[UPGRADE] it depends on network",
	}
	fmt.Println(serverStatus, statusContent[serverStatus.status])

	serverStatus.room = "545068"
	rooms[serverStatus.room] = &RoomStatus{}
	ch := make(chan *DanmuEvent)
	StartBlive(serverStatus.room, ch, HTML)
	for {
		select {
		case c := <-ch:
			handleDanmuEvent(c)
		}
	}
}
