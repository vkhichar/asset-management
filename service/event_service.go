package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/domain"
)

type EventService interface {
	PostUpdateUserEvent(context.Context, *domain.User) (string, error)
}

type eventSvc struct{}

func NewEventService() EventService {
	return &eventSvc{}
}

func (evSvc *eventSvc) PostUpdateUserEvent(ctx context.Context, user *domain.User) (string, error) {
	request := contract.UpdateUserEventRequest{}
	request.EventType = "user"
	request.Data = user
	reqEvent, errJson := json.Marshal(request)

	if errJson != nil {
		fmt.Printf("Event service: Error while json marshal. Error: %s", errJson.Error())
		return "", errJson
	}

	r := bytes.NewReader(reqEvent)

	//req, errNewReq := http.NewRequest("POST", config.GetIpAddress()+":"+config.GetEventAppPort()+"/events", r)
	req, errNewReq := http.NewRequest("POST", "http://34.70.86.33:9035/events", r)
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

	var responseObj contract.UpdateUserEventResponse

	errJsonUnmar := json.Unmarshal(body, &responseObj)

	if errJsonUnmar != nil {
		fmt.Printf("Event service: Error while json unmarshal. Error: %s", errJsonUnmar.Error())
		return "", errJsonUnmar
	}

	eventId := strconv.Itoa(responseObj.Id)

	return eventId, nil
}
