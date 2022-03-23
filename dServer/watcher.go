package main

import (
	"time"
)

// run: register (check) <- run -> go
type WatcherStruct struct {
	IsRunning bool
	stop      chan bool
}

var UrlWatcher map[string]WatcherStruct

func init() {
	UrlWatcher = make(map[string]WatcherStruct)
}

func goUrlWatcher(name string, t time.Duration, url string, data string, method string, f func(string)) {
	for {
		select {
		case <-UrlWatcher[name].stop:
			break
		case <-time.After(t):
			go cover(func() { f(CorsAccess(url, data, method)) })
		}
	}
}

func CheckUrlWatcher(name string) bool {
	_, ok := UrlWatcher[name]
	return ok
}

func RegisterUrlWatcher(name string) bool {
	if CheckUrlWatcher(name) {
		return false
	} else {
		UrlWatcher[name] = WatcherStruct{}
		return true
	}
}

func RunUrlWatcher(name string, t time.Duration, url string, data string, method string, f func(string)) {
	if CheckUrlWatcher(name) {
		StopUrlWatcher(name)
	} else {
		RegisterUrlWatcher(name)
	}
	go goUrlWatcher(name, t, url, data, method, f)
}

func StopUrlWatcher(name string) bool {
	if CheckUrlWatcher(name) && UrlWatcher[name].IsRunning {
		UrlWatcher[name].stop <- true
		return true
	} else {
		return false
	}
}

func UnregisterWatcher(name string) bool {
	if CheckUrlWatcher(name) {
		delete(UrlWatcher, name)
		return true
	} else {
		return false
	}
}
