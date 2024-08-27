package constant

const (
	MinerImage        = "cesslab/cess-miner"
	GenBlockInterval  = 6 // unit: second
	ConfPath          = "/opt/cess/watchdog/config.yaml"
	MinerConfPath     = "/opt/miner/config.yaml"
	HttpMaxRetry      = 3
	HttpRetryWaitTime = 5
	HttpTimeout       = 30
	TimeFormat        = "2006-01-02 15:04:05"
)

const (
	Size1kib = 1024
	Size1mib = 1024 * Size1kib
	Size1gib = 1024 * Size1mib
)

const (
	Unknown  = "unknown"
	Discord  = "discord"
	Slack    = "slack"
	Telegram = "telegram" // do not support now
	Teams    = "teams"
	Lark     = "lark"
	DingTalk = "ding"
	WeChat   = "wechat"
)

const (
	HttpPostContentType = "application/json"
	ScanApiUrl          = "https://scan-api.cess.network"
	DefaultDescription  = "The Storage Node is not in a positive status or has received punishment"
	DefaultURL          = "https://scan.cess.network"
	ScanAccountURL      = "https://scan.cess.network/account/"
	ScanBlockURL        = "https://scan.cess.network/block/"
	ScanExtrinsicURL    = "https://scan.cess.network/extrinsic/"
	DefaultRpcUrl       = "wss://testnet-rpc.cess.network/ws/"
	LocalRpcUrl         = "ws://127.0.0.1/ws/"
)

const (
	MinerStatus          = "Frozen"
	NoSubmitSvcProof     = "NoSubmitSvcProof"
	SvcProofResIncorrect = "SvcProofResIncorrect"
	AlertStaticPath      = "/opt/cess/watchdog/alert/"
)
