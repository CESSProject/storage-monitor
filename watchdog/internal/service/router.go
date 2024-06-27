package service

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", healthCheck)
	r.GET("/health_check", healthCheck)
	r.GET("/list", list)
	r.GET("/hosts", getHosts)
	r.GET("/clients", getClientsStatus)
	//r.Group("/").Use(localhostOnly()).POST("/config", setConfig)
	r.GET("/config", getConfig)
	r.POST("/config", setConfig)
	r.GET("/toggle", getToggle)
	r.POST("/toggle", setToggle)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func localhostOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.RemoteAddr != "127.0.0.1" && c.Request.RemoteAddr != "::1" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Access denied",
			})
			return
		}
		c.Next()
	}
}
