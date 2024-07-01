package service

import (
	"github.com/CESSProject/watchdog/constant"
	"github.com/CESSProject/watchdog/internal/core"
	"github.com/CESSProject/watchdog/internal/log"
	"github.com/CESSProject/watchdog/internal/model"
	"github.com/CESSProject/watchdog/internal/util"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"time"
)

// watchdog godoc
// @Schemes
// @Description Service HealthCheck
// @Tags HealthCheck
// @Success 200 {string} ok
// @Router / [get]
func healthCheck(c *gin.Context) {
	res := "ok"
	c.JSON(http.StatusOK, res)
}

// watchdog godoc
// @Description  List miners in each host
// @Tags         List Miners by host
// @Produce      json
// @Param        host   query  string   false  "Host IP"
// @Success      200  {object}  []MinerInfoVO
// @Router       /list  [get]
func list(c *gin.Context) {
	host := c.Query("host")
	data := getListByCondition(host)
	c.JSON(http.StatusOK, data)
}

// watchdog godoc
// @Description  List host
// @Tags         Get Hosts
// @Success      200  {object}  []string
// @Router       /hosts [get]
func getHosts(c *gin.Context) {
	res := make([]string, 0)
	for hostIP := range core.Clients {
		res = append(res, hostIP)
	}
	c.JSON(http.StatusOK, res)
}

// watchdog godoc
// @Description  Get Clients Status
// @Tags         Get Hosts
// @Success      200  {object} map[string]string
// @Router       /clients [get]
func getClientsStatus(c *gin.Context) {
	res := map[string]string{}
	for _, client := range core.Clients {
		if client.Updating {
			res[client.Host] = "Running"
		} else {
			res[client.Host] = "Sleep"
		}
	}
	c.JSON(http.StatusOK, res)
}

// watchdog godoc
// @Description  Update watchdog configuration
// @Tags         Update Config
// @Accept       json
// @Produce      json
// @Param        model.yamlConfig body model.YamlConfig true "YAML Configuration"
// @Success      200 {object} model.YamlConfig
// @Router       /config [post]
func setConfig(c *gin.Context) {
	var newConfig model.YamlConfig
	if err := c.ShouldBindJSON(&newConfig); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	configTemp, err := util.LoadConfigFile(constant.ConfPath)
	if err != nil {
		log.Logger.Errorf("Fail to load %s", constant.ConfPath)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Fail to load config file from /opt/monitor/config.yaml"})
		return
	}
	// remove old config
	util.RemoveFields(configTemp, "hosts", "scrapeInterval", "alert")
	// add new config
	util.AddFields(configTemp, newConfig)
	err = util.SaveConfigFile(constant.ConfPath, configTemp)
	if err != nil {
		log.Logger.Errorf("Fail to save file to: %v", constant.ConfPath)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Fail to save config file to /opt/monitor/config.yaml"})
		return
	}
	log.Logger.Infof("Save new config %v to: %s", configTemp, constant.ConfPath)
	go runWithNewConf()
	c.JSON(http.StatusOK, gin.H{"message": "update Watchdog config success"})
}

// watchdog godoc
// @Description  Get watchdog configuration
// @Tags         Get Config
// @Produce      json
// @Success      200 {object} model.YamlConfig
// @Router       /config [get]
func getConfig(c *gin.Context) {
	var conf model.YamlConfig
	conf = core.CustomConfig
	for i := 0; i < len(conf.Alert.Webhook); i++ {
		conf.Alert.Webhook[i] = splitURLByTopLevelDomain(conf.Alert.Webhook[i])
	}
	for i := 0; i < len(conf.Alert.Email.Receiver); i++ {
		conf.Alert.Email.Receiver[i] = replaceFirstThreeChars(conf.Alert.Email.Receiver[i])
	}
	conf.Alert.Email.SenderAddr = replaceFirstThreeChars(conf.Alert.Email.SenderAddr)
	conf.Alert.Email.SmtpPassword = "*"
	conf.Server = model.YamlConfig{}.Server
	c.JSON(http.StatusOK, conf)
}

// watchdog godoc
// @Description  Get Alert Status
// @Tags         Get Alert Status
// @Produce      json
// @Success      200 {object} bool
// @Router       /toggle [get]
func getToggle(c *gin.Context) {
	status := core.CustomConfig.Alert.Enable
	c.JSON(http.StatusOK, status)
}

// watchdog godoc
// @Description  Set Alert Status
// @Tags         Set Alert Status
// @Accept       json
// @Produce      json
// @Param        model.AlertToggle body model.AlertToggle true "Alert Toggle Status"
// @Success      200 {object} model.AlertToggle
// @Router       /toggle [post]
func setToggle(c *gin.Context) {
	var alertToggle = model.AlertToggle{}
	if err := c.ShouldBindJSON(&alertToggle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	conf := core.CustomConfig
	conf.Alert.Enable = alertToggle.Status
	data, err := yaml.Marshal(conf)
	if err != nil {
		log.Logger.Errorf("Fail to parse conf: %v", conf)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Fail to parse conf"})
		return
	}
	err = os.WriteFile(constant.ConfPath, data, 0644)
	if err != nil {
		log.Logger.Errorf("Fail to save file to: %v", constant.ConfPath)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Fail to save config file to /opt/monitor/config.yaml"})
		return
	}
	core.CustomConfig.Alert.Enable = alertToggle.Status
	log.Logger.Infof("Switch alert status to: %v", alertToggle.Status)
	c.JSON(http.StatusOK, gin.H{"message": "updateConfig alert status success"})
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
			for minerName := range VO.MinerInfoList {
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
		minerInfo.Conf.Mnemonic = "-"
		minerInfoArray = append(minerInfoArray, *minerInfo)
	}
	return minerInfoArray
}

func replaceFirstThreeChars(s string) string {
	// 123456@cess.cloud -> ***456@cess.cloud
	if len(s) < 5 {
		return s
	}
	return "***" + s[3:]
}

func splitURLByTopLevelDomain(inputURL string) string {
	// https://example.com/bot/v2/hook/4bb9bfc7-dat4-41g9-962d-d8b4c139f37c -> https://example.com/***
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		log.Logger.Warnf("Parse webhook url err: %v", err)
		return ""
	}
	hostname := parsedURL.Hostname()
	lastDotIndex := strings.LastIndex(hostname, ".")
	if lastDotIndex == -1 {
		log.Logger.Warnf("no top-level domain found in Webhook URL")
		return ""
	}
	res := parsedURL.Scheme + "://" + hostname + "/***"
	return res
}

func runWithNewConf() {
	for key := range core.Clients {
		core.Clients[key].Active = false
	}
	try := 0
	for {
		// max scrapeInterval is 300, sleep time is 5s, 300/5=60
		if try > 60 {
			break
		}
		if canProceed() {
			err := core.InitWatchdogConfig()
			if err != nil {
				return
			}
			core.InitSmtpConfig()
			core.InitWebhookConfig()
			err = core.InitWatchdogClients(core.CustomConfig)
			if err != nil {
				log.Logger.Fatalf("Init CESS Node Monitor Service With New Conf Failed: %v", err)
			}
			err = core.RunWatchdogClients(core.CustomConfig)
			if err != nil {
				log.Logger.Fatalf("Fail to run with new clients: %v", err)
			}
			break
		} else {
			log.Logger.Infof("Run with new config failed, Some watchdog clients might running, retrying (%d/%d)", try+1, 60)
		}
		try++
		time.Sleep(time.Duration(5) * time.Second)
	}
	log.Logger.Info("Run with new config success")
}

func canProceed() bool {
	for _, client := range core.Clients {
		if client.Updating {
			return false
		}
	}
	return true
}
