package node

import (
	"pandora/pkg/p2p"
)

type Service struct {
	p2pService *p2p.Service
}

func New() (*Service, error) {
	p2pService, err := p2p.NewService()
	if err != nil {
		return nil, err
	}

	return &Service{
		p2pService: p2pService,
	}, nil
}

func (s *Service) Run() {
	s.p2pService.Run()
}

func (s *Service) Close() {
	s.p2pService.Close()
}
