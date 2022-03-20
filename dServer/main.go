package main

import "fmt"

/*
cw.htm -> :8080 -> unreceived danmu / command
                          ^-ch- formatter <- blive <- bilibili
*/

type danmu struct {
	contents []string
	i        int
}

var purse int

func main() {
	ch := make(chan string)
	price := make(chan int)
	sc := make(chan schat)
	var d danmu
	HtmlFormatter("545068", ch, price, sc)
	for {
		select {
		case s := <-ch:
			d.contents = append(d.contents, s)
			if len(d.contents) > 2e4 {
				d.contents = d.contents[2000:]
				d.i -= 2000
				if d.i < 0 {
					d.i = 0
				}
			}
			fmt.Println(s)
		case p := <-price:
			purse += p
			fmt.Println(purse)
		}
	}
}
