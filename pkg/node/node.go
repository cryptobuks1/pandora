package node

import (
	"github.com/rs/zerolog/log"

	"pandora/pkg/api"
	"pandora/pkg/p2p"
)

type Service struct {
	p2pService *p2p.Service
	apiService *api.Service
}

func NewService() (*Service, error) {
	p2pService, err := p2p.NewService()
	if err != nil {
		return nil, err
	}

	apiService := api.NewService()

	return &Service{
		apiService: apiService,
		p2pService: p2pService,
	}, nil
}

func (s *Service) Start() {
	go func() {
		if err := s.apiService.Start(); err != nil {
			log.Err(err).Msg("running api service error")
		}
	}()

	s.p2pService.Start()
}

func (s *Service) Stop() {
	s.p2pService.Stop()
	s.apiService.Stop()
}
