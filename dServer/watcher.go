package main

import (
	"fmt"
	"time"
)

type Watcher struct {
	IsRunning bool
	Done      chan bool
}

func GetWatcher() Watcher {
	return Watcher{
		IsRunning: true,
		Done:      make(chan bool),
	}
}

func (c *Watcher) Stop() {
	if c.IsRunning {
		c.IsRunning = false
		close(c.Done)
	}
}

func StartWatcher(t time.Duration, f func()) Watcher {
	c := GetWatcher()
	go runUrl(c, t, f)
	return c
}

func runUrl(c Watcher, t time.Duration, f func()) {
	for {
		select {
		case <-c.Done:
			return
		case <-time.After(t):
			go cover(f)
		}
	}
}

func cover(f func()) {
	defer func() {
		if pan := recover(); pan != nil {
			fmt.Printf("event error: %v\n", pan)
		}
	}()
	f()
}
