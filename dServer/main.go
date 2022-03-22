package main

import "dServer/settings"

/*
cw.htm -> :8080 -> unreceived danmu / command
                          ^-ch- formatter <- blive <- bilibili
*/
func main() {
	settings.ReadFlags()
	Start()
}
