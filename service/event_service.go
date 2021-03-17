package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/vkhichar/asset-management/config"
	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/domain"
)

const EventResource = "/events"
const AssetMaintenanceEvent = "ASSET_MAINTENANCE_ACTIVITY"

type EventService interface {
<<<<<<< HEAD
	PostUserEvent(context.Context, *domain.User) (string, error)
	PostEvent(ctx context.Context, req domain.MaintenanceActivity) (int, error)
}

type eventSvc struct{}

func NewEventService() EventService {
	return &eventSvc{}
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

func (evSvc *eventSvc) PostEvent(ctx context.Context, req domain.MaintenanceActivity) (int, error) {
=======
	PostEvent(ctx context.Context, req domain.MaintenanceActivity) (string, error)
}

type eventService struct {
	client *http.Client
}

func NewEventService() EventService {
	return &eventService{
		client: &http.Client{
			Timeout: time.Second * time.Duration(config.GetEventApiTimeout()),
		},
	}
}

func (service *eventService) PostEvent(ctx context.Context, req domain.MaintenanceActivity) (string, error) {
>>>>>>> added event service

	reqBody, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	eventReqBody, _ := json.Marshal(contract.NewEventRequest(AssetMaintenanceEvent, reqBody))

	//res, err := service.client.Post(config.GetEventServiceUrl()+EventResource, "application/json", bytes.NewBuffer(eventReqBody))

	httpreq, err := http.NewRequest("POST", config.GetEventServiceUrl()+EventResource, bytes.NewBuffer(eventReqBody))
	httpreq.Header.Add("Content-type", "application/json")
	res, err := service.client.Do(httpreq)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		fmt.Println("Failed to create event due to ", res.StatusCode)
		return "", errors.New("Event not created")
	}

	resBody, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		fmt.Println("Failed to read response\n", err)
		return "", err
	}

	return string(resBody), err
}
