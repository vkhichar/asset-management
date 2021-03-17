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

	"github.com/vkhichar/asset-management/config"
	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/customerrors"
	"github.com/vkhichar/asset-management/domain"
)

type EventService interface {
	PostAssetMaintenanceActivityEvent(context.Context, *domain.MaintenanceActivity) (string, error)
}
type eventSvc struct{}

func NewEventService() EventService {
	return &eventSvc{}
}

func (evSvc *eventSvc) PostAssetMaintenanceActivityEvent(ctx context.Context, resBody *domain.MaintenanceActivity) (string, error) {
	req := contract.CreateAssetMaintenanceEventRequest{}
	req.EventType = "maintenanceactivity"
	req.Data = resBody
	reqEvent, _ := json.Marshal(req)
	r := bytes.NewReader(reqEvent)

	reqst, err := http.NewRequest("POST", config.GetUrl()+":"+config.GetEventAppPort()+"/events", r)
	if err != nil {
		fmt.Printf("Event service request: error:%s", err.Error())
		return "", err
	}
	client := http.Client{
		Timeout: 3 * time.Second,
	}
	resp, err := client.Do(reqst)
	if err != nil {
		fmt.Printf("Event service error while getting response from client.Do: error:%s", err.Error())
		return "", customerrors.ResponseTimeLimitExceeded
	}
	body, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		fmt.Printf("Event service read: error:%s", errRead.Error())
		return "", errRead
	}

	errJsonUnmarshl := json.Unmarshal(body, &contract.CreateMaintenanceActivityEvent)
	if errJsonUnmarshl != nil {
		fmt.Printf("Event service error in unmarshaling: error:%s", errJsonUnmarshl.Error())
		return "", errJsonUnmarshl
	}
	eventConvert := strconv.Itoa(contract.CreateMaintenanceActivityEvent.Id)
	return eventConvert, nil
}
