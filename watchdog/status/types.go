package status

type Container struct {
	ID      string   `json:"id"`
	Names   []string `json:"names"`
	Name    string   `json:"name"`
	Image   string   `json:"image"`
	ImageID string   `json:"image_id"`
	Command string   `json:"command"`
	Created int64    `json:"created"`
	State   string   `json:"state"`
	Status  string   `json:"status"`
}

type ContainerStat struct {
	ID            string `json:"id"`
	CPUPercent    int64  `json:"cpu"`
	MemoryPercent int64  `json:"memory"`
	MemoryUsage   int64  `json:"mem_usage"`
}

type Configfile struct {
	Rpc         []string `name:"Rpc" toml:"Rpc" yaml:"Rpc"`
	Boot        []string `name:"Boot" toml:"Boot" yaml:"Boot"`
	Mnemonic    string   `name:"Mnemonic" toml:"Mnemonic" yaml:"Mnemonic"`
	EarningsAcc string   `name:"EarningsAcc" toml:"EarningsAcc" yaml:"EarningsAcc"`
	Workspace   string   `name:"Workspace" toml:"Workspace" yaml:"Workspace"`
	Port        int      `name:"Port" toml:"Port" yaml:"Port"`
	UseSpace    uint64   `name:"UseSpace" toml:"UseSpace" yaml:"UseSpace"`
	UseCpu      uint8    `name:"UseCpu" toml:"UseCpu" yaml:"UseCpu"`
}
