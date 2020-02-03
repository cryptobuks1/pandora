package node

import (
	"github.com/rs/zerolog/log"

	"pandora/pkg/api"
	"pandora/pkg/p2p"
)

type Node struct {
	p2pSrv *p2p.Server
	apiSrv *api.Server
}

func NewNode() (*Node, error) {
	p2pSrv, err := p2p.NewServer()
	if err != nil {
		return nil, err
	}

	apiSrv := api.NewServer()

	return &Node{
		apiSrv: apiSrv,
		p2pSrv: p2pSrv,
	}, nil
}

func (n *Node) Start() {
	go func() {
		if err := n.apiSrv.Start(); err != nil {
			log.Err(err).Msg("running api service error")
		}
	}()

	n.p2pSrv.Start()
}

func (n *Node) Stop() {
	n.p2pSrv.Stop()
	n.apiSrv.Stop()
}
