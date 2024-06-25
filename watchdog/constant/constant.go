package constant

import "time"

const (
	MinerImage       = "cesslab/cess-miner"
	GenBlockInterval = 6 // unit: second
	ConfPath         = "/opt/monitor/config.yaml"
	HttpMaxRetry     = 3
	TimeFormat       = "2006-01-02 15:04:05"
)

const (
	LogMaxAge    = 86400 * time.Second
	RotationTime = 604800 * time.Second
)

const (
	Size1kib = 1024
	Size1mib = 1024 * Size1kib
	Size1gib = 1024 * Size1mib
	Size1tib = 1024 * Size1gib
)
