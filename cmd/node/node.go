package main

import (
	"github.com/rs/zerolog/log"

	"pandora/pkg/cli"
)

func main() {
	if err := cli.Cmd().Execute(); err != nil {
		log.Fatal().Err(err).Msg("execute command error")
	}
}
