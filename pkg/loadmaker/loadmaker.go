package loadmaker

import (
	"net/http"
	"time"
)

const (
	DelayHeaderParam = "x-proxy-delay"
)

type ProxyCoordinator struct {
	LoadMakers []LoadMaker
	TargetTime time.Time
	Client     http.Client
}

type LoadMaker struct {
	Request  http.Request
	ProxyUrl string // with port if IP
	Endpoint string
}

func (pc *ProxyCoordinator) New(proxyUrl string, endpoint string, method string, body any) *LoadMaker {
	lm := LoadMaker{
		ProxyUrl: proxyUrl,
		Endpoint: endpoint,
		Request:  http.NewRequest(method, proxyUrl),
	}
}

func New(proxyIP string, request http.Request) *ProxyCoordinator {

}
