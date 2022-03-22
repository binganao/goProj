package main

import "dServer/settings"

/*
cw.htm -> :8080 -> unreceived danmu / command
                          ^-ch- formatter <- blive <- bilibili

main -> server -> router -> api -> handler
*/
func main() {
	settings.ReadFlags()
	Start()
}
