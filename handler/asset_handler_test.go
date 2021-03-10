package handler_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vkhichar/asset-management/handler"
	mockService "github.com/vkhichar/asset-management/service/mocks"
)

func TestAssetHandler_ListAllAssets_When_ReturnsError(t *testing.T) {
	ctx := context.Background()
	req, err := http.NewRequest("GET", "/assets", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	expectedErr := string(`{"error":"something went wrong"}`)

	mockAssetService := &mockService.MockAssetService{}
	mockAssetService.On("ListAssets", ctx).Return(nil, errors.New("Something went wrong"))

	handlerTest := http.HandlerFunc(handler.ListAssetHandler(mockAssetService))
	handlerTest.ServeHTTP(rr, req)

	assert.JSONEq(t, expectedErr, rr.Body.String())

}
