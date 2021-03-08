package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/vkhichar/asset-management/service"
)

type MockTokenService struct {
	mock.Mock
}

func (m *MockTokenService) GenerateToken(c *service.Claims) (string, error) {
	args := m.Called(c)

	if args[1] != nil {
		return args[0].(string), args[1].(error)
	}

	return args[0].(string), nil
}

func (m *MockTokenService) ValidateToken(token string) (*service.Claims, error) {
	// TODO: define mock method

	return nil, nil
}
