package loadmaker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	DelayHeaderParam = "x-proxy-delay"
)

type ProxyCoordinator struct {
	LoadMakers []*LoadMaker
	TargetTime time.Time
	Client     http.Client
}

type LoadMaker struct {
	ProxyUrl string // with port if IP
	Endpoint string
	Request  *http.Request
}

func NewProxyCoordinator(d time.Duration, opts ...ProxyCoordinatorOption) *ProxyCoordinator {
	pc := &ProxyCoordinator{
		LoadMakers: nil,
		TargetTime: time.Now().Add(d),
		Client:     http.Client{}, // Default client unless otherwise specified via WithClient option
	}
	for _, opt := range opts {
		opt(pc)
	}
	return pc
}

type ProxyCoordinatorOption func(pc *ProxyCoordinator)

func WithClient(client http.Client) ProxyCoordinatorOption {
	return func(pc *ProxyCoordinator) {
		pc.Client = client
	}
}

func (pc *ProxyCoordinator) AddLoadMaker(proxyUrl string, endpoint string, method string, body any) (*ProxyCoordinator, error) {
	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(body)
	if err != nil {
		return nil, fmt.Errorf("could not encode %v to json: %w", body, err)
	}
	req, err := http.NewRequest(method, proxyUrl+"/"+endpoint, buf)
	if err != nil {
		return nil, fmt.Errorf("failed creating http request for proxy %s: %w", proxyUrl, err)
	}
	req.Header.Set(DelayHeaderParam, fmt.Sprintf("%di", pc.TargetTime.Sub(time.Now()).Milliseconds()))
	lm := &LoadMaker{
		ProxyUrl: proxyUrl,
		Endpoint: endpoint,
		Request:  req,
	}
	pc.LoadMakers = append(pc.LoadMakers, lm)
	return pc, nil
}
