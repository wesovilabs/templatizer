package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wesovilabs/templatizer/pkg/handlers"
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Disposition, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func SetUpRouter(basePath string) *gin.Engine {
	router := gin.Default()
	router.Use(corsMiddleware())
	v1 := router.Group(basePath)
	{
		v1.POST("/parameters", handlers.LoadParamaters)
		v1.POST("/template", handlers.ProcessTemplate)
	}
	return router
}
