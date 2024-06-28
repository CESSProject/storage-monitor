package main

import (
	"github.com/CESSProject/watchdog/docs"
	"github.com/CESSProject/watchdog/internal/core"
	"github.com/CESSProject/watchdog/internal/log"
	"github.com/CESSProject/watchdog/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strconv"
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
	var httpPort string
	if core.CustomConfig.Server.External {
		httpPort = ":" + strconv.Itoa(core.CustomConfig.Server.Http.Port)
	} else {
		httpPort = "127.0.0.1:" + strconv.Itoa(core.CustomConfig.Server.Http.Port)
	}
	if core.CustomConfig.Server.Https.CertPath != "" && core.CustomConfig.Server.Https.KeyPath != "" {
		var httpsPort string
		if core.CustomConfig.Server.External {
			httpsPort = ":" + strconv.Itoa(core.CustomConfig.Server.Https.Port)
		} else {
			httpsPort = "127.0.0.1:" + strconv.Itoa(core.CustomConfig.Server.Https.Port)
		}
		go func() {
			if err := router.Run(httpPort); err != nil {
				log.Logger.Errorf("Fail to start watchdog server at %s : %v", httpPort, err)
			}
		}()
		go func() {
			if err := router.RunTLS(httpsPort, core.CustomConfig.Server.Https.CertPath, core.CustomConfig.Server.Https.KeyPath); err != nil {
				log.Logger.Errorf("Fail to start watchdog server at %s : %v", httpsPort, err)
			}
		}()
		select {}
	} else {
		err := router.Run(httpPort)
		if err != nil {
			return
		}
	}
}
