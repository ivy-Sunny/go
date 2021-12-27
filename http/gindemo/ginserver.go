package http

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func MainGin() {
	r := gin.Default()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	/**
	middleware
	*/
	r.Use(func(c *gin.Context) {
		now := time.Now()
		c.Next()
		/**
		path, log latency, response code
		*/
		logger.Info(
			"incoming request",
			zap.String("path", c.Request.URL.Path),
			zap.Int("statis", c.Writer.Status()),
			zap.Duration("elapsed", time.Now().Sub(now)),
		)
	}, func(c *gin.Context) {
		c.Set("requestId", rand.Int())
		c.Next()
	})

	r.GET("/ping", func(c *gin.Context) {
		h := gin.H{"message": "pong"}

		if rid, exists := c.Get("requestId"); exists {
			h["requestId"] = rid
		}
		c.JSON(200, h)
	})

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})

	r.Run()
}
