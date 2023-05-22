package program

import (
	"svc/internal/svc"

	"github.com/kardianos/service"
)

type program struct {
	svc svc.Service
}

func New(svc svc.Service) *program {
	return &program{svc: svc}
}

func (p *program) Start(s service.Service) error {
	p.svc.Start()
	return nil
}

func (p *program) Stop(s service.Service) error {
	p.svc.Stop()
	return nil
}
