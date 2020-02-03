package p2p

import (
	"context"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/rs/zerolog"
)

type mdnsNotifee struct {
	host host.Host

	logger zerolog.Logger
}

func (m *mdnsNotifee) HandlePeerFound(pi peer.AddrInfo) {
	if err := m.host.Connect(context.TODO(), pi); err != nil {
		m.logger.Err(err).Msg("establishing connection between host and peer error")
	}
}
