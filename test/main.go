package main

import (
	"fmt"
	"time"
)

func main() {
	a()
	<-time.After(time.Second * 10)
}

func a() {
	go func() {
		for {
			fmt.Println(12)
			<-time.After(time.Millisecond * 100)
		}
	}()
	<-time.After(time.Second)
}
