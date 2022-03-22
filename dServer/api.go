package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	STATUS_OK = 200
)

func SetHeaderHTML(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
}

func Ok(c *gin.Context) {
	c.JSON(STATUS_OK, gin.H{
		"test": "ok",
	})
}

func GetApi(c *gin.Context) {
	cmd := c.Request.URL.Query()
	if len(cmd) == 0 {
		GetDanmu(c, ServerStatus.i)
		ServerStatus.i = len(History)
		return
	}

	var action string
	for k := range cmd {
		action = k
		break
	}
	statement := map[string]func(c *gin.Context){
		//`^/d+$`: ,
		`^history$`: GetHistory,
		//`^restart$`: ,
		//`^upgrade$`: ,
		`^status$`: GetStatus,
		`^clients$`: func(c *gin.Context) {
			c.JSON(STATUS_OK, ServerStatus.clients)
		},
		//`^kick$`: ,
		`^call:`: func(c *gin.Context) {
			addDanmu(action[strings.Index(action, ":")+1:])
			c.String(STATUS_OK, "[CALLING]")
		},
		`^js:`: func(c *gin.Context) {
			if ServerStatus.pop == 1 {
				ServerStatus.pop = 9999
			}
			js := action[strings.Index(action, ":")+1:]
			addDanmu("[JS] " + js)
			c.String(STATUS_OK, "[JS-EXCUTING] "+js)
		},
		//`^cors:`: ,
		`^time$`: func(c *gin.Context) {
			c.String(STATUS_OK, strconv.FormatInt(time.Now().UnixMilli(), 10))
		},
		//`^s4f_:`: ,
		//`store`: ,
	}

	isValid := false
	for k, v := range statement {
		if ok, _ := regexp.MatchString("(?i)"+k, action); ok {
			isValid = true
			v(c)
			break
		}
	}
	if !isValid {
		c.String(STATUS_OK, "[err] Invalid: "+action)
	}
}

func PostApi(c *gin.Context) {}

func GetStatus(c *gin.Context) {
	c.JSON(STATUS_OK, gin.H{
		"room":           ServerStatus.room,
		"pop":            ServerStatus.pop,
		"unread":         len(History) - ServerStatus.i,
		"status":         ServerStatus.status,
		"status_content": StatusList[ServerStatus.status],
	})
}

func Reverse(original []string) []string {
	out := make([]string, len(original))
	for i, k := range original {
		out[len(original)-i-1] = k
	}
	return out
}

func GetDanmu(c *gin.Context, i int) {
	SetHeaderHTML(c)
	res := ""
	if i < len(History) {
		res = strings.Join(History[i:], "<br>")
	}
	c.String(STATUS_OK, res)
}

func GetHistory(c *gin.Context) {
	SetHeaderHTML(c)
	limit := 2000
	former := len(History) - limit
	var res string
	if former > 0 {
		res = fmt.Sprintf("%s<details><summary>%d shown, %d left</summary>%s</details>", strings.Join(Reverse(History[former:]), "<br>"), limit, former, strings.Join(Reverse(History[:former]), "<br>"))
	} else {
		res = strings.Join(Reverse(History), "<br>")
	}
	c.String(STATUS_OK, res)
}

func GetFavicon(c *gin.Context) {
	c.Status(204)
}
