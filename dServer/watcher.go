package main

import (
	"fmt"
	"time"
)

type Watcher struct {
	IsRunning bool
	Done      chan bool
}

func StartUrlWatcher(t time.Duration, url string, data string, method string, f func(string)) Watcher {
	c := Watcher{
		IsRunning: true,
		Done:      make(chan bool),
	}
	go c.run(t, url, data, method, f)
	return c
}

func (c *Watcher) run(t time.Duration, url string, data string, method string, f func(string)) {
	for {
		select {
		case <-c.Done:
			return
		case <-time.After(t):
			go cover(func() { f(CorsAccess(url, data, method)) })
		}
	}
}

func (c *Watcher) Stop() {
	if c.IsRunning {
		c.IsRunning = false
		c.Done <- true
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
