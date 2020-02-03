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

type Server struct {
	host  host.Host
	sub   *pubsub.Subscription
	topic *pubsub.Topic

	logger zerolog.Logger
}

func NewServer() (*Server, error) {
	srv := &Server{}

	srv.logger = log.Logger.With().Str(constants.LoggerComponentKey, "p2p").Logger()

	listenAddrs := libp2p.ListenAddrStrings(cfg.Cfg.P2P.ListenAddr)

	identity, err := srv.prepareIdentity()
	if err != nil {
		return nil, err
	}

	var kad *dht.IpfsDHT
	newKad := func(h host.Host) (routing.PeerRouting, error) {
		var err error
		kad, err = dht.New(context.TODO(), h)
		return kad, err
	}
	rtg := libp2p.Routing(newKad)

	hst, err := libp2p.New(
		context.TODO(),
		identity,
		listenAddrs,
		rtg,
	)
	if err != nil {
		return nil, err
	}
	srv.host = hst

	if err := srv.preparePubSub(); err != nil {
		return nil, err
	}

	if err := srv.connectToBootstrapPeer(); err != nil {
		return nil, err
	}

	if err := srv.prepareMDNS(); err != nil {
		return nil, err
	}

	if err := kad.Bootstrap(context.TODO()); err != nil {
		return nil, err
	}

	return srv, nil
}

func (srv *Server) prepareMDNS() error {
	mdns, err := discovery.NewMdnsService(context.TODO(), srv.host, time.Second*10, "")
	if err != nil {
		return err
	}
	mdns.RegisterNotifee(&mdnsNotifee{host: srv.host, logger: srv.logger})
	return nil
}

func (srv *Server) preparePubSub() error {
	pubSub, err := pubsub.NewGossipSub(context.TODO(), srv.host)
	if err != nil {
		return err
	}

	topic, err := pubSub.Join(topic)
	if err != nil {
		return err
	}
	srv.topic = topic

	sub, err := topic.Subscribe()
	if err != nil {
		return err
	}
	srv.sub = sub

	return nil
}

func (srv *Server) prepareIdentity() (libp2p.Option, error) {
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

func (srv *Server) connectToBootstrapPeer() error {
	if cfg.Cfg.P2P.BootstrapPeer != "" {
		multiAddr, err := multiaddr.NewMultiaddr(cfg.Cfg.P2P.BootstrapPeer)
		if err != nil {
			return err
		}

		peerInfo, err := peer.AddrInfoFromP2pAddr(multiAddr)
		if err != nil {
			return err
		}

		if err := srv.host.Connect(context.TODO(), *peerInfo); err != nil {
			return err
		}

		srv.logger.Debug().Msgf("connected to bootstrap peer %s:%s", peerInfo.ID, peerInfo.Addrs[0])
		return nil
	}

	srv.logger.Debug().Msg("there aren't any known bootstrap peers")
	return nil
}

func (srv *Server) Stop() {
	if err := srv.host.Close(); err != nil {
		srv.logger.Err(err).Msg("close host error")
	}
}

func (srv *Server) Start() {
	srv.logger.Debug().Msgf("node %s listening on %s", srv.host.ID(), srv.host.Addrs()[0])

	doneC := make(chan struct{})

	go inputLoop(srv.topic, srv.logger, doneC)
	go pubSubHandler(srv.sub, srv.logger)

	<-doneC
}
