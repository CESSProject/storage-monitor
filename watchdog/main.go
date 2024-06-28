package main

import (
	"github.com/CESSProject/watchdog/docs"
	"github.com/CESSProject/watchdog/internal/core"
	"github.com/CESSProject/watchdog/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	core.Run()
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}
	router.Use(cors.New(corsConfig))
	service.RegisterRoutes(router)
	err := router.Run(":13090")
	if err != nil {
		return
	}
}
