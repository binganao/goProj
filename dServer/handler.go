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

func CorsAccess(url string, data string, method string, ori_headers map[string][]string) string {
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36",
		"Accept":     "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	}
	if v, ok := ori_headers["Content-Type"]; ok {
		headers["Content-Type"] = v[0]
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

func ChangeRoom(room string) (res string) {
	room_id, _ := strconv.Atoi(room)
	if room_id > 0 && room_id < 1e15 {
		res = "[RECV] Room<b>" + room + "</b>"
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
	} else if room_id == 0 {
		res = "[RECV] Room Keeps"
	} else {
		res = "[err] Not in safe range: " + room
	}
	return
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
	SetHeaderHTML(c)
	res := ""
	i := ServerStatus.i
	if i < len(History) {
		res = strings.Join(History[i:], "<br>")
	}
	c.String(HTTP_OK, res)
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
	c.String(HTTP_OK, res)
}

func GetFavicon(c *gin.Context) {
	c.Status(204)
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
