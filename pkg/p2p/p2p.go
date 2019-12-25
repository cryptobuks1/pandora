package p2p

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/routing"
	"github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p-mplex"
	"github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p-secio"
	"github.com/libp2p/go-libp2p-yamux"
	"github.com/libp2p/go-libp2p/p2p/discovery"
	"github.com/libp2p/go-tcp-transport"
	"github.com/multiformats/go-multiaddr"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"pandora/pkg/cfg"
	"pandora/pkg/constants"
)

type Service struct {
	host   host.Host
	sub    *pubsub.Subscription
	pubsub *pubsub.PubSub
	topic  *pubsub.Topic
	logger zerolog.Logger
}

func NewService() (*Service, error) {
	logger := log.Logger.With().Str(constants.LoggerComponentKey, "p2p").Logger()

	transports := libp2p.ChainOptions(
		libp2p.Transport(tcp.NewTCPTransport),
	)

	muxers := libp2p.ChainOptions(
		libp2p.Muxer("/yamux/1.0.0", sm_yamux.DefaultTransport),
		libp2p.Muxer("/mplex/6.7.0", peerstream_multiplex.DefaultTransport),
	)

	security := libp2p.Security(secio.ID, secio.New)

	listenAddrs := libp2p.ListenAddrStrings(cfg.Cfg.P2P.ListenAddresses...)

	var kademlia *dht.IpfsDHT
	newKademlia := func(h host.Host) (routing.PeerRouting, error) {
		var err error
		kademlia, err = dht.New(context.TODO(), h)
		return kademlia, err
	}
	routing := libp2p.Routing(newKademlia)

	host, err := libp2p.New(
		context.TODO(),
		transports,
		listenAddrs,
		muxers,
		security,
		routing,
	)
	if err != nil {
		return nil, err
	}

	logger.Debug().Msgf("host id %s", host.ID())

	pubsub, err := pubsub.NewGossipSub(context.TODO(), host)
	if err != nil {
		return nil, err
	}
	topic, err := pubsub.Join(topic)
	if err != nil {
		return nil, err
	}
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, err
	}

	for _, addr := range host.Addrs() {
		logger.Debug().Msgf("listening on %s", addr)
	}

	if cfg.Cfg.P2P.BootstrapPeer != "" {
		multiAddr, err := multiaddr.NewMultiaddr(cfg.Cfg.P2P.BootstrapPeer)
		if err != nil {
			return nil, err
		}

		addrInfo, err := peer.AddrInfoFromP2pAddr(multiAddr)
		if err != nil {
			return nil, err
		}

		if err := host.Connect(context.TODO(), *addrInfo); err != nil {
			return nil, err
		}

		logger.Debug().Msgf("connected to %s", addrInfo.ID)
	}

	mdns, err := discovery.NewMdnsService(context.TODO(), host, time.Second*10, "")
	if err != nil {
		return nil, err
	}
	mdns.RegisterNotifee(&mdnsNotifee{h: host, ctx: context.TODO(), logger: logger})

	if err := kademlia.Bootstrap(context.TODO()); err != nil {
		return nil, err
	}

	return &Service{
		host:   host,
		logger: logger,
		sub:    sub,
		pubsub: pubsub,
		topic:  topic,
	}, nil
}

func (s *Service) Close() {
	if err := s.host.Close(); err != nil {
		s.logger.Err(err).Msg("close host error")
	}
}

func (s *Service) Run() {
	doneC := make(chan struct{})

	go chatInputLoop(s.topic, s.logger, doneC)
	go pubsubHandler(s.sub, s.logger)

	select {
	case <-doneC:
		return
	}
}
