package main

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	HTTP_OK      = 200
	HTTP_INVALID = 400
)

func SetHeaderHTML(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
}

func Ok(c *gin.Context) {
	c.JSON(HTTP_OK, gin.H{
		"test": "ok",
	})
}

func ApplyMatch(c *gin.Context, statement map[string]func(c *gin.Context), cmd string) {
	isValid := false
	for k, v := range statement {
		if ok, _ := regexp.MatchString("(?i)"+k, cmd); ok {
			isValid = true
			v(c)
			break
		}
	}
	if !isValid {
		c.JSON(HTTP_INVALID, gin.H{"invalid": cmd})
	}
}

func CheckQuery(c *gin.Context) string {
	RecordClient(c)
	cmds := c.Request.URL.Query()
	var cmd string
	if len(cmds) != 0 {
		for k := range cmds {
			cmd = k
			break
		}
	}
	return cmd
}

func HTMLString(c *gin.Context, s string) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(HTTP_OK, s)
}

func ParseGet(c *gin.Context) {
	cmd := CheckQuery(c)
	if len(cmd) == 0 {
		GetDanmu(c)
		return
	}
	SetHeaderHTML(c)
	statement := map[string]func(c *gin.Context){
		`^\d+$`: func(c *gin.Context) {
			HTMLString(c, ChangeRoom(cmd))
		},
		`^history$`: GetHistory,
		`^restart$`: func(c *gin.Context) {
			ServerStatus.pop = 1
			HTMLString(c, "[RESTART] RECV OK")
			control <- ControlStruct{cmd: CMD_RESTART}
		},
		`^upgrade$`: func(c *gin.Context) {
			ServerStatus.pop = 1
			HTMLString(c, "[UPGRADE] Depends on network")
			control <- ControlStruct{cmd: CMD_UPGRADE}
		},
		`^status$`: GetStatus,
		`^clients$`: func(c *gin.Context) {
			// LOWER == private == emit
			c.JSON(HTTP_OK, ServerStatus.clients)
		},
		`^kick$`: func(c *gin.Context) {
			SetKick(c)
			GetDanmu(c)
		},
		`^call:`: func(c *gin.Context) {
			addDanmu(cmd[strings.Index(cmd, ":")+1:])
			HTMLString(c, "[CALLING]")
		},
		`^js:`: func(c *gin.Context) {
			if ServerStatus.pop == 1 {
				ServerStatus.pop = 9999
			}
			js := cmd[strings.Index(cmd, ":")+1:]
			addDanmu("[JS] " + js)
			HTMLString(c, "[JS-EXCUTING] "+js)
		},
		`^cors:`: func(c *gin.Context) {
			s, _ := c.GetRawData()
			HTMLString(c, CorsAccess(cmd[strings.Index(cmd, ":")+1:], string(s), "GET", c.Request.Header))
		},
		`^time$`: func(c *gin.Context) {
			HTMLString(c, strconv.FormatInt(time.Now().UnixMilli(), 10))
		},
		//`^s4f_:`: ,
		`store`: func(c *gin.Context) {
			c.JSON(HTTP_OK, ServerStatus.store)
		},
		//any
	}
	ApplyMatch(c, statement, cmd)
}

func ParsePost(c *gin.Context) {
	cmd := CheckQuery(c)
	if len(cmd) == 0 {
		c.JSON(HTTP_INVALID, gin.H{"reason": "Empty Query"})
		return
	}
	statement := map[string]func(c *gin.Context){
		`^store$`: func(c *gin.Context) {
			s, _ := c.GetRawData()
			ServerStatus.store = string(s)
		},
		`^cors:`: func(c *gin.Context) {
			s, _ := c.GetRawData()
			HTMLString(c, CorsAccess(cmd[strings.Index(cmd, ":")+1:], string(s), "POST", c.Request.Header))
		},
	}
	ApplyMatch(c, statement, cmd)
}

func ParsePut(c *gin.Context) {
	cmd := CheckQuery(c)
	if len(cmd) == 0 {
		c.JSON(HTTP_INVALID, gin.H{"reason": "Empty Query"})
		return
	}
	statement := map[string]func(c *gin.Context){
		`^cors:`: func(c *gin.Context) {
			s, _ := c.GetRawData()
			HTMLString(c, CorsAccess(cmd[strings.Index(cmd, ":")+1:], string(s), "PUT", c.Request.Header))
		},
	}
	ApplyMatch(c, statement, cmd)
}
