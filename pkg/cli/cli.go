package cli

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"pandora/pkg/cfg"
	"pandora/pkg/logger"
	"pandora/pkg/node"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if err := cfg.Prepare(cmd); err != nil {
				log.Fatal().Err(err).Msg("prepare config file error")
			}

			logger.Initialize()
		},

		Run: func(cmd *cobra.Command, args []string) {
			nodeService, err := node.NewService()
			if err != nil {
				log.Fatal().Err(err).Msg("initialize node service error")
			}

			go func() {
				nodeService.Start()
			}()

			interrupt := make(chan os.Signal, 1)
			signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
			<-interrupt
			log.Info().Msg("handle SIGINT, SIGTERM, SIGQUIT")

			nodeService.Stop()
		},
	}

	cmd.PersistentFlags().String("listen", "", "listen address")
	cmd.PersistentFlags().String("boot", "", "bootstrap peer")
	cmd.PersistentFlags().String("log_level", "", "log level")
	cmd.PersistentFlags().String("private", "", "private key")
	cmd.PersistentFlags().Int("http_port", 0, "api http port")

	return cmd
}
