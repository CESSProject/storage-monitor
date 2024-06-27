package util

import (
	"fmt"
	"github.com/CESSProject/cess-go-sdk/chain"
	"github.com/CESSProject/watchdog/constant"
	"github.com/CESSProject/watchdog/internal/model"
	"github.com/btcsuite/btcutil/base58"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"gopkg.in/yaml.v3"
	"math/big"
	"net"
	"os"
)

func ParseMinerConfigFile(data []byte) (model.MinerConfigFile, error) {
	var conf model.MinerConfigFile
	err := yaml.Unmarshal(data, &conf) // unmarshal = parse
	if err != nil {
		return conf, err
	}
	return conf, nil
}

func TransferMinerInfoToMinerStat(info chain.MinerInfo) (model.MinerStat, error) {
	var minerStat = model.MinerStat{}
	minerStat.PeerId = base58.Encode([]byte(string(info.PeerId[:])))
	minerStat.Collaterals = BigNumConversion(info.Collaterals)
	minerStat.Debt = BigNumConversion(info.Debt)
	minerStat.Status = string(info.State)
	minerStat.DeclarationSpace = StorageSpaceUnitConversion(info.DeclarationSpace)
	minerStat.IdleSpace = StorageSpaceUnitConversion(info.IdleSpace)
	minerStat.ServiceSpace = StorageSpaceUnitConversion(info.ServiceSpace)
	minerStat.LockSpace = StorageSpaceUnitConversion(info.LockSpace)
	minerStat.IsPunished = make([][]bool, 0)
	return minerStat, nil
}

func BigNumConversion(value types.U128) string {
	bigIntValue, ok := new(big.Int).SetString(value.String(), 10)
	if !ok {
		return ""
	}
	bigRatValue := new(big.Rat).SetInt(bigIntValue)
	divisor := new(big.Rat).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
	result := new(big.Rat).Quo(bigRatValue, divisor)
	resultFloat, _ := result.Float64()
	return fmt.Sprintf("%.4f", resultFloat)
}

func StorageSpaceUnitConversion(value types.U128) string {
	var result string
	if value.IsUint64() {
		v := value.Uint64()
		if v >= (constant.Size1gib * 1024 * 1024 * 1024) {
			result = fmt.Sprintf("%.2f EiB", float64(v)/float64(constant.Size1gib*1024*1024*1024))
			return result
		}
		if v >= (constant.Size1gib * 1024 * 1024) {
			result = fmt.Sprintf("%.2f PiB", float64(v)/float64(constant.Size1gib*1024*1024))
			return result
		}
		if v >= (constant.Size1gib * 1024) {
			result = fmt.Sprintf("%.2f TiB", float64(v)/float64(constant.Size1gib*1024))
			return result
		}
		if v >= (constant.Size1gib) {
			result = fmt.Sprintf("%.2f GiB", float64(v)/float64(constant.Size1gib))
			return result
		}
		if v >= (constant.Size1mib) {
			result = fmt.Sprintf("%.2f MiB", float64(v)/float64(constant.Size1mib))
			return result
		}
		if v >= (constant.Size1kib) {
			result = fmt.Sprintf("%.2f KiB", float64(v)/float64(constant.Size1kib))
			return result
		}
		result = fmt.Sprintf("%v Bytes", v)
		return result
	}
	v := new(big.Int).SetBytes(value.Bytes())
	v.Quo(v, new(big.Int).SetUint64(constant.Size1gib*1024*1024*1024))
	result = fmt.Sprintf("%v EiB", v)
	return result
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
	config["scrapeInterval"] = conf.ScrapeInterval
	config["hosts"] = conf.Hosts
	config["alert"] = conf.Alert
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
