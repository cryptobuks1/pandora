package main

import (
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"pandora/pkg/cfg"
	"pandora/pkg/node"
)

var cli = &cobra.Command{
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})

		if err := cfg.Prepare(cmd); err != nil {
			log.Fatal().Err(err).Msg("prepare config file error")
		}

		level, err := zerolog.ParseLevel(strings.ToLower(cfg.Cfg.LogLevel))
		if err != nil {
			log.Fatal().Err(err).Msg("parse log level from config error")
		}
		if level == zerolog.NoLevel {
			log.Logger = log.Logger.Level(zerolog.TraceLevel)
		} else {
			log.Logger = log.
				Logger.
				Level(level)
		}
		if log.Logger.GetLevel() == zerolog.TraceLevel {
			log.Logger = log.Logger.With().Caller().Logger()
		}
	},

	Run: func(cmd *cobra.Command, args []string) {
		node, err := node.New()
		if err != nil {
			log.Fatal().Err(err).Msg("initialize node service error")
		}

		go func() {
			node.Run()
		}()

		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		<-interrupt
		log.Info().Msg("handle SIGINT, SIGTERM, SIGQUIT")

		node.Close()
	},
}

func main() {
	cli.PersistentFlags().String("listen", "", "listen address")
	cli.PersistentFlags().String("boot", "", "bootstrap peer")

	if err := cli.Execute(); err != nil {
		log.Fatal().Err(err).Msg("execute node error")
	}
}
