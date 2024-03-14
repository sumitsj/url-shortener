package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/sumitsj/url-shortener/models"
	"github.com/sumitsj/url-shortener/repositories/mocks"
	"testing"
)

func TestUrlService_GenerateShortUrl(t *testing.T) {
	url := "www.google.com"
	expectedShortUrl := "http://localhost:8080/short/"

	repository := mocks.UrlMappingRepository{}
	repository.On("Create", models.URLMapping{
		OriginalUrl:  url,
		ShortenedUrl: expectedShortUrl,
	}).Return(nil)

	service := CreateUrlService(&repository)
	shortUrl := service.GenerateShortUrl(url)
	assert.Contains(t, shortUrl, expectedShortUrl)
	assert.Equal(t, 34, len(shortUrl))
}
