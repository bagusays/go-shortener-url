package mock

import (
	"url-shortener/models"

	"github.com/stretchr/testify/mock"
)

// Create a MockStore struct with an embedded mock instance
type MockShortenerRepo struct {
	mock.Mock
}

func (m *MockShortenerRepo) FindByShortLink(shortUrl string) (*models.Link, error) {
	// This allows us to pass in mocked results, so that the mock store will return whatever we define
	args := m.Called()
	result := args.Get(0)
	// return the values which we define
	return result.(*models.Link), args.Error(1)
}

func (m *MockShortenerRepo) CreateLink(data models.Link) error {
	// This allows us to pass in mocked results, so that the mock store will return whatever we define
	args := m.Called()
	// result := args.Get(0)
	// return the values which we define
	return args.Error(0)
}

func (m *MockShortenerRepo) EditLink(data models.Link) error {
	// This allows us to pass in mocked results, so that the mock store will return whatever we define
	args := m.Called()
	// result := args.Get(0)
	// return the values which we define
	return args.Error(0)
}

func (m *MockShortenerRepo) DeleteLink(data models.DeleteLink) error {
	// This allows us to pass in mocked results, so that the mock store will return whatever we define
	args := m.Called()
	// result := args.Get(0)
	// return the values which we define
	return args.Error(0)
}
