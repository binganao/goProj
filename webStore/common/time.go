package common

import "time"

func NowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func BeforeDay(days int) (dayTime string) {
	return time.Now().Add(time.Duration(days) * -24 * time.Hour).Format("2006-01-02")
}
