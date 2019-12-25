package p2p

import (
	"context"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/rs/zerolog"
)

type mdnsNotifee struct {
	h      host.Host
	ctx    context.Context
	logger zerolog.Logger
}

func (m *mdnsNotifee) HandlePeerFound(pi peer.AddrInfo) {
	if err := m.h.Connect(m.ctx, pi); err != nil {
		m.logger.Err(err).Msg("establishing connection between host and peer error")
	}
}
