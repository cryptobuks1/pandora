package node

import (
	"pandora/pkg/p2p"
)

type Service struct {
	p2pService *p2p.Service
}

func NewService() (*Service, error) {
	p2pService, err := p2p.NewService()
	if err != nil {
		return nil, err
	}

	return &Service{
		p2pService: p2pService,
	}, nil
}

func (s *Service) Run() error {
	s.p2pService.Run()
	return nil
}

func (s *Service) Close() {
	s.p2pService.Close()
}
