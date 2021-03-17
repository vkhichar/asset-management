package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/vkhichar/asset-management/domain"
)

type MockEventService struct {
	mock.Mock
}

func (m *MockEventService) PostUserEvent(ctx context.Context, user domain.User) string {

	args := m.Called(ctx, user)

	var newUser string
	if args[0] != nil {
		newUser = args[0].(string)
	}
	return newUser

}
