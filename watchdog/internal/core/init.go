package core

import (
	"github.com/CESSProject/watchdog/constant"
	"github.com/CESSProject/watchdog/internal/model"
	"github.com/CESSProject/watchdog/internal/util"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"math"
	"os"
)

type MinerInfoVO struct {
	Host          string
	MinerInfoList []MinerInfo
}

var CustomConfig model.YamlConfig

var SmtpConfigPoint *util.SmtpConfig
var WebhookConfigPoint *util.WebhookConfig

func Run() {
	err := InitMonitorConfig()
	if err != nil {
		return
	}
	InitSmtpConfig()
	InitWebhookConfig()
	log.Println("Service run with Host: ", CustomConfig.Hosts)
	err = InitWatchdogClients(CustomConfig)
	if err != nil {
		log.Fatal("Init docker clients failed: ", err)
	}
	err = RunWatchdogClients(CustomConfig)
	if err != nil {
		log.Fatal("Run docker clients failed", err)
		return
	}
}

func InitMonitorConfig() error {
	yamlFile, err := os.ReadFile(constant.ConfPath)
	if err != nil {
		log.Fatalf("Error when read file from /opt/monitor/config.yaml: %v", err)
		return err
	}
	CustomConfig = model.YamlConfig{}
	// yaml.Unmarshal
	//For string types, the zero value is the empty string "".
	//For numeric types, the zero value is 0.
	//For Boolean types, the zero value is false.
	//For pointer types, the zero value is nil.
	err = yaml.Unmarshal(yamlFile, &CustomConfig)
	if err != nil {
		log.Fatalf("Error when parse file from /opt/monitor/config.yaml: %v", err)
		return err
	}
	// 30 <= ScrapeInterval <= 600
	CustomConfig.ScrapeInterval = int(math.Max(30, math.Min(float64(CustomConfig.ScrapeInterval), 600)))
	log.Println("Service run with config file: \n", CustomConfig)
	return nil
}

func InitSmtpConfig() {
	if CustomConfig.Alert.Email.SmtpEndpoint == "" ||
		CustomConfig.Alert.Email.SmtpPort == 0 ||
		CustomConfig.Alert.Email.SenderAddr == "" ||
		CustomConfig.Alert.Email.SmtpPassword == "" ||
		len(CustomConfig.Alert.Email.Receiver) == 0 {
		return
	}
	SmtpConfigPoint = &util.SmtpConfig{
		SmtpUrl:      CustomConfig.Alert.Email.SmtpEndpoint,
		SmtpPort:     CustomConfig.Alert.Email.SmtpPort,
		SenderAddr:   CustomConfig.Alert.Email.SenderAddr,
		SmtpPassword: CustomConfig.Alert.Email.SmtpPassword,
		Receiver:     CustomConfig.Alert.Email.Receiver,
	}
}

func InitWebhookConfig() {
	if len(CustomConfig.Alert.Webhook) == 0 {
		return
	}
	WebhookConfigPoint = &util.WebhookConfig{
		Webhooks: CustomConfig.Alert.Webhook,
	}
}
