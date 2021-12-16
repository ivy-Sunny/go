package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	logger.Warn("warning test")
	gin.New()
}
