package main

import (
	"fmt"
	//"io/uiutil"
	//"net/http"
	//"os"
	//"strconv"
	"time"
)

func main() {
	defer func(start time.Time) {
		fmt.Println("Total use:\t", time.Since(start))
	}(time.Now())
	a := func() {
		fmt.Println(1)
	}
	go a()
	fmt.Println(2)

	b := map[string]int{}
	b["asd"] = 1
	o, k := b["s"]
	fmt.Println(b["asd"], b["aas"], o, k, len(b))
}
