package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/hoseazhai/gopl_learning/blog-service/global"
	"github.com/hoseazhai/gopl_learning/blog-service/pkg/logger"
	"time"
)

type AccessLogWrite struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWrite) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := AccessLogWrite{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWriter

		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()

		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}
		s := "access log: method: %s, status_code: %d, " +
			"begin_time: %d, end_time: %d"
		global.Logger.WithFields(fields).Infof(c, s,
			c.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime,
		)
	}

}
