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
	PostUserEvent(context.Context, *domain.User) (string, error)
}

type eventSvc struct{}

func NewEventService() EventService {
	return &eventSvc{}
}

func (e *eventSvc) PostAssetEvent(ctx context.Context, asset *domain.Asset) (string, error) {
	object := contract.UpdateAssetEvent{}
	object.EvenType = "asset"
	object.Data = asset
	Revent, err := json.Marshal(object)
	if err != nil {
		fmt.Printf("Error while Marshaling %s", err.Error())
		return "", err
	}
	r := bytes.NewReader(Revent)
	rec, errReq := http.NewRequest("POST", "http://34.70.86.33:"+config.GetEventAppPort()+"/events", r)
	if errReq != nil {
		fmt.Printf("Error in http request %s", errReq.Error())
		return "", errReq
	}
	client := http.Client{
		Timeout: 3 * time.Second,
	}
	resp, errPost := client.Do(rec)
	if errPost != nil {
		fmt.Printf("Event service: Error while sending Post request to event. Error: %s", errPost.Error())
		return "", errPost
	}

	body, errReadAll := ioutil.ReadAll(resp.Body)
	if errReadAll != nil {
		fmt.Printf("Error While Performing ReadAll %s", errReadAll.Error())
	}
	return string(body), nil
}
func (evSvc *eventSvc) PostUserEvent(ctx context.Context, user *domain.User) (string, error) {
	request := contract.UpdateUserEventRequest{}
	request.EventType = "user"
	request.Data = user
	reqEvent, errJson := json.Marshal(request)

	if errJson != nil {
		fmt.Printf("Event service: Error while json marshal. Error: %s", errJson.Error())
		return "", errJson
	}

	r := bytes.NewReader(reqEvent)

	req, errNewReq := http.NewRequest("POST", "http://34.70.86.33:"+config.GetEventAppPort()+"/events", r)
	if errNewReq != nil {
		fmt.Printf("Event service: Error while sending Post request to event. Error: %s", errNewReq.Error())
		return "", errNewReq
	}
	client := &http.Client{
		Timeout: time.Second * 3,
	}
	resp, errPost := client.Do(req)

	if errPost != nil {
		fmt.Printf("Event service: Error while sending Post request to event. Error: %s", errPost.Error())
		return "", errPost
	}

	body, errBodyRead := ioutil.ReadAll(resp.Body)
	if errBodyRead != nil {
		fmt.Printf("Event service: Error while reading response body. Error: %s", errBodyRead.Error())
		return "", errBodyRead
	}
	return string(body), nil
}
