package loadmaker

import (
	"net/http"
	"sync"
	"time"
)

type ProxyCoordinator struct {
	LoadMakers []*LoadMaker
	TargetTime time.Time
	Client     *http.Client
}

func NewProxyCoordinator(d time.Duration, opts ...ProxyCoordinatorOption) *ProxyCoordinator {
	pc := &ProxyCoordinator{
		LoadMakers: nil,
		TargetTime: time.Now().Add(d),
		Client:     http.DefaultClient, // Default client unless otherwise specified via WithClient option
	}
	for _, opt := range opts {
		opt(pc)
	}
	return pc
}

type ProxyCoordinatorOption func(pc *ProxyCoordinator)

func WithClient(client *http.Client) ProxyCoordinatorOption {
	return func(pc *ProxyCoordinator) {
		pc.Client = client
	}
}

func (pc *ProxyCoordinator) Start() {
	var wg sync.WaitGroup

	for _, lm := range pc.LoadMakers {
		lm := lm // capture forloop value (see gobyexample.com)

		wg.Add(1)
		go func() {
			defer wg.Done()
			lm.Start(pc.Client)
		}()
	}

	wg.Wait()
}
