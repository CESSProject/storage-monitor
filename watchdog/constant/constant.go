package constant

import "time"

const (
	MinerImage       = "cesslab/cess-miner"
	GenBlockInterval = 6 // unit: second
	ConfPath         = "/opt/monitor/config.yaml"
)

const (
	LogMaxAge    = 86400 * time.Second
	RotationTime = 604800 * time.Second
)
