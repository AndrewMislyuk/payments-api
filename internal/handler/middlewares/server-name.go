package middlewares

import (
	"github.com/gin-gonic/gin"
)

type XServerName struct {
	gin.ResponseWriter
	Host string
}

func (x *XServerName) WriteHeader(statusCode int) {
	x.Header().Set("X-Server-Name", x.Host)
	x.ResponseWriter.WriteHeader(statusCode)
}

func (x *XServerName) Write(b []byte) (int, error) {
	return x.ResponseWriter.Write(b)
}

func NewXServerName(c *gin.Context) {
	blw := &XServerName{ResponseWriter: c.Writer, Host: c.Request.Host}
	c.Writer = blw
	c.Next()
}
