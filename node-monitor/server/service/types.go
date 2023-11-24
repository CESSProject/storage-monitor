package service

type MinerInfoDisplay struct {
	ContainerInfo Container     `json:"container_info"`
	ContainerStat ContainerStat `json:"container_status"`
	Metadata      MinerMetadata `json:"miner_metadata"`
}

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

type MinerMetadata struct {
	Name            string `json:"name"`
	PeerId          string `json:"peer_id"`
	State           string `json:"state"`
	StakingAmount   string `json:"staking_amount"`
	ValidatedSpace  uint64 `json:"validated_space"`
	UsedSpace       uint64 `json:"used_space"`
	LockedSpace     uint64 `json:"locked_space"`
	StakingAccount  string `json:"staking_account"`
	EarningsAccount string `json:"earnings_account"`
}
