package service

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	r.GET("/list", list)
	r.GET("/hosts", getHosts)
	r.GET("/health_check", healthCheck)
	r.GET("/", healthCheck)
	r.POST("/update", update)
}
