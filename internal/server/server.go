package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ss49919201/diary/internal/handler"
)

func Run() error {
	r := gin.New()

	r.GET("/health", func(c *gin.Context) {
		handler.HealthCheck(c.Request.Context(), c.Writer, c.Request)
	})

	return r.Run() // 0.0.0.0:8080
}
