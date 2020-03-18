package config

type Config struct {
	LogLevel string `mapstructure:"log_level"`
	P2P      *P2P   `mapstructure:"p2p"`
	API      *API   `mapstructure:"api"`
}

type P2P struct {
	ListenAddr    string `mapstructure:"listen_addr"`
	BootstrapPeer string `mapstructure:"bootstrap_peer"`
	PrivateKey    string `mapstructure:"private_key"`
}

type API struct {
	Port int `mapstructure:"port"`
}
