package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type PageResult struct {
	Total int64       `json:"total"`
	List  interface{} `json:"list"`
}

func Success(message string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{200, message, data})
}

func SuccessPage(message string, data interface{}, rows int64, c *gin.Context) {
	page := &PageResult{rows, data}
	c.JSON(http.StatusOK, Response{200, message, page})
}

func Failed(message string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{400, message, 0})
}
