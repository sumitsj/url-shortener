package services

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/sumitsj/url-shortener/models"
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

	shortUrl, _ := service.GenerateShortUrl(url)

	assert.Contains(t, shortUrl, expectedShortUrl)
	assert.Equal(t, 30, len(shortUrl))
	repository.AssertExpectations(t)
	appConfig.AssertExpectations(t)
}

func TestUrlService_GenerateShortUrl_ShouldReturnExistingShortUrlIfExists(t *testing.T) {
	url := "www.google.com"
	shortKey := "abcdef"
	expectedShortUrl := "http://localhost:8080/s/" + shortKey

	repository := mocks.NewUrlMappingRepository(t)
	repository.On("GetByOriginalUrl", url).Return(&models.URLMapping{
		OriginalUrl: url,
		ShortKey:    shortKey,
	}, nil)

	appConfig := mocks2.NewAppConfig(t)
	appConfig.On("GetServerAddr").Return("http://localhost")
	appConfig.On("GetServerPort").Return("8080")

	service := CreateUrlService(appConfig, repository)

	shortUrl, _ := service.GenerateShortUrl(url)

	assert.Equal(t, expectedShortUrl, shortUrl)
	repository.AssertNumberOfCalls(t, "Create", 0)
	repository.AssertExpectations(t)
}

func TestUrlService_GenerateShortUrl_ShouldReturnErrorInCaseOfDbIssue(t *testing.T) {
	url := "www.google.com"

	repository := mocks.NewUrlMappingRepository(t)
	repository.On("GetByOriginalUrl", url).Return(nil, errors.New("document not found"))
	repository.On("Create", mock.AnythingOfType("*models.URLMapping")).Return(errors.New("connection failed"))

	appConfig := mocks2.NewAppConfig(t)

	service := CreateUrlService(appConfig, repository)

	shortUrl, err := service.GenerateShortUrl(url)

	assert.Equal(t, "", shortUrl)
	assert.Error(t, err)
	assert.Equal(t, "Failed to save url mapping for url: www.google.com.\nError: connection failed", err.Error())
}

func TestUrlService_GetOriginalUrlBy(t *testing.T) {
	url := "www.google.com"
	shortenedUrl := "http://localhost:8080/s/abcdef"

	repository := mocks.NewUrlMappingRepository(t)
	repository.On("GetByShortKey", shortenedUrl).Return(&models.URLMapping{
		OriginalUrl: url,
		ShortKey:    shortenedUrl,
	}, nil)

	appConfig := mocks2.NewAppConfig(t)
	service := CreateUrlService(appConfig, repository)

	originalUrl, _ := service.GetOriginalUrlBy(shortenedUrl)

	assert.Equal(t, url, originalUrl)
	repository.AssertExpectations(t)
}

func TestUrlService_GetOriginalUrlBy_ShouldReturnErrorInCaseOfDbIssue(t *testing.T) {
	shortenedUrl := "http://localhost:8080/s/abcdef"
	errorMessage := "db error"

	repository := mocks.NewUrlMappingRepository(t)
	repository.On("GetByShortKey", shortenedUrl).Return(nil, errors.New(errorMessage))

	appConfig := mocks2.NewAppConfig(t)
	service := CreateUrlService(appConfig, repository)

	originalUrl, err := service.GetOriginalUrlBy(shortenedUrl)

	assert.Equal(t, "", originalUrl)
	assert.Error(t, err)
	assert.Equal(t, errorMessage, err.Error())
	repository.AssertExpectations(t)
}
