package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/vkhichar/asset-management/config"
	"github.com/vkhichar/asset-management/domain"
	"github.com/vkhichar/asset-management/service"
	"gopkg.in/h2non/gock.v1"
)

func TestPostEvent_AssetMaintenanceActivity_When_Success(t *testing.T) {
	defer gock.Off()
	activity := domain.MaintenanceActivity{
		ID:      1,
		AssetId: uuid.New(),
		Cost:    25,
	}
	gock.New(config.GetEventServiceUrl()).
		Post(service.EventResource).
		Reply(200).
		JSON(map[string]int{"id": 1})

	eventService := service.NewEventService()
	resp, err := eventService.PostEvent(context.Background(), activity)

	assert.Nil(t, err)
	assert.JSONEq(t, `{"id": 1}`, resp)
}

func TestPostEvent_AssetMaintenanceActivity_When_Failed(t *testing.T) {
	defer gock.Off()
	gock.New(config.GetEventServiceUrl()).
		Post(service.EventResource).
		Reply(400) // testing 400 request

	eventService := service.NewEventService()
	_, err := eventService.PostEvent(context.Background(), domain.MaintenanceActivity{})

	assert.NotNil(t, err)

}

func TestPostEvent_AssetMaintenanceActivity_When_TimeoutError(t *testing.T) {
	defer gock.Off()
	gock.New(config.GetEventServiceUrl()).
		Post(service.EventResource).
		ReplyError(errors.New("Timeout error"))

	eventService := service.NewEventService()
	_, err := eventService.PostEvent(context.Background(), domain.MaintenanceActivity{})

	assert.NotNil(t, err)

}
