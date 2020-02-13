package services

import "time"

type ExpiredChecks struct {
	registry *ServiceRegistry
	ttl      time.Duration
	done     chan bool
}

func NewExpiredChecks(registry *ServiceRegistry, ttl time.Duration) *ExpiredChecks {
	return &ExpiredChecks{registry, ttl, make(chan bool)}
}

func (sec *ExpiredChecks) Start(frequency int) {
	go func() {
		for {
			time.Sleep(time.Duration(int64(sec.ttl) / int64(frequency)))

			select {
			case <-sec.done:
				return
			default:
				sec.registry.FilterExpired(sec.ttl)
			}
		}
	}()
}

func (sec *ExpiredChecks) StartLazy() {
	sec.Start(1)
}

func (sec *ExpiredChecks) Stop() {
	sec.done <- true
}
