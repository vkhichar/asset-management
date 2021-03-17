package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/vkhichar/asset-management/config"
	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/domain"
)

type EventService interface {
	PostAssetEvent(ctx context.Context, asset *domain.Asset) (string, error)
}

type eventSvc struct{}

func NewEventService() EventService {
	return &eventSvc{}
}

func (e *eventSvc) PostAssetEvent(ctx context.Context, asset *domain.Asset) (string, error) {
	obj := contract.CreateAssetEvent{}
	obj.EventType = "asset"
	obj.Data = asset
	postBody, err := json.Marshal(obj)
	if err != nil {
		fmt.Printf("Event Service: Error while marshaling %s", err.Error())
		return "", err
	}

	responseBody := bytes.NewReader(postBody)

	req, err := http.NewRequest("POST", "http://34.70.86.33:"+config.GetEventPort()+"/events", responseBody)

	if err != nil {
		fmt.Printf("Event Service: error during http request %s:", err.Error())
		return "", err
	}

	var netClient = &http.Client{
		Timeout: time.Second * 3,
	}

	response, err := netClient.Do(req)

	if err != nil {
		fmt.Printf("Event server: Request Timeout %s: Taking more than %v seconds", err.Error(), response)
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Event Service : error while converting into byte stream: %s", err.Error())
		return "", err
	}

	return string(body), nil
}
