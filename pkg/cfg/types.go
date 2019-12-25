package cfg

type Config struct {
	LogLevel string `mapstructure:"log_level"`
	P2P      *P2P   `mapstructure:"p2p"`
}

type P2P struct {
	ListenAddr    string `mapstructure:"listen_addr"`
	BootstrapPeer string `mapstructure:"bootstrap_peer"`
}
