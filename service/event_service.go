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
	PostUserEvent(ctx context.Context, user *domain.User) (string, error)
}

type eventSvc struct{}

func NewEventService() EventService {
	return &eventSvc{}
}

type Client struct{}

func (e *eventSvc) PostUserEvent(ctx context.Context, user *domain.User) (string, error) {

	object := contract.CreateUserEvent{
		EventType: "user",
		Data:      user,
	}
	postBody, err := json.Marshal(object)
	if err != nil {
		fmt.Printf("Error while marshaling in event service: %s", err.Error())
		return "", err
	}
	responseBody := bytes.NewReader(postBody)

	re, err := http.NewRequest("POST", "http://34.70.86.33:"+config.GetEventAppPort()+"/events", responseBody)
	if err != nil {
		fmt.Printf("event service: error while newrequest: %s", err.Error())
		return "", err
	}

	client := http.Client{
		Timeout: 3 * time.Second,
	}

	resp, err := client.Do(re)
	if err != nil {
		fmt.Printf("Error in event service while getting response in client.do: %s", err.Error())
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error in event service :%s", err.Error())
		return "", err
	}

	return string(body), nil
}
