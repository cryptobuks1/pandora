package p2p

import (
	"context"
	"encoding/hex"
	"time"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/routing"
	"github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/p2p/discovery"
	"github.com/multiformats/go-multiaddr"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"pandora/pkg/cfg"
	"pandora/pkg/constants"
)

type Service struct {
	host  host.Host
	sub   *pubsub.Subscription
	topic *pubsub.Topic

	logger zerolog.Logger
}

func NewService() (*Service, error) {
	s := &Service{}

	logger := log.Logger.With().Str(constants.LoggerComponentKey, "p2p").Logger()
	s.logger = logger

	listenAddrs := libp2p.ListenAddrStrings(cfg.Cfg.P2P.ListenAddr)

	identity, err := s.prepareIdentity()
	if err != nil {
		return nil, err
	}

	// TODO move kademlia and routing initialization to separate function
	var kademlia *dht.IpfsDHT
	newKademlia := func(h host.Host) (routing.PeerRouting, error) {
		var err error
		kademlia, err = dht.New(context.TODO(), h)
		return kademlia, err
	}
	routing := libp2p.Routing(newKademlia)

	host, err := libp2p.New(
		context.TODO(),
		identity,
		listenAddrs,
		routing,
	)
	if err != nil {
		return nil, err
	}
	s.host = host

	if err := s.preparePubSub(); err != nil {
		return nil, err
	}

	if err := s.connectToBootstrapPeer(); err != nil {
		return nil, err
	}

	if err := s.prepareMDNS(); err != nil {
		return nil, err
	}

	if err := kademlia.Bootstrap(context.TODO()); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Service) prepareMDNS() error {
	mdns, err := discovery.NewMdnsService(context.TODO(), s.host, time.Second*10, "")
	if err != nil {
		return err
	}
	mdns.RegisterNotifee(&mdnsNotifee{h: s.host, logger: s.logger})
	return nil
}

func (s *Service) preparePubSub() error {
	pubSub, err := pubsub.NewGossipSub(context.TODO(), s.host)
	if err != nil {
		return err
	}

	topic, err := pubSub.Join(topic)
	if err != nil {
		return err
	}
	s.topic = topic

	sub, err := topic.Subscribe()
	if err != nil {
		return err
	}
	s.sub = sub

	return nil
}

func (s *Service) prepareIdentity() (libp2p.Option, error) {
	identity := func(privateKey crypto.PrivKey) libp2p.Option {
		return libp2p.Identity(privateKey)
	}

	if cfg.Cfg.P2P.PrivateKey != "" {
		privateKeyBuf, err := hex.DecodeString(cfg.Cfg.P2P.PrivateKey)
		if err != nil {
			return nil, err
		}

		privateKey, err := crypto.UnmarshalPrivateKey(privateKeyBuf)
		if err != nil {
			return nil, err
		}

		return identity(privateKey), nil
	}

	privateKey, _, err := crypto.GenerateKeyPair(crypto.Ed25519, -1)
	if err != nil {
		return nil, err
	}

	return identity(privateKey), nil
}

func (s *Service) connectToBootstrapPeer() error {
	if cfg.Cfg.P2P.BootstrapPeer != "" {
		multiAddr, err := multiaddr.NewMultiaddr(cfg.Cfg.P2P.BootstrapPeer)
		if err != nil {
			return err
		}

		peerInfo, err := peer.AddrInfoFromP2pAddr(multiAddr)
		if err != nil {
			return err
		}

		if err := s.host.Connect(context.TODO(), *peerInfo); err != nil {
			return err
		}

		s.logger.Debug().Msgf("connected to bootstrap peer %s:%s", peerInfo.ID, peerInfo.Addrs[0])
		return nil
	}

	s.logger.Debug().Msg("there aren't any known bootstrap peers")
	return nil
}

func (s *Service) Stop() {
	if err := s.host.Close(); err != nil {
		s.logger.Err(err).Msg("close host error")
	}
}

func (s *Service) Start() {
	s.logger.Debug().Msgf("node %s listening on %s", s.host.ID(), s.host.Addrs()[0])

	doneC := make(chan struct{})

	go inputLoop(s.topic, s.logger, doneC)
	go pubSubHandler(s.sub, s.logger)

	<-doneC
}
