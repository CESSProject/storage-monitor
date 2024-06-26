package service

import (
	"github.com/CESSProject/watchdog/constant"
	"github.com/CESSProject/watchdog/internal/core"
	"github.com/CESSProject/watchdog/internal/log"
	"github.com/CESSProject/watchdog/internal/model"
	"github.com/CESSProject/watchdog/internal/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
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
	log.Logger.Infof("Update config to: %v", newConfig)
	configTemp, err := util.LoadConfigFile("/opt/cess/mineradm/config.yaml")
	if err != nil {
		log.Logger.Errorf("Fail to load /opt/cess/mineradm/config.yaml")
	}

	// remove old config
	util.RemoveFields(configTemp, "Hosts", "ScrapeInterval", "Alert")
	// add new config
	util.AddFields(configTemp, newConfig)

	err = util.SaveConfigFile(constant.ConfPath, configTemp)
	if err != nil {
		log.Logger.Errorf("Fail to save /opt/cess/mineradm/config.yaml")
	}
	c.JSON(http.StatusOK, gin.H{"message": "update conf success"})
}

type MinerInfoVO struct {
	Host          string
	MinerInfoList []core.MinerInfo
}

func getListByCondition(hostIp string) []MinerInfoVO {
	var res []MinerInfoVO
	if hostIp != "" {
		if _, ok := core.Clients[hostIp]; ok {
			VO := MinerInfoVO{
				Host:          hostIp,
				MinerInfoList: getMinersListByClientInfo(core.Clients[hostIp].MinerInfoMap),
			}
			for minerName, _ := range VO.MinerInfoList {
				VO.MinerInfoList[minerName].Conf.Mnemonic = ""
			}
			res = make([]MinerInfoVO, 0)
			res = append(res, VO)
			sort.Slice(res, func(i, j int) bool {
				return res[i].Host < res[j].Host
			})
		} else {
			log.Logger.Warnf("Host IP not found: %s", hostIp)
		}
	} else {
		res = make([]MinerInfoVO, 0)
		for k, v := range core.Clients {
			VO := MinerInfoVO{
				Host:          k,
				MinerInfoList: getMinersListByClientInfo(v.MinerInfoMap),
			}
			res = append(res, VO)
		}
		sort.Slice(res, func(i, j int) bool {
			return res[i].Host < res[j].Host
		})
	}
	return res
}

func getMinersListByClientInfo(minerMap map[string]*core.MinerInfo) []core.MinerInfo {
	var minerInfoArray []core.MinerInfo
	for _, minerInfo := range minerMap {
		minerInfoArray = append(minerInfoArray, *minerInfo)
	}
	return minerInfoArray
}
