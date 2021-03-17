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
	PostAssetEvent(context.Context, *domain.Asset) (string, error)
}

type eventsvc struct{}

func NewEventService() EventService {
	return &eventsvc{}
}

func (e *eventsvc) PostAssetEvent(ctx context.Context, asset *domain.Asset) (string, error) {
	object := contract.UpdateAssetEvent{}
	object.EvenType = "asset"
	object.Data = asset
	Revent, err := json.Marshal(object)
	if err != nil {
		fmt.Printf("Error while Marshaling %s", err.Error())
		return "", err
	}
	r := bytes.NewReader(Revent)
	rec, errReq := http.NewRequest("POST", "http://34.70.86.33:"+config.GetEventPort()+"/events", r)
	if errReq != nil {
		fmt.Printf("Error in http request %s", errReq.Error())
		return "", errReq
	}
	client := http.Client{
		Timeout: 3 * time.Second,
	}
	resp, errClient := client.Do(rec)
	if errClient != nil {
		fmt.Printf("EventService: Error in Client.Do %s", errClient.Error())
		return "", errClient
	}

	body, errReadAll := ioutil.ReadAll(resp.Body)
	if errReadAll != nil {
		fmt.Printf("Error While Performing ReadAll %s", errReadAll.Error())
	}
	return string(body), nil
}
