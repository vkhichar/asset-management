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
	PostAssetEventCreateAsset(ctx context.Context, asset *domain.Asset) (string, error)
	PostUserEvent(context.Context, *domain.User) (string, error)
	PostMaintenanceActivity(ctx context.Context, req domain.MaintenanceActivity) (string, error)
}

type eventSvc struct {
	client *http.Client
}

func NewEventService() EventService {
	return &eventSvc{
		client: &http.Client{
			Timeout: time.Second * time.Duration(config.GetEventApiTimeout()),
		},
	}
}

func (e *eventSvc) PostAssetEventCreateAsset(ctx context.Context, asset *domain.Asset) (string, error) {
	obj := contract.CreateAssetEvent{}
	obj.EventType = "asset"
	obj.Data = asset
	postBody, err := json.Marshal(obj)
	if err != nil {
		fmt.Printf("Event Service: Error while marshaling %s", err.Error())
		return "", err
	}

	responseBody := bytes.NewReader(postBody)

	req, err := http.NewRequest("POST", "http://34.70.86.33:"+config.GetEventAppPort()+"/events", responseBody)

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

func (service *eventSvc) PostMaintenanceActivity(ctx context.Context, req domain.MaintenanceActivity) (string, error) {

	reqBody, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	eventReqBody, _ := json.Marshal(contract.NewEventRequest(AssetMaintenanceEvent, reqBody))

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
