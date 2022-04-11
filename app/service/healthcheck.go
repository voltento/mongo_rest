package service

import "sync"

type HealthCheck interface {
	HealthCheck() error
}

type HealthChecker struct {
	services []HealthCheck
	wg       sync.WaitGroup
	err      chan error
}

func NewHealthChecker() *HealthChecker {
	return &HealthChecker{
		services: []HealthCheck{},
		err:      make(chan error, 1),
	}
}

func (h *HealthChecker) Healthy() error {

	for _, s := range h.services {
		h.wg.Add(1)
		go func() {
			if err := s.HealthCheck(); err != nil {
				h.err <- err
			}
			defer h.wg.Done()
		}()
	}

	go func() {
		h.wg.Wait()
		h.err <- nil
	}()
	return <-h.err
}

func (h *HealthChecker) RegisterService(s HealthCheck) {
	h.services = append(h.services, s)
}
