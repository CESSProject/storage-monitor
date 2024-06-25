package model

type HostItem struct {
	IP       string `yaml:"ip"`                  // host ip
	Port     string `yaml:"port"`                // docker api port
	CAPath   string `yaml:"ca_path,omitempty"`   // /etc/docker/127.0.0.1/ca.pem
	CertPath string `yaml:"cert_path,omitempty"` // /etc/docker/127.0.0.1/cert.pem
	KeyPath  string `yaml:"key_path,omitempty"`  // /etc/docker/127.0.0.1/key.pem
}

type AlertContent struct {
	AlertTime     string
	HostIp        string
	ContainerName string
	Description   string
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
	Collaterals      string   `json:"collaterals"`
	Debt             string   `json:"debt"`
	Status           string   `json:"status"`            // positive, exit, frozen, unready(register on chain but no get a tag from tee)
	DeclarationSpace string   `json:"declaration_space"` // unit: TiB
	IdleSpace        string   `json:"idle_space"`
	ServiceSpace     string   `json:"service_space"`
	LockSpace        string   `json:"lock_space"` // upload file allocated to this miner but not get a proof from tee yet, it can be serviceSpace after get proof from tee
	IsPunished       [][]bool `json:"is_punished"`
	TotalReward      string   `json:"total_reward"`
	RewardIssued     string   `json:"reward_issued"`
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
			SmtpPort     int      `yaml:"smtp_port"`
			SenderAddr   string   `yaml:"smtp_account"`
			SmtpPassword string   `yaml:"smtp_password"`
			Receiver     []string `yaml:"receiver"`
		} `yaml:"email"`
	} `yaml:"alert"`
}
