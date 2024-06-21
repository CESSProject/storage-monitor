package model

import (
	"math/big"
)

type HostItem struct {
	IP       string `yaml:"ip,omitempty"`        // host ip
	Port     string `yaml:"port,omitempty"`      // docker api port
	CAPath   string `yaml:"ca_path,omitempty"`   // /etc/docker/127.0.0.1/ca.pem
	CertPath string `yaml:"cert_path,omitempty"` // /etc/docker/127.0.0.1/cert.pem
	KeyPath  string `yaml:"key_path,omitempty"`  // /etc/docker/127.0.0.1/key.pem
}

type MailContent struct {
	AlertTime     string
	HostIp        string
	ContainerName string
	Description   string
}

type WebhookContent struct {
	MsgType string `json:"msg_type"`
	Content struct {
		Text string `json:"text"`
	} `json:"content"`
}

type Container struct {
	ID            string   `json:"id"`
	Names         []string `json:"names"`
	Name          string   `json:"name"`
	Image         string   `json:"image"`
	ImageID       string   `json:"image_id"`
	Command       string   `json:"command"`
	Created       int64    `json:"created"`
	State         string   `json:"state"`
	Status        string   `json:"status"`
	CPUPercent    float64  `json:"cpu_percent"`
	MemoryPercent float64  `json:"memory_percent"`
	MemoryUsage   uint64   `json:"mem_usage"`
}

type ContainerStat struct {
	CPUPercent    float64
	MemoryPercent float64
	MemoryUsage   uint64
}

type MinerStat struct {
	PeerId           string   `json:"peer_id"`
	Collaterals      *big.Int `json:"collaterals"`
	Debt             *big.Int `json:"debt"`
	Status           string   `json:"status"`            // positive, exit, frozen, unready(register on chain but no get a tag from tee)
	DeclarationSpace *big.Int `json:"declaration_space"` // unit: TiB
	IdleSpace        *big.Int `json:"idle_space"`
	ServiceSpace     *big.Int `json:"service_space"`
	LockSpace        *big.Int `json:"lock_space"` // upload file allocated to this miner but not get a proof from tee yet, it can be serviceSpace after get proof from tee
	IsPunished       [][]bool `json:"is_punished"`
	TotalReward      *big.Int `json:"total_reward"`
	RewardIssued     *big.Int `json:"reward_issued"`
}

type MinerConfigFile struct {
	Name        string   `name:"Name" toml:"Name" yaml:"Name"`
	Port        int      `name:"Port" toml:"Port" yaml:"Port"`
	EarningsAcc string   `name:"EarningsAcc" toml:"EarningsAcc" yaml:"EarningsAcc"`
	StakingAcc  string   `name:"StakingAcc" toml:"StakingAcc" yaml:"StakingAcc"`
	Mnemonic    string   `name:"Mnemonic" toml:"Mnemonic" yaml:"Mnemonic"`
	Rpc         []string `name:"Rpc" toml:"Rpc" yaml:"Rpc"`
	UseSpace    uint64   `name:"UseSpace" toml:"UseSpace" yaml:"UseSpace"`
	Workspace   string   `name:"Workspace" toml:"Workspace" yaml:"Workspace"`
	UseCpu      uint8    `name:"UseCpu" toml:"UseCpu" yaml:"UseCpu"`
	TeeList     []string `name:"TeeList" toml:"TeeList" yaml:"TeeList"`
	Boot        []string `name:"Boot" toml:"Boot" yaml:"Boot"`
}

type YamlConfig struct {
	Hosts          []HostItem `yaml:"hosts"`
	ScrapeInterval int        `yaml:"scrapeInterval,omitempty"`
	Alert          struct {
		Enable  bool     `yaml:"enable"`
		Webhook []string `yaml:"webhook"`
		Email   struct {
			SmtpEndpoint string   `yaml:"smtp_endpoint"`
			SmtpPort     string   `yaml:"smtp_port"`
			SenderAddr   string   `yaml:"sender_addr"`
			SmtpPassword string   `yaml:"smtp_password"`
			Receiver     []string `yaml:"receiver"`
		} `yaml:"email"`
	} `yaml:"alert"`
}
