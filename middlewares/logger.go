package middlewares

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	file, err := os.OpenFile("gin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logger := log.New(file, "", log.LstdFlags)

	return func(c *gin.Context) {
		// Before processing the request
		start := time.Now()

		// Process the request
		c.Next()

		// After processing the request
		end := time.Now()
		latency := end.Sub(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		path := c.Request.URL.Path

		// Log base request info
		logger.Printf("ClientIP: %s Method: %s Path: %s Status: %d Latency: %v", clientIP, method, path, statusCode, latency)

		// Log error if one occurred
		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				logger.Printf("Error in %v: %s", e.Meta, e.Err)
			}
		}
	}
}
