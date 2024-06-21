package util

import (
	"github.com/CESSProject/cess-go-sdk/chain"
	"github.com/CESSProject/watchdog/internal/model"
	"github.com/btcsuite/btcutil/base58"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"net"
	"os"
)

func ParseMinerConfigFile(data []byte) (model.MinerConfigFile, error) {
	var conf model.MinerConfigFile
	err := yaml.Unmarshal(data, &conf) // unmarshal = parse
	if err != nil {
		return conf, errors.Wrap(err, "parse miner config file error")
	}
	return conf, nil
}

func TransferMinerInfoToMinerStat(info chain.MinerInfo) (model.MinerStat, error) {
	var minerStat = model.MinerStat{}
	minerStat.PeerId = base58.Encode([]byte(string(info.PeerId[:])))
	minerStat.Collaterals = info.Collaterals.Int
	minerStat.Debt = info.Debt.Int
	minerStat.Status = string(info.State)
	minerStat.DeclarationSpace = info.DeclarationSpace.Int
	minerStat.IdleSpace = info.IdleSpace.Int
	minerStat.ServiceSpace = info.ServiceSpace.Int
	minerStat.LockSpace = info.LockSpace.Int
	minerStat.IsPunished = make([][]bool, 0)
	return minerStat, nil
}

func LoadConfigFile(filePath string) (map[interface{}]interface{}, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var configTemp map[interface{}]interface{}
	err = yaml.Unmarshal(data, &configTemp)
	if err != nil {
		return nil, err
	}

	return configTemp, nil
}

func RemoveFields(config map[interface{}]interface{}, fields ...string) {
	for _, field := range fields {
		delete(config, field)
	}
}

func AddFields(config map[interface{}]interface{}, conf model.YamlConfig) {
	config["ScrapeInterval"] = conf.ScrapeInterval
	config["Host"] = conf.Hosts
	config["Alert"] = conf.Alert
}

func SaveConfigFile(filePath string, config map[interface{}]interface{}) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}

func IsPrivateIP(ip net.IP) bool {
	if ip.IsLoopback() {
		return true
	}
	privateIPBlocks := []*net.IPNet{
		{IP: net.ParseIP("10.0.0.0"), Mask: net.CIDRMask(8, 32)},     // 10.0.0.0/8
		{IP: net.ParseIP("172.16.0.0"), Mask: net.CIDRMask(12, 32)},  // 172.16.0.0/12
		{IP: net.ParseIP("192.168.0.0"), Mask: net.CIDRMask(16, 32)}, // 192.168.0.0/16
	}
	for _, block := range privateIPBlocks {
		if block.Contains(ip) {
			return true
		}
	}
	return false
}
