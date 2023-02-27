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

// LoadMaker acts as a thread, which continuously sends Request at ProxyUrl/Endpoint asynchronously.
type LoadMaker struct {
	UID        string
	ProxyUrl   string // with port if IP
	Endpoint   string
	Request    *http.Request
	TargetTime time.Time
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
	lm := &LoadMaker{
		ProxyUrl:   proxyUrl,
		Endpoint:   endpoint,
		Request:    req,
		TargetTime: pc.TargetTime,
	}
	pc.LoadMakers = append(pc.LoadMakers, lm)
	return pc, nil
}

func (lm *LoadMaker) Start(client *http.Client) {
	for lm.TargetTime.After(time.Now()) {
		lm.Request.Header.Set(DelayHeaderParam, fmt.Sprintf("%di", lm.TargetTime.Sub(time.Now()).Milliseconds()))
		_, err := client.Do(lm.Request)
		if err != nil {
			fmt.Printf("Failed with request: %+v\n", *lm.Request)
		}
	}
}
