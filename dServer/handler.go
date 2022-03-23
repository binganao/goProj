package main

import (
	"dServer/settings"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func CorsAccess(url string, data string, method string, ori_headers ...map[string][]string) string {
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36",
		"Accept":     "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	}
	if len(ori_headers) != 0 {
		if v, ok := ori_headers[0]["Content-Type"]; ok {
			headers["Content-Type"] = v[0]
		}
	}

	request, _ := http.NewRequest(method, url, strings.NewReader(data))
	for k, v := range headers {
		request.Header.Set(k, v)
	}
	client := &http.Client{Timeout: time.Second * time.Duration(settings.Timeout)}
	response, err := client.Do(request)
	//defer response.Body.Close()

	if err != nil {
		return "[ERR:cors] " + err.Error() + "@" + url
	} else if response.StatusCode >= 400 {
		return "[" + strconv.Itoa(response.StatusCode) + ":cors] " + url
	} else {
		body, _ := ioutil.ReadAll(response.Body)
		return string(body)
	}
}

func insertStatus() {
	js, _ := json.Marshal(GetServerStatus())
	History = append(History, "<!--"+string(js)+"-->")
}

//readfromlive

func GetServerStatus() gin.H {
	return gin.H{
		"room":           ServerStatus.room,
		"other_room":     ServerStatus.other_room,
		"pop":            ServerStatus.pop,
		"unread":         len(History) - ServerStatus.i,
		"status":         ServerStatus.status,
		"status_content": StatusList[ServerStatus.status],
	}
}

func ChangeRoom(room string) string {
	room_id, _ := strconv.Atoi(room)
	if room_id > 0 && room_id < 1e15 {
		if ServerStatus.other_room != "" || ServerStatus.room != room {
			fmt.Println("[kill:" + room + "]")
			ServerStatus.room = room
			ServerStatus.pop = 0
			ServerStatus.other_room = ""
			control <- ControlStruct{
				cmd:  CMD_CHANGE_ROOM,
				room: room,
			}
		} else {
			fmt.Println("[recv:butSame]")
		}
		return "[RECV] Room<b>" + room + "</b>"
	} else if room_id == 0 {
		return "[RECV] Room Keeps"
	} else {
		return "[err] Not in safe range: " + room
	}
}

func GetStatus(c *gin.Context) {
	c.JSON(HTTP_OK, GetServerStatus())
}

func Reverse(original []string) []string {
	out := make([]string, len(original))
	for i, k := range original {
		out[len(original)-i-1] = k
	}
	return out
}

func GetDanmu(c *gin.Context) {
	if CheckKick(c) {
		HTMLString(c, "[JS]danmuOff('KICKED')")
		return
	}
	res := ""
	for j := 0; j <= 1; j++ {
		i := ServerStatus.i
		if i < len(History) {
			res = strings.Join(History[i:], "<br>")
			ServerStatus.i = len(History)
			break
		} else if j == 0 {
			<-time.After(time.Second * 15)
		}
	}
	HTMLString(c, res)
}

func GetHistory(c *gin.Context) {
	limit := 2000
	former := len(History) - limit
	var res string
	if former > 0 {
		res = fmt.Sprintf("%s<details><summary>%d shown, %d left</summary>%s</details>", strings.Join(Reverse(History[former:]), "<br>"), limit, former, strings.Join(Reverse(History[:former]), "<br>"))
	} else {
		res = strings.Join(Reverse(History), "<br>")
	}
	HTMLString(c, res)
}

func GetFavicon(c *gin.Context) {
	c.Status(204)
}

func RecordClient(c *gin.Context) {
	ua := c.Request.UserAgent()
	path := c.Request.RequestURI
	if v, ok := ServerStatus.clients[ua]; ok {
		last, _ := time.Parse("2006-01-02T15:04:05MST", v.Last+time.Now().Format("MST"))
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

func SetKick(c *gin.Context) {
	ua := c.Request.UserAgent()
	if _, ok := ServerStatus.clients[ua]; ok {
		for i, j := range ServerStatus.clients {
			if i == ua {
				j.Kick = ""
			} else {
				j.Kick = time.Now().Format("2006-01-02T15:04:05")
			}
		}
	}
}

func CheckKick(c *gin.Context) bool {
	if k, ok := ServerStatus.clients[c.Request.UserAgent()]; ok && k.Kick != "" {
		expire, _ := time.Parse("2006-01-02T15:04:05", k.Kick)
		k.Kick = ""
		if time.Now().Sub(expire) < 120*time.Second {
			return true
		}
	}
	return false
}
