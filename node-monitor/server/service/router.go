package service

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	r.Static("/app", "./dist")
	r.GET("/miners", GetCacheData)
	r.POST("/push", ReceivedCacheData)
	r.GET("/container", func(c *gin.Context) {
		c.File("./dist/index.html")
	})
}
