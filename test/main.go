package main

import (
	"fmt"
	"time"
)

type Watcher struct {
	IsRunning bool
	Done      chan bool
}

func StartUrlWatcher(t time.Duration, w ...Watcher) (c Watcher) {
	if len(w) != 0 {
		c = w[0]
	} else {
		c = Watcher{
			IsRunning: true,
			Done:      make(chan bool),
		}
	}
	go c.run(t)
	return c
}

func (c *Watcher) run(t time.Duration) {
	for {
		fmt.Println(c.IsRunning)
		select {
		case <-c.Done:
			//	c.IsRunning = false
			break
		case <-time.After(t):
			fmt.Println(123)
		}
	}
	fmt.Println(c.IsRunning)
}

func (c *Watcher) Stop() {
	fmt.Println(c.IsRunning)
	if c.IsRunning {
		c.IsRunning = false
		c.Done <- true
	}
}

func main() {
	c := StartUrlWatcher(time.Second)
	<-time.After(time.Second * 2)
	c.Stop()
	<-time.After(time.Hour)
}
