package handler

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *CustomerHandler, middleware ...gin.HandlerFunc) *gin.Engine {
	r := gin.Default()

	for _, m := range middleware {
		r.Use(m)
	}

	r.POST("/customers", handler.Create)
	r.GET("/customers/:id", handler.GetByID)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy"})
	})
	r.GET("/ready", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ready"})
	})

	return r
}
