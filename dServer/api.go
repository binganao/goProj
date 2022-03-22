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

func RecordClient(c *gin.Context) {
	ua := c.Request.UserAgent()
	path := c.Request.RequestURI
	if v, ok := ServerStatus.clients[ua]; ok {
		last, _ := time.Parse("2006-01-02T15:04:05", v.Last)
		v.Interval = int(time.Now().Sub(last) / time.Second)
		v.Reads++
	} else {
		ServerStatus.clients[ua] = &ClientsStruct{
			First:    time.Now().Format("2006-01-02T15:04:05"),
			Interval: 0,
			Path:     []string{},
			Reads:    1,
			Kick:     "",
		}
	}

	ServerStatus.clients[ua].Last = time.Now().Format("2006-01-02T15:04:05")
	paths := ServerStatus.clients[ua].Path
	if len(paths) >= 4 {
		paths = paths[len(paths)-4:]
	}
	ServerStatus.clients[ua].Path = append(paths, path)
}

func ParseGet(c *gin.Context) {
	cmd := CheckQuery(c)
	if len(cmd) == 0 {
		GetDanmu(c)
		return
	}
	statement := map[string]func(c *gin.Context){
		`^/d+$`: func(c *gin.Context) {
			c.String(HTTP_OK, ChangeRoom(cmd))
		},
		`^history$`: GetHistory,
		`^restart$`: func(c *gin.Context) {
			ServerStatus.pop = 1
			c.String(HTTP_OK, "[RESTART] RECV OK")
			control <- ControlStruct{cmd: CMD_RESTART}
		},
		`^upgrade$`: func(c *gin.Context) {
			ServerStatus.pop = 1
			c.String(HTTP_OK, "[UPGRADE] Depends on network")
			control <- ControlStruct{cmd: CMD_RESTART}
		},
		`^status$`: GetStatus,
		`^clients$`: func(c *gin.Context) {
			c.JSON(HTTP_OK, ServerStatus.clients)
		},
		`^kick$`: func(c *gin.Context) {
			SetKick(c)
			GetDanmu(c)
		},
		`^call:`: func(c *gin.Context) {
			addDanmu(cmd[strings.Index(cmd, ":")+1:])
			c.String(HTTP_OK, "[CALLING]")
		},
		`^js:`: func(c *gin.Context) {
			if ServerStatus.pop == 1 {
				ServerStatus.pop = 9999
			}
			js := cmd[strings.Index(cmd, ":")+1:]
			addDanmu("[JS] " + js)
			c.String(HTTP_OK, "[JS-EXCUTING] "+js)
		},
		`^cors:`: func(c *gin.Context) {
			s, _ := c.GetRawData()
			c.String(HTTP_OK, CorsAccess(cmd[strings.Index(cmd, ":")+1:], string(s), "GET", c.Request.Header))
		},
		`^time$`: func(c *gin.Context) {
			c.String(HTTP_OK, strconv.FormatInt(time.Now().UnixMilli(), 10))
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
			c.String(HTTP_OK, CorsAccess(cmd[strings.Index(cmd, ":")+1:], string(s), "POST", c.Request.Header))
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
			c.String(HTTP_OK, CorsAccess(cmd[strings.Index(cmd, ":")+1:], string(s), "PUT", c.Request.Header))
		},
	}
	ApplyMatch(c, statement, cmd)
}
