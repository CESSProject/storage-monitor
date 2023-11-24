package main

import (
	"github.com/CESSProject/node-monitor/service"
	"github.com/gin-gonic/gin"
)

func main() {
	service.InitCache()
	router := gin.Default()
	service.RegisterRoutes(router)
	router.Run(":8088")
}
