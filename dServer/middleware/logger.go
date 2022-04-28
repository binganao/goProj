package middleware

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

var (
	Logger   = log.New()
	LogEntry *log.Entry
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		spend := fmt.Sprintf("%.1fms", float64(time.Since(start).Microseconds())/1000)
		status := c.Writer.Status()
		clientIP := c.ClientIP()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method
		url := c.Request.URL.RequestURI()
		Log := Logger.WithFields(log.Fields{
			"Use":      spend,
			"Path":     url,
			"Method":   method,
			"ClientIP": clientIP,
		})
		if len(c.Errors) > 0 {
			Log.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if status >= 500 {
			Log.Error()
		} else if status >= 400 {
			Log.Warn()
		} else {
			Log.Info()
		}
	}
}

func init() {
	LogEntry = log.NewEntry(Logger).WithField("service", "dServer")
}
