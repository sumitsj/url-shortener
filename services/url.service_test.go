package services

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/sumitsj/url-shortener/repositories/mocks"
	mocks2 "github.com/sumitsj/url-shortener/services/mocks"
	"testing"
)

func TestUrlService_GenerateShortUrl(t *testing.T) {
	url := "www.google.com"
	expectedShortUrl := "http://localhost:8080/s/"

	repository := mocks.NewUrlMappingRepository(t)
	repository.On("GetByOriginalUrl", url).Return(nil, errors.New("not found"))
	repository.On("Create", mock.AnythingOfType("*models.URLMapping")).Return(nil)

	appConfig := mocks2.NewAppConfig(t)
	appConfig.On("GetServerAddr").Return("http://localhost")
	appConfig.On("GetServerPort").Return("8080")

	service := CreateUrlService(appConfig, repository)

	shortUrl := service.GenerateShortUrl(url)

	assert.Contains(t, shortUrl, expectedShortUrl)
	assert.Equal(t, 30, len(shortUrl))
	repository.AssertNumberOfCalls(t, "Create", 1)
	repository.AssertNumberOfCalls(t, "GetByOriginalUrl", 1)
	appConfig.AssertNumberOfCalls(t, "GetServerAddr", 1)
	appConfig.AssertNumberOfCalls(t, "GetServerPort", 1)
}
