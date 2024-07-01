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
	CPUPercent    string   `json:"cpu_percent"`
	MemoryPercent string   `json:"memory_percent"`
	MemoryUsage   string   `json:"mem_usage"`
}

type ContainerStat struct {
	CPUPercent    string
	MemoryPercent string
	MemoryUsage   string
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
	UseSpace    int      `name:"UseSpace" toml:"UseSpace" yaml:"UseSpace"`
	Workspace   string   `name:"Workspace" toml:"Workspace" yaml:"Workspace"`
	UseCpu      int      `name:"UseCpu" toml:"UseCpu" yaml:"UseCpu"`
	TeeList     []string `name:"TeeList" toml:"TeeList" yaml:"TeeList"`
	Boot        []string `name:"Boot" toml:"Boot" yaml:"Boot"`
}

type YamlConfig struct {
	Server struct {
		External bool `yaml:"external" json:"external"`
		Http     struct {
			Port int `yaml:"http_port" json:"http_port"`
		} `yaml:"http"`
		Https struct {
			Port     int    `yaml:"https_port" json:"https_port"`
			CertPath string `yaml:"cert_path" json:"cert_path"`
			KeyPath  string `yaml:"key_path" json:"key_path"`
		} `yaml:"https"`
	} `yaml:"server" json:"server"`
	Hosts          []HostItem `yaml:"hosts" json:"hosts"`
	ScrapeInterval int        `yaml:"scrapeInterval" json:"scrapeInterval"`
	Alert          struct {
		Enable  bool     `yaml:"enable" json:"enable"`
		Webhook []string `yaml:"webhook,omitempty" json:"webhook,omitempty"`
		Email   struct {
			SmtpEndpoint string   `yaml:"smtp_endpoint,omitempty" json:"smtp_endpoint,omitempty"`
			SmtpPort     int      `yaml:"smtp_port,omitempty" json:"smtp_port,omitempty"`
			SenderAddr   string   `yaml:"smtp_account,omitempty" json:"smtp_account,omitempty"`
			SmtpPassword string   `yaml:"smtp_password,omitempty" json:"smtp_password,omitempty"`
			Receiver     []string `yaml:"receiver,omitempty" json:"receiver,omitempty"`
		} `yaml:"email"`
	} `yaml:"alert" json:"alert"`
}

type AlertToggle struct {
	Status bool `name:"enable"`
}
