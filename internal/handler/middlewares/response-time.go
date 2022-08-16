package middlewares

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type XResponseTimer struct {
	gin.ResponseWriter
	start time.Time
}

func (x *XResponseTimer) WriteHeader(statusCode int) {
	duration := time.Since(x.start).Microseconds()
	format := strconv.Itoa(int(duration))

	x.Header().Set("X-Response-Time", format)
	x.ResponseWriter.WriteHeader(statusCode)
}

func (x *XResponseTimer) Write(b []byte) (int, error) {
	return x.ResponseWriter.Write(b)
}

func NewXResponseTimer(c *gin.Context) {
	blw := &XResponseTimer{ResponseWriter: c.Writer, start: time.Now()}
	c.Writer = blw
	c.Next()
}
