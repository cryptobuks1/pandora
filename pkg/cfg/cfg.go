package cfg

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const configFilename = "pandora"

var Cfg = &Config{
	P2P: &P2P{},
}

func Prepare(cmd *cobra.Command) error {
	v := viper.New()

	v.SetConfigName(configFilename)
	v.AddConfigPath("./")
	v.AddConfigPath("./contrib/")

	_ = v.BindPFlag("p2p.listen_addr", cmd.Flags().Lookup("listen"))
	_ = v.BindPFlag("p2p.bootstrap_peer", cmd.Flags().Lookup("boot"))

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil
		}
		return err
	}
	if err := v.Unmarshal(&Cfg); err != nil {
		return err
	}

	return nil
}
