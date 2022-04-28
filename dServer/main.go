package main

import (
	"dServer/service"
)

/*
cw.htm -> :8080 -> unreceived danmu / command
                          ^-ch- formatter <- blive <- bilibili

main -> server -> router -> api -> handler
*/
func main() {
	service.Start()
}
