package cfg

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Cfg = &Config{
	P2P: &P2P{},
	API: &API{},
}

func Prepare(cmd *cobra.Command) error {
	v := viper.New()

	_ = v.BindPFlag("p2p.listen_addr", cmd.Flags().Lookup("listen"))
	_ = v.BindPFlag("p2p.bootstrap_peer", cmd.Flags().Lookup("boot"))
	_ = v.BindPFlag("p2p.private_key", cmd.Flags().Lookup("private"))
	_ = v.BindPFlag("log_level", cmd.Flags().Lookup("log_level"))
	_ = v.BindPFlag("api.port", cmd.Flags().Lookup("http_port"))

	if err := v.Unmarshal(&Cfg); err != nil {
		return err
	}

	return nil
}
