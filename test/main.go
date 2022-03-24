package main

import (
	"fmt"
	"time"
)

var ch chan struct{}

func main() {
	ch = make(chan struct{})

	go func() {
		select {
		case <-ch:
			fmt.Println("123")
		case <-time.After(time.Minute):
			fmt.Println("TIMEOUT1")
		}
	}()
	go func() {
		select {
		case <-ch:
			fmt.Println("456")
		case <-time.After(time.Minute):
			fmt.Println("TIMEOUT2")
		}
	}()
	<-time.After(time.Second * 2)
	close(ch)
	//ch <- struct{}{}
	a, ok := <-ch
	fmt.Println(a, ok)
	<-time.After(time.Hour)
}
