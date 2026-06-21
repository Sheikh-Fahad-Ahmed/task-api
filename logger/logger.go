package logger

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func CustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		latency := time.Since(startTime)

		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path

		var err string
		if len(c.Errors) > 0 {
			err = " | Error: " + c.Errors.String()
		}

		log.Printf("endpoint %s | %v | %d | %s%s", method, latency, status, path, err)
	}
}
