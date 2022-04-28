package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, World!")
	<-time.After(time.Second * 100)
}
