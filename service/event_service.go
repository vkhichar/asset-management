package service

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/domain"
)

type EventService interface {
	PostUserEvent(context.Context, *domain.User) string
}

type eventSvc struct{}

func NewEventService() EventService {
	return &eventSvc{}
}

func (evSvc *eventSvc) PostUserEvent(ctx context.Context, user *domain.User) string {
	request := contract.UpdateUserEventRequest{}
	request.EventType = "user"
	request.Data = user
	reqEvent, _ := json.Marshal(request)
	r := bytes.NewReader(reqEvent)
	resp, _ := http.Post("http://34.70.86.33:9035/events", "application/json", r)
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
