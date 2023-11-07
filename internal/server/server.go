package server

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ss49919201/diary/internal/dicontainer"
	"github.com/ss49919201/diary/internal/handler"
)

func Run() error {
	r := gin.New()

	// log request handling elapsed time
	r.Use(logRequestHandlingElapsedTime)

	r.GET("/health", func(c *gin.Context) {
		dicontainer.MustInvoke[handler.HealthCheck]().Do(c.Writer, c.Request)
	})

	r.GET("/articles", func(c *gin.Context) {
		dicontainer.MustInvoke[handler.ListArticles]().Do(c.Writer, c.Request)
	})

	return r.Run() // 0.0.0.0:8080
}

func logRequestHandlingElapsedTime(ctx *gin.Context) {
	start := time.Now()
	ctx.Next()
	log.Printf("elapsed: %dns\n", time.Since(start).Nanoseconds())
}
