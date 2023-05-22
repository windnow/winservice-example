package svc

import (
	"context"
	"log"
	"sync"
	"time"
)

type svc struct {
	wg     sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
}

func New(ctx context.Context, cancel context.CancelFunc) *svc {

	return &svc{ctx: ctx, cancel: cancel}

}

func (s *svc) Start() {

	s.wg.Add(1)
	go s.working()

	s.wg.Wait()
}
func (s *svc) Stop() {
	s.cancel()
}

type Service interface {
	Start()
	Stop()
}

func (s *svc) working() {
	defer s.wg.Done()

	finishChan := make(chan struct{})

	var counter int = 0
	log.Println("Starting service")

	go func(c chan struct{}) {
	WL:
		for {
			select {
			case <-s.ctx.Done():
				break WL
			default:
				time.Sleep(1 * time.Second)
				counter++
				if counter > 10 {
					break WL
				}

			}
		}

		c <- struct{}{}

	}(finishChan)

	<-finishChan

	log.Println("Finishing")

}
