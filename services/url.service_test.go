package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUrlService_GenerateShortUrl(t *testing.T) {
	service := CreateUrlService()
	shortUrl := service.GenerateShortUrl("www,google.com")
	assert.Contains(t, shortUrl, "http://localhost:8080/short/")
	assert.Equal(t, 34, len(shortUrl))
}
