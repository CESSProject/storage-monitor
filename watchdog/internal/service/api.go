package service

import (
	"github.com/CESSProject/watchdog/constant"
	"github.com/CESSProject/watchdog/internal/core"
	"github.com/CESSProject/watchdog/internal/model"
	"github.com/CESSProject/watchdog/internal/util"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func healthCheck(c *gin.Context) {
	res := "ok"
	c.JSON(200, res)
}

// get miner info list
func list(c *gin.Context) {
	host := c.Query("host")
	data := getListByCondition(host)
	c.JSON(200, data)
}

// get host ip list
func getHosts(c *gin.Context) {
	res := make([]string, 0)
	for hostIP, _ := range core.Clients {
		res = append(res, hostIP)
	}
	c.JSON(200, res)
}

// update monitor service' configuration
func update(c *gin.Context) {
	var newConfig = model.YamlConfig{}
	if err := c.ShouldBindJSON(&newConfig); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("Update config to: \n", newConfig)
	configTemp, err := util.LoadConfigFile("/opt/cess/mineradm/config.yaml")
	if err != nil {
		log.Fatalf("Fail to load /opt/cess/mineradm/config.yaml")
	}

	// remove old config
	util.RemoveFields(configTemp, "Hosts", "ScrapeInterval", "Alert")
	// add new config
	util.AddFields(configTemp, newConfig)

	err = util.SaveConfigFile(constant.ConfPath, configTemp)
	if err != nil {
		log.Fatalf("Fail to save /opt/cess/mineradm/config.yaml")
	}
	c.JSON(http.StatusOK, gin.H{"message": "update conf success"})
}
