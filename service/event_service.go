package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/vkhichar/asset-management/config"
	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/domain"
)

const EventResource = "/events"
const AssetMaintenanceEvent = "ASSET_MAINTENANCE_ACTIVITY"

type EventService interface {
	PostAssetEvent(context.Context, *domain.Asset) (string, error)
	PostCreateUserEvent(ctx context.Context, user *domain.User) (string, error)
	PostUpdateUserEvent(context.Context, *domain.User) (string, error)
	PostAssetEventCreateAsset(ctx context.Context, asset *domain.Asset) (string, error)
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

func (e *eventSvc) PostCreateUserEvent(ctx context.Context, user *domain.User) (string, error) {

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

	}

	return string(body), nil
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
		return "", errReadAll
	}
	var responseObj contract.AssetEventResponse
	errJsonMarshal := json.Unmarshal(body, &responseObj)
	if errJsonMarshal != nil {
		fmt.Printf("Event Service : Error While UnMarshaling :%s", errJsonMarshal.Error())
		return "", errJsonMarshal
	}
	// eventId := strconv.Itoa(responseObj.ID)
	return string(responseObj.ID), nil
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
