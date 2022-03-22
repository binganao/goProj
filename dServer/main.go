package main

/*
cw.htm -> :8080 -> unreceived danmu / command
                          ^-ch- formatter <- blive <- bilibili
*/
import (
	"dServer/handler"
	"dServer/settings"
)

func main() {
	settings.ReadFlags()
	handler.Start()
}
