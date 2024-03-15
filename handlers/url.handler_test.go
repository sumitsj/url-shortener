package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sumitsj/url-shortener/constants"
	"github.com/sumitsj/url-shortener/contracts"
	"github.com/sumitsj/url-shortener/services/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUrlHandler_HandleRedirect(t *testing.T) {
	shortKey := "abcdef"
	shortenedUrl := "http://localhost:8080/s/abcdef"
	originalUrl := "www.google.com"
	engine := gin.Default()

	urlService := mocks.NewUrlService(t)
	urlService.On("FormatShortUrl", shortKey).Return(shortenedUrl)
	urlService.On("GetOriginalUrlBy", shortenedUrl).Return(originalUrl, nil)

	handler := CreateUrlHandler(urlService)
	engine.GET("/redirect/:shortKey", handler.HandleRedirect)

	responseRecorder := httptest.NewRecorder()

	url := fmt.Sprintf("/redirect/%v", shortKey)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	engine.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusMovedPermanently, responseRecorder.Code)
}

func TestUrlHandler_HandleRedirect_ShouldReturnErrorResponseIfShortUrlNotFound(t *testing.T) {
	shortKey := "abcdef"
	shortenedUrl := "http://localhost:8080/s/abcdef"
	engine := gin.Default()

	urlService := mocks.NewUrlService(t)
	urlService.On("FormatShortUrl", shortKey).Return(shortenedUrl)
	urlService.On("GetOriginalUrlBy", shortenedUrl).Return("", errors.New("not found"))

	handler := CreateUrlHandler(urlService)
	engine.GET("/redirect/:shortKey", handler.HandleRedirect)

	responseRecorder := httptest.NewRecorder()

	url := fmt.Sprintf("/redirect/%v", shortKey)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	engine.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusNotFound, responseRecorder.Code)
}

func TestUrlHandler_ShortenUrl(t *testing.T) {
	shortenedUrl := "http://localhost:8080/s/abcdef"
	originalUrl := "www.google.com"
	engine := gin.Default()

	urlService := mocks.NewUrlService(t)
	urlService.On("GenerateShortUrl", originalUrl).Return(shortenedUrl, nil)

	handler := CreateUrlHandler(urlService)
	engine.POST("/short", handler.ShortenUrl)

	responseRecorder := httptest.NewRecorder()

	var body = []byte(fmt.Sprintf(`{"URL": "%s" }`, originalUrl))
	req, _ := http.NewRequest(http.MethodPost, "/short", bytes.NewBuffer(body))
	engine.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	response, err := readResponseBody(responseRecorder)
	require.NoError(t, err)
	assert.Equal(t, shortenedUrl, response.ShortenedUrl)
	assert.Equal(t, originalUrl, response.OriginalUrl)
	assert.Equal(t, "", response.Error)
}

func TestUrlHandler_ShortenUrl_ShouldReturnInternalServerError(t *testing.T) {
	originalUrl := "www.google.com"
	engine := gin.Default()

	urlService := mocks.NewUrlService(t)
	urlService.On("GenerateShortUrl", originalUrl).Return("", errors.New("db error"))

	handler := CreateUrlHandler(urlService)
	engine.POST("/short", handler.ShortenUrl)

	responseRecorder := httptest.NewRecorder()

	var body = []byte(fmt.Sprintf(`{"URL": "%s" }`, originalUrl))
	req, _ := http.NewRequest(http.MethodPost, "/short", bytes.NewBuffer(body))
	engine.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusInternalServerError, responseRecorder.Code)

	response, err := readResponseBody(responseRecorder)
	require.NoError(t, err)
	assert.Equal(t, "", response.ShortenedUrl)
	assert.Equal(t, "", response.OriginalUrl)
	assert.Equal(t, constants.InternalServerErrorMessage, response.Error)
}

func TestUrlHandler_ShortenUrl_ShouldReturnBadRequestError(t *testing.T) {
	engine := gin.Default()
	urlService := mocks.NewUrlService(t)
	handler := CreateUrlHandler(urlService)
	engine.POST("/short", handler.ShortenUrl)

	responseRecorder := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/short", nil)
	engine.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	response, err := readResponseBody(responseRecorder)
	require.NoError(t, err)
	assert.Equal(t, "", response.ShortenedUrl)
	assert.Equal(t, "", response.OriginalUrl)
	assert.Equal(t, constants.RequestParsingErrorMessage, response.Error)
}

func readResponseBody(responseRecorder *httptest.ResponseRecorder) (contracts.ShortenUrlResponse, error) {
	var response contracts.ShortenUrlResponse
	err := json.NewDecoder(responseRecorder.Body).Decode(&response)
	return response, err
}
