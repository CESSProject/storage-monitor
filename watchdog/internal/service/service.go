package service

import (
	"github.com/CESSProject/watchdog/constant"
	"github.com/CESSProject/watchdog/internal/core"
	"github.com/CESSProject/watchdog/internal/model"
	"github.com/CESSProject/watchdog/internal/util"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"math"
	"os"
	"sort"
)

type MinerInfoVO struct {
	Host          string
	MinerInfoList []core.MinerInfo
}

var CustomConfig model.YamlConfig

var SmtpConfig util.SmtpConfig

func Init() {
	err := ReadConfigFile()
	if err != nil {
		log.Fatal("Failed to read config file: ", constant.ConfPath)
		return
	}
	SmtpConfig = InitSmtpConfig()
	log.Println("Miners Host: ", CustomConfig.Hosts)
	err = core.InitWatchdogClients(CustomConfig)
	if err != nil {
		log.Fatal("Init docker clients failed: ", err)
	}
	err = core.RunWatchdogClients(CustomConfig)
	if err != nil {
		log.Fatal("Run docker clients failed", err)
		return
	}
}

func ReadConfigFile() error {
	yamlFile, err := os.ReadFile(constant.ConfPath)
	if err != nil {
		log.Fatalf("Error when read file from /opt/monitor/config.yaml: %v", err)
		return err
	}
	CustomConfig = model.YamlConfig{}
	err = yaml.Unmarshal(yamlFile, &CustomConfig)
	if err != nil {
		log.Fatalf("Error when parse file from /opt/monitor/config.yaml: %v", err)
		return err
	}
	// 30 <= ScrapeInterval <= 600
	CustomConfig.ScrapeInterval = int(math.Max(30, math.Min(float64(CustomConfig.ScrapeInterval), 600)))
	log.Println("Config file content: \n", CustomConfig)
	return nil
}

func InitSmtpConfig() util.SmtpConfig {
	c := util.SmtpConfig{
		SmtpUrl:      CustomConfig.Alert.Email.SmtpEndpoint,
		SmtpPort:     CustomConfig.Alert.Email.SmtpPort,
		SenderAddr:   CustomConfig.Alert.Email.SenderAddr,
		SmtpPassword: CustomConfig.Alert.Email.SmtpPassword,
		Receiver:     CustomConfig.Alert.Email.Receiver,
	}
	return c
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
			log.Println("HostIp not found")
		}
	} else {
		res = make([]MinerInfoVO, 0)
		log.Println("Get All Host Miners Info")
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
