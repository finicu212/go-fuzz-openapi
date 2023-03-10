package loadmaker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"go_fuzz_openapi/gen"
	"log"
	"net/http"
	"time"
)

const (
	DelayHeaderParam = "x-proxy-delay"
)

// LoadMaker acts as a thread, which continuously sends Request at ProxyUrl/Endpoint asynchronously until TargetTime.
// All the requests made will have DelayHeaderParam set such that they all land at approximately TargetTime.
type LoadMaker struct {
	UID         string // TODO
	RequestBody string
	Request     *http.Request
	TargetTime  time.Time
}

func (pc *ProxyCoordinator) AddLoadMaker(proxyUrl string, endpoint string, method string, requestBodyName string) (*ProxyCoordinator, error) {
	fmt.Printf("AddLoadMaker(%s, %s, %s, %s)\n", proxyUrl, endpoint, method, requestBodyName)

	req, err := http.NewRequest(method, proxyUrl+"/"+endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed creating http request for proxy %s: %w", proxyUrl, err)
	}
	lm := &LoadMaker{
		UID:         uuid.NewString(),
		RequestBody: requestBodyName,
		Request:     req,
		TargetTime:  pc.TargetTime,
	}
	pc.LoadMakers = append(pc.LoadMakers, lm)
	return pc, nil
}

func (lm *LoadMaker) Start(client *http.Client) {
	for lm.TargetTime.After(time.Now()) {
		lm.Request.Header.Set(DelayHeaderParam, fmt.Sprintf("%di", lm.TargetTime.Sub(time.Now()).Milliseconds()))

		f := gen.GetFakedStructByName(lm.RequestBody)

		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(&f)
		if err != nil {
			log.Printf("could not encode %+v to json: %s", f, err.Error())
		}

		_, err = client.Do(lm.Request)
		if err != nil {
			log.Printf("failed with request: %+v: %s", *lm.Request, err.Error())
		}
	}
}
